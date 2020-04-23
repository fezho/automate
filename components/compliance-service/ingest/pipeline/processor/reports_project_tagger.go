package processor

import (
	"context"

	"github.com/chef/automate/api/interservice/authz"
	"github.com/chef/automate/components/compliance-service/ingest/pipeline/message"
	"github.com/chef/automate/components/compliance-service/reporting/relaxting"
	"github.com/chef/automate/lib/stringutils"
	"github.com/sirupsen/logrus"
)

// BundleReportProjectTagger - Build a project tagger processor for InSpec reports
func BundleReportProjectTagger(authzClient authz.ProjectsClient) message.CompliancePipe {
	return func(in <-chan message.Compliance) <-chan message.Compliance {
		return reportProjectTagger(in, authzClient)
	}
}

// This processor is bundling all the messages that are currently in the 'in' channel. The bundling of messages
// decreases the number of times the authz-service is called for project rules. The way it works is when a message
// comes in, we make a call to the authz-service for the rules. We use these rules for all the messages that are
// currently in the queue. The 'bundleSize' is the number of messages that can use the current project rules from authz.
// When the bundleSize zero or less we need to refetch the project rules.
func reportProjectTagger(in <-chan message.Compliance, authzClient authz.ProjectsClient) <-chan message.Compliance {
	if authzClient == nil {
		logrus.Error("no authz client found for project tagging; skipping project tagging")
		return in
	}
	out := make(chan message.Compliance, 100)
	go func() {
		bundleSize := 0
		var projectRulesCollection map[string]*authz.ProjectRules
		for msg := range in {
			if bundleSize <= 0 {
				bundleSize = len(in)
				logrus.WithFields(logrus.Fields{
					"message_id": msg.Report.ReportUuid,
					"bundleSize": bundleSize,
				}).Debug("BundleProjectTagging - Update Project rules")
				projectRulesCollection = getProjectRulesFromAuthz(msg.Ctx, authzClient)
			} else {
				// Skip
				bundleSize--
			}

			projectTags := findMatchingProjects(msg.InspecReport, projectRulesCollection)

			msg.InspecReport.Projects = projectTags
			msg.InspecSummary.Projects = projectTags

			out <- msg
		}
		close(out)
	}()

	return out
}

func getProjectRulesFromAuthz(ctx context.Context, authzClient authz.ProjectsClient) map[string]*authz.ProjectRules {
	projectsCollection, err := authzClient.ListRulesForAllProjects(ctx, &authz.ListRulesForAllProjectsReq{})

	if err != nil {
		// If there is an error getting the project rules from authz crash the service.
		logrus.WithError(err).Fatal("Could not fetch project rules from authz")
	}
	return projectsCollection.ProjectRules
}

func findMatchingProjects(report *relaxting.ESInSpecReport, projects map[string]*authz.ProjectRules) []string {
	matchingProjects := make([]string, 0)

	for projectName, projectRules := range projects {
		if reportMatchesRules(report, projectRules.Rules) {
			matchingProjects = append(matchingProjects, projectName)
		}
	}

	return matchingProjects
}

// Only one rule has to match for the entire project to match (ORed together).
func reportMatchesRules(report *relaxting.ESInSpecReport, rules []*authz.ProjectRule) bool {
	for _, rule := range rules {
		if rule.Type == authz.ProjectRuleTypes_NODE && reportMatchesAllConditions(report, rule.Conditions) {
			return true
		}
	}

	return false
}

// All the conditions must be true for a rule to be true (ANDed together).
// If there are no conditions then the rule is false
func reportMatchesAllConditions(report *relaxting.ESInSpecReport, conditions []*authz.Condition) bool {
	if len(conditions) == 0 {
		return false
	}

	for _, condition := range conditions {
		switch condition.Attribute {
		case authz.ProjectRuleConditionAttributes_ENVIRONMENT:
			if !stringutils.SliceContains(condition.Values, report.Environment) {
				return false
			}
		case authz.ProjectRuleConditionAttributes_CHEF_ROLE:
			foundMatch := false
			for _, projectRole := range condition.Values {
				if stringutils.SliceContains(report.Roles, projectRole) {
					foundMatch = true
					break
				}
			}
			if !foundMatch {
				return false
			}
		case authz.ProjectRuleConditionAttributes_CHEF_SERVER:
			if !stringutils.SliceContains(condition.Values, report.SourceFQDN) {
				return false
			}
		case authz.ProjectRuleConditionAttributes_CHEF_ORGANIZATION:
			if !stringutils.SliceContains(condition.Values, report.OrganizationName) {
				return false
			}
		case authz.ProjectRuleConditionAttributes_CHEF_POLICY_GROUP:
			if !stringutils.SliceContains(condition.Values, report.PolicyGroup) {
				return false
			}
		case authz.ProjectRuleConditionAttributes_CHEF_POLICY_NAME:
			if !stringutils.SliceContains(condition.Values, report.PolicyName) {
				return false
			}
		case authz.ProjectRuleConditionAttributes_CHEF_TAG:
			foundMatch := false
			for _, projectChefTag := range condition.Values {
				if stringutils.SliceContains(report.ChefTags, projectChefTag) {
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
