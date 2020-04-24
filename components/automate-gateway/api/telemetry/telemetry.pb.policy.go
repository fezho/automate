// Code generated by protoc-gen-policy. DO NOT EDIT.
// source: components/automate-gateway/api/telemetry/telemetry.proto

package telemetry

import policy "github.com/chef/automate/components/automate-gateway/api/iam/v2/policy"

func init() {
	policy.MapMethodTo("/chef.automate.api.telemetry.Telemetry/GetTelemetryConfiguration", "system:config", "system:telemetryConfig:get", "GET", "/api/v0/telemetry/config", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
}
