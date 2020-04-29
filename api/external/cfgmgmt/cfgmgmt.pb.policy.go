// Code generated by protoc-gen-policy. DO NOT EDIT.
// source: api/external/cfgmgmt/cfgmgmt.proto

package cfgmgmt

import (
	request "github.com/chef/automate/api/external/cfgmgmt/request"
	query "github.com/chef/automate/api/external/common/query"
	policy "github.com/chef/automate/components/automate-gateway/api/iam/v2/policy"
)

func init() {
	policy.MapMethodTo("/chef.automate.api.cfgmgmt.ConfigMgmt/GetNodes", "infra:nodes", "infra:nodes:list", "GET", "/api/v0/cfgmgmt/nodes", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.Nodes); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "start":
					return m.Start
				case "end":
					return m.End
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.cfgmgmt.ConfigMgmt/GetRuns", "infra:nodes:{node_id}", "infra:nodes:list", "GET", "/api/v0/cfgmgmt/nodes/{node_id}/runs", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.Runs); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "node_id":
					return m.NodeId
				case "start":
					return m.Start
				case "end":
					return m.End
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.cfgmgmt.ConfigMgmt/GetNodesCounts", "infra:nodes", "infra:nodes:list", "GET", "/api/v0/cfgmgmt/stats/node_counts", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
	policy.MapMethodTo("/chef.automate.api.cfgmgmt.ConfigMgmt/GetRunsCounts", "infra:nodes", "infra:nodes:list", "GET", "/api/v0/cfgmgmt/stats/run_counts", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.RunsCounts); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "start":
					return m.Start
				case "end":
					return m.End
				case "node_id":
					return m.NodeId
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.cfgmgmt.ConfigMgmt/GetCheckInCountsTimeSeries", "infra:nodes", "infra:nodes:list", "GET", "/api/v0/cfgmgmt/stats/checkin_counts_timeseries", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
	policy.MapMethodTo("/chef.automate.api.cfgmgmt.ConfigMgmt/GetMissingNodeDurationCounts", "infra:nodes", "infra:nodes:list", "GET", "/api/v0/cfgmgmt/stats/missing_node_duration_counts", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
	policy.MapMethodTo("/chef.automate.api.cfgmgmt.ConfigMgmt/GetNodeRun", "infra:nodes:{node_id}", "infra:nodes:get", "GET", "/api/v0/cfgmgmt/nodes/{node_id}/runs/{run_id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.NodeRun); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "node_id":
					return m.NodeId
				case "run_id":
					return m.RunId
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.cfgmgmt.ConfigMgmt/GetSuggestions", "infra:nodes", "infra:nodes:list", "GET", "/api/v0/cfgmgmt/suggestions", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*query.Suggestion); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "type":
					return m.Type
				case "text":
					return m.Text
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.cfgmgmt.ConfigMgmt/GetOrganizations", "infra:nodes", "infra:nodes:list", "GET", "/api/v0/cfgmgmt/organizations", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
	policy.MapMethodTo("/chef.automate.api.cfgmgmt.ConfigMgmt/GetSourceFqdns", "infra:nodes", "infra:nodes:list", "GET", "/api/v0/cfgmgmt/source_fqdns", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
	policy.MapMethodTo("/chef.automate.api.cfgmgmt.ConfigMgmt/GetAttributes", "infra:nodes:{node_id}", "infra:nodes:get", "GET", "/api/v0/cfgmgmt/nodes/{node_id}/attribute", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.Node); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "node_id":
					return m.NodeId
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.cfgmgmt.ConfigMgmt/GetVersion", "system:service:version", "system:serviceVersion:get", "GET", "/api/v0/cfgmgmt/version", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
	policy.MapMethodTo("/chef.automate.api.cfgmgmt.ConfigMgmt/GetPolicyCookbooks", "infra:nodes:{revision_id}", "infra:nodes:list", "GET", "/api/v0/cfgmgmt/policy_revision/{revision_id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.PolicyRevision); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "revision_id":
					return m.RevisionId
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.cfgmgmt.ConfigMgmt/GetErrors", "infra:nodes", "infra:nodes:list", "GET", "/api/v0/cfgmgmt/errors", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
	policy.MapMethodTo("/chef.automate.api.cfgmgmt.ConfigMgmt/GetNodeMetadataCounts", "infra:nodes", "infra:nodes:list", "GET", "/api/v0/cfgmgmt/node_metadata_counts", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.NodeMetadataCounts); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "start":
					return m.Start
				case "end":
					return m.End
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.cfgmgmt.ConfigMgmt/GetNodeRunsDailyStatusTimeSeries", "infra:nodes", "infra:nodes:list", "GET", "/api/v0/cfgmgmt/node_runs_daily_status_time_series", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.NodeRunsDailyStatusTimeSeries); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "node_id":
					return m.NodeId
				default:
					return ""
				}
			})
		}
		return ""
	})
}
