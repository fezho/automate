package processor

import (
	"context"

	log "github.com/sirupsen/logrus"

	"github.com/chef/automate/api/interservice/authz"
	"github.com/chef/automate/components/ingest-service/backend"
	"github.com/chef/automate/components/ingest-service/pipeline/message"
	"github.com/chef/automate/lib/stringutils"
)

// BuildRunProjectTagger - Build a project tagger for CCRs
func BuildRunProjectTagger(authzClient authz.ProjectsClient) message.ChefRunPipe {
	return func(in <-chan message.ChefRun) <-chan message.ChefRun {
		return runBundleProjectTagger(in, authzClient)
	}
}

// This processor is bundling all the messages that are currently in the 'in' channel. The bundling
// of messages decreases the number of times the authz-service is called for project rules. The way
// it works is when a message comes in, we make a call to the authz-service for the rules. We use
// these rules for all the messages that are currently in the queue. The 'bundleSize' is the number
// of messages that can use the current project rules from authz.
func runBundleProjectTagger(in <-chan message.ChefRun,
	authzClient authz.ProjectsClient) <-chan message.ChefRun {
	out := make(chan message.ChefRun, 100)
	go func() {
		bundleSize := 0
		var projectRulesCollection map[string]*authz.ProjectRules
		for msg := range in {
			if bundleSize <= 0 {
				bundleSize = len(in)
				log.WithFields(log.Fields{
					"message_id": msg.ID,
					"bundleSize": bundleSize,
				}).Debug("BundleProjectTagging - Update Project rules")
				projectRulesCollection = getProjectRulesFromAuthz(msg.Ctx, authzClient)
			} else {
				// Skip
				bundleSize--
			}

			msg.Node.Projects = findMatchingProjects(msg.Node, projectRulesCollection)

			out <- msg
		}
		close(out)
	}()

	return out
}

func findMatchingProjects(node backend.Node, projects map[string]*authz.ProjectRules) []string {
	matchingProjects := make([]string, 0)

	for projectName, project := range projects {
		if nodeMatchesRules(node, project.Rules) {
			matchingProjects = append(matchingProjects, projectName)
		}
	}

	return matchingProjects
}

func getProjectRulesFromAuthz(ctx context.Context,
	authzClient authz.ProjectsClient) map[string]*authz.ProjectRules {
	projectsCollection, err := authzClient.ListRulesForAllProjects(ctx, &authz.ListRulesForAllProjectsReq{})
	if err != nil {
		// If there is an error getting the project rules from authz crash the service.
		log.WithError(err).Fatal("Could not fetch project rules from authz")
	}

	return projectsCollection.ProjectRules
}

// Only one rule has to be true for the project to match (ORed together).
func nodeMatchesRules(node backend.Node, rules []*authz.ProjectRule) bool {
	for _, rule := range rules {
		if rule.Type == authz.ProjectRuleTypes_NODE && nodeMatchesAllConditions(node, rule.Conditions) {
			return true
		}
	}

	return false
}

// All the conditions must be true for a rule to be true (ANDed together).
// If there are no conditions then the rule is false
func nodeMatchesAllConditions(node backend.Node, conditions []*authz.Condition) bool {
	if len(conditions) == 0 {
		return false
	}

	for _, condition := range conditions {
		switch condition.Attribute {
		case authz.ProjectRuleConditionAttributes_CHEF_SERVER:
			if !stringutils.SliceContains(condition.Values, node.SourceFqdn) {
				return false
			}
		case authz.ProjectRuleConditionAttributes_CHEF_ORGANIZATION:
			if !stringutils.SliceContains(condition.Values, node.OrganizationName) {
				return false
			}
		case authz.ProjectRuleConditionAttributes_ENVIRONMENT:
			if !stringutils.SliceContains(condition.Values, node.Environment) {
				return false
			}
		case authz.ProjectRuleConditionAttributes_CHEF_POLICY_GROUP:
			if !stringutils.SliceContains(condition.Values, node.PolicyGroup) {
				return false
			}
		case authz.ProjectRuleConditionAttributes_CHEF_POLICY_NAME:
			if !stringutils.SliceContains(condition.Values, node.PolicyName) {
				return false
			}
		case authz.ProjectRuleConditionAttributes_CHEF_ROLE:
			foundMatch := false
			for _, projectRole := range condition.Values {
				if stringutils.SliceContains(node.Roles, projectRole) {
					foundMatch = true
					break
				}
			}
			if !foundMatch {
				return false
			}
		case authz.ProjectRuleConditionAttributes_CHEF_TAG:
			foundMatch := false
			for _, projectRole := range condition.Values {
				if stringutils.SliceContains(node.ChefTags, projectRole) {
					foundMatch = true
					break
				}
			}
			if !foundMatch {
				return false
			}
		default:
			return false
		}
	}

	return true
}
