// Code generated by protoc-gen-policy. DO NOT EDIT.
// source: components/automate-gateway/api/nodes/nodes.proto

package nodes

import policyv2 "github.com/chef/automate/components/automate-gateway/authz/policy_v2"

func init() {
	policyv2.MapMethodTo("/chef.automate.api.nodes.v1.NodesService/Create", "infra:nodes", "infra:nodes:create", "POST", "/nodes", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*Node); ok {
			return policyv2.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "name":
					return m.Name
				case "platform":
					return m.Platform
				case "platform_version":
					return m.PlatformVersion
				case "manager":
					return m.Manager
				case "status":
					return m.Status
				case "connection_error":
					return m.ConnectionError
				case "state":
					return m.State
				case "name_prefix":
					return m.NamePrefix
				default:
					return ""
				}
			})
		}
		return ""
	})
	policyv2.MapMethodTo("/chef.automate.api.nodes.v1.NodesService/Read", "infra:nodes:{id}", "infra:nodes:get", "GET", "/nodes/id/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*Id); ok {
			return policyv2.ExpandParameterizedResource(unexpandedResource, func(want string) string {
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
	policyv2.MapMethodTo("/chef.automate.api.nodes.v1.NodesService/Update", "infra:nodes:{id}", "infra:nodes:update", "PUT", "/nodes/id/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*Node); ok {
			return policyv2.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "name":
					return m.Name
				case "platform":
					return m.Platform
				case "platform_version":
					return m.PlatformVersion
				case "manager":
					return m.Manager
				case "status":
					return m.Status
				case "connection_error":
					return m.ConnectionError
				case "state":
					return m.State
				case "name_prefix":
					return m.NamePrefix
				default:
					return ""
				}
			})
		}
		return ""
	})
	policyv2.MapMethodTo("/chef.automate.api.nodes.v1.NodesService/Delete", "infra:nodes:{id}", "infra:nodes:delete", "DELETE", "/nodes/id/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*Id); ok {
			return policyv2.ExpandParameterizedResource(unexpandedResource, func(want string) string {
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
	policyv2.MapMethodTo("/chef.automate.api.nodes.v1.NodesService/BulkDeleteById", "infra:nodes", "infra:nodes:delete", "POST", "/nodes/delete/ids", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
	policyv2.MapMethodTo("/chef.automate.api.nodes.v1.NodesService/List", "infra:nodes", "infra:nodes:list", "POST", "/nodes/search", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*Query); ok {
			return policyv2.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "sort":
					return m.Sort
				default:
					return ""
				}
			})
		}
		return ""
	})
	policyv2.MapMethodTo("/chef.automate.api.nodes.v1.NodesService/Rerun", "infra:nodes:{id}", "infra:nodes:rerun", "GET", "/nodes/rerun/id/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*Id); ok {
			return policyv2.ExpandParameterizedResource(unexpandedResource, func(want string) string {
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
	policyv2.MapMethodTo("/chef.automate.api.nodes.v1.NodesService/BulkDelete", "infra:nodes", "infra:nodes:delete", "POST", "/nodes/delete", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*Query); ok {
			return policyv2.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "sort":
					return m.Sort
				default:
					return ""
				}
			})
		}
		return ""
	})
	policyv2.MapMethodTo("/chef.automate.api.nodes.v1.NodesService/BulkCreate", "infra:nodes", "infra:nodes:create", "POST", "/nodes/bulk-create", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
}
