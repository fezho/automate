// Code generated by protoc-gen-policy. DO NOT EDIT.
// source: components/automate-gateway/api/iam/v2/rules.proto

package v2

import (
	policy "github.com/chef/automate/components/automate-gateway/api/iam/v2/policy"
	request "github.com/chef/automate/components/automate-gateway/api/iam/v2/request"
)

func init() {
	policy.MapMethodTo("/chef.automate.api.iam.v2.Rules/CreateRule", "iam:projects:{project_id}", "iam:projects:update", "POST", "/iam/v2/projects/{project_id}/rules", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.CreateRuleReq); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "project_id":
					return m.ProjectId
				case "name":
					return m.Name
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.iam.v2.Rules/UpdateRule", "iam:projects:{project_id}", "iam:projects:update", "PUT", "/iam/v2/projects/{project_id}/rules/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.UpdateRuleReq); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "project_id":
					return m.ProjectId
				case "name":
					return m.Name
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.iam.v2.Rules/GetRule", "iam:projects:{project_id}", "iam:projects:get", "GET", "/iam/v2/projects/{project_id}/rules/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.GetRuleReq); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "project_id":
					return m.ProjectId
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.iam.v2.Rules/ListRulesForProject", "iam:projects:{id}", "iam:projects:get", "GET", "/iam/v2/projects/{id}/rules", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.ListRulesForProjectReq); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.iam.v2.Rules/DeleteRule", "iam:projects:{project_id}", "iam:projects:update", "DELETE", "/iam/v2/projects/{project_id}/rules/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.DeleteRuleReq); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "project_id":
					return m.ProjectId
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.iam.v2.Rules/ApplyRulesStart", "iam:rules", "iam:rules:apply", "POST", "/iam/v2/apply-rules", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
	policy.MapMethodTo("/chef.automate.api.iam.v2.Rules/ApplyRulesCancel", "iam:rules", "iam:rules:cancel", "DELETE", "/iam/v2/apply-rules", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
	policy.MapMethodTo("/chef.automate.api.iam.v2.Rules/ApplyRulesStatus", "iam:rules", "iam:rules:status", "GET", "/iam/v2/apply-rules", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
}
