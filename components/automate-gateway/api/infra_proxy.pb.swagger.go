package api

func init() {
	Swagger.Add("infra_proxy", `{
  "swagger": "2.0",
  "info": {
    "title": "api/external/infra_proxy/infra_proxy.proto",
    "version": "version not set"
  },
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/infra/servers": {
      "get": {
        "operationId": "GetServers",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.GetServers"
            }
          }
        },
        "tags": [
          "InfraProxy"
        ]
      },
      "post": {
        "operationId": "CreateServer",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.CreateServer"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.request.CreateServer"
            }
          }
        ],
        "tags": [
          "InfraProxy"
        ]
      }
    },
    "/infra/servers/{id}": {
      "get": {
        "operationId": "GetServer",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.GetServer"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "description": "Chef Infra Server ID.",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "InfraProxy"
        ]
      },
      "delete": {
        "operationId": "DeleteServer",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.DeleteServer"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "description": "Chef Infra Server ID.",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "InfraProxy"
        ]
      },
      "put": {
        "operationId": "UpdateServer",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.UpdateServer"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "description": "Chef Infra Server ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.request.UpdateServer"
            }
          }
        ],
        "tags": [
          "InfraProxy"
        ]
      }
    },
    "/infra/servers/{server_id}/orgs": {
      "get": {
        "operationId": "GetOrgs",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.GetOrgs"
            }
          }
        },
        "parameters": [
          {
            "name": "server_id",
            "description": "Chef Infra Server ID.",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "InfraProxy"
        ]
      },
      "post": {
        "operationId": "CreateOrg",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.CreateOrg"
            }
          }
        },
        "parameters": [
          {
            "name": "server_id",
            "description": "Chef Infra Server ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.request.CreateOrg"
            }
          }
        ],
        "tags": [
          "InfraProxy"
        ]
      }
    },
    "/infra/servers/{server_id}/orgs/{id}": {
      "get": {
        "operationId": "GetOrg",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.GetOrg"
            }
          }
        },
        "parameters": [
          {
            "name": "server_id",
            "description": "Chef Infra Server ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "id",
            "description": "Chef organization ID.",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "InfraProxy"
        ]
      },
      "delete": {
        "operationId": "DeleteOrg",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.DeleteOrg"
            }
          }
        },
        "parameters": [
          {
            "name": "server_id",
            "description": "Chef Infra Server ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "id",
            "description": "Chef organization ID.",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "InfraProxy"
        ]
      },
      "put": {
        "operationId": "UpdateOrg",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.UpdateOrg"
            }
          }
        },
        "parameters": [
          {
            "name": "server_id",
            "description": "Chef Infra Server ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "id",
            "description": "Chef organization ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.request.UpdateOrg"
            }
          }
        ],
        "tags": [
          "InfraProxy"
        ]
      }
    },
    "/infra/servers/{server_id}/orgs/{org_id}/affected-nodes/{chef_type}/{name}": {
      "get": {
        "operationId": "GetAffectedNodes",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.AffectedNodes"
            }
          }
        },
        "parameters": [
          {
            "name": "server_id",
            "description": "ID of the Server.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "org_id",
            "description": "ID of the Org.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "chef_type",
            "description": "Type of the chef object (e.g. 'cookbooks', 'roles', 'chef_environment').",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "name",
            "description": "Name of the chef object.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "version",
            "description": "Version of the chef object.",
            "in": "query",
            "required": false,
            "type": "string"
          }
        ],
        "tags": [
          "InfraProxy"
        ]
      }
    },
    "/infra/servers/{server_id}/orgs/{org_id}/clients": {
      "get": {
        "operationId": "GetClients",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.Clients"
            }
          }
        },
        "parameters": [
          {
            "name": "server_id",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "org_id",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "InfraProxy"
        ]
      }
    },
    "/infra/servers/{server_id}/orgs/{org_id}/clients/{name}": {
      "get": {
        "operationId": "GetClient",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.Client"
            }
          }
        },
        "parameters": [
          {
            "name": "server_id",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "org_id",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "name",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "InfraProxy"
        ]
      }
    },
    "/infra/servers/{server_id}/orgs/{org_id}/cookbooks": {
      "get": {
        "operationId": "GetCookbooks",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.Cookbooks"
            }
          }
        },
        "parameters": [
          {
            "name": "server_id",
            "description": "ID of the Server",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "org_id",
            "description": "ID of the Org.",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "InfraProxy"
        ]
      }
    },
    "/infra/servers/{server_id}/orgs/{org_id}/cookbooks/{name}": {
      "get": {
        "operationId": "GetCookbookVersions",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.CookbookVersions"
            }
          }
        },
        "parameters": [
          {
            "name": "server_id",
            "description": "ID of the Server.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "org_id",
            "description": "ID of the Org.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "name",
            "description": "Name of the cookbook.",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "InfraProxy"
        ]
      }
    },
    "/infra/servers/{server_id}/orgs/{org_id}/cookbooks/{name}/{version}": {
      "get": {
        "operationId": "GetCookbook",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.Cookbook"
            }
          }
        },
        "parameters": [
          {
            "name": "server_id",
            "description": "ID of the Server.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "org_id",
            "description": "ID of the Org.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "name",
            "description": "Name of the cookbook.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "version",
            "description": "Version of the cookbook.",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "InfraProxy"
        ]
      }
    },
    "/infra/servers/{server_id}/orgs/{org_id}/cookbooks/{name}/{version}/file-content": {
      "get": {
        "operationId": "GetCookbookFileContent",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.CookbookFileContent"
            }
          }
        },
        "parameters": [
          {
            "name": "server_id",
            "description": "ID of the server.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "org_id",
            "description": "ID of the org.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "name",
            "description": "Name of the cookbook.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "version",
            "description": "Version of the cookbook.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "url",
            "description": "Cookbook data file URL.",
            "in": "query",
            "required": false,
            "type": "string"
          }
        ],
        "tags": [
          "InfraProxy"
        ]
      }
    },
    "/infra/servers/{server_id}/orgs/{org_id}/data_bags/{name}": {
      "get": {
        "operationId": "GetDataBags",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.DataBags"
            }
          }
        },
        "parameters": [
          {
            "name": "server_id",
            "description": "Chef Infra Server ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "org_id",
            "description": "Chef organization ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "name",
            "description": "Data bag name.",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "InfraProxy"
        ]
      }
    },
    "/infra/servers/{server_id}/orgs/{org_id}/data_bags/{name}/{item}": {
      "get": {
        "operationId": "GetDataBagItem",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.DataBag"
            }
          }
        },
        "parameters": [
          {
            "name": "server_id",
            "description": "Chef Infra Server ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "org_id",
            "description": "Chef organization ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "name",
            "description": "Data bag name.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "item",
            "description": "Data bag item name.",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "InfraProxy"
        ]
      }
    },
    "/infra/servers/{server_id}/orgs/{org_id}/environments": {
      "get": {
        "operationId": "GetEnvironments",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.Environments"
            }
          }
        },
        "parameters": [
          {
            "name": "server_id",
            "description": "Chef Infra Server ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "org_id",
            "description": "Chef organization ID.",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "InfraProxy"
        ]
      }
    },
    "/infra/servers/{server_id}/orgs/{org_id}/environments/{name}": {
      "get": {
        "operationId": "GetEnvironment",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.Environment"
            }
          }
        },
        "parameters": [
          {
            "name": "server_id",
            "description": "Chef Infra Server ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "org_id",
            "description": "Chef organization ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "name",
            "description": "Environment name.",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "InfraProxy"
        ]
      }
    },
    "/infra/servers/{server_id}/orgs/{org_id}/policyfiles": {
      "get": {
        "operationId": "GetPolicyfiles",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.Policyfiles"
            }
          }
        },
        "parameters": [
          {
            "name": "server_id",
            "description": "Chef Infra Server ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "org_id",
            "description": "Chef Organization ID.",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "InfraProxy"
        ]
      }
    },
    "/infra/servers/{server_id}/orgs/{org_id}/policyfiles/{name}": {
      "get": {
        "operationId": "GetPolicyfile",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.Policyfile"
            }
          }
        },
        "parameters": [
          {
            "name": "server_id",
            "description": "Chef Infra Server ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "org_id",
            "description": "Chef Organization ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "name",
            "description": "Policyfile name.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "revision_id",
            "description": "Policyfile revision ID.",
            "in": "query",
            "required": false,
            "type": "string"
          }
        ],
        "tags": [
          "InfraProxy"
        ]
      }
    },
    "/infra/servers/{server_id}/orgs/{org_id}/roles": {
      "get": {
        "operationId": "GetRoles",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.Roles"
            }
          }
        },
        "parameters": [
          {
            "name": "server_id",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "org_id",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "InfraProxy"
        ]
      }
    },
    "/infra/servers/{server_id}/orgs/{org_id}/roles/{name}": {
      "get": {
        "operationId": "GetRole",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.Role"
            }
          }
        },
        "parameters": [
          {
            "name": "server_id",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "org_id",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "name",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "InfraProxy"
        ]
      }
    },
    "/infra/version": {
      "get": {
        "operationId": "GetVersion",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.common.version.VersionInfo"
            }
          }
        },
        "tags": [
          "InfraProxy"
        ]
      }
    }
  },
  "definitions": {
    "chef.automate.api.common.version.VersionInfo": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        },
        "version": {
          "type": "string"
        },
        "sha": {
          "type": "string"
        },
        "built": {
          "type": "string"
        }
      }
    },
    "chef.automate.api.infra_proxy.request.CreateOrg": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "description": "Chef organization ID."
        },
        "name": {
          "type": "string",
          "description": "Chef organization name."
        },
        "admin_user": {
          "type": "string",
          "description": "Chef organization admin user."
        },
        "admin_key": {
          "type": "string",
          "description": "Chef organization admin key."
        },
        "server_id": {
          "type": "string",
          "description": "Chef Infra Server ID."
        },
        "projects": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "List of projects this chef organization belongs to. May be empty."
        }
      }
    },
    "chef.automate.api.infra_proxy.request.CreateServer": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "description": "Chef Infra Server ID."
        },
        "name": {
          "type": "string",
          "description": "Chef Infra Server name."
        },
        "fqdn": {
          "type": "string",
          "description": "Chef Infra Server FQDN."
        },
        "ip_address": {
          "type": "string",
          "description": "Chef infra Server IP address."
        }
      }
    },
    "chef.automate.api.infra_proxy.request.UpdateOrg": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "description": "Chef organization ID."
        },
        "name": {
          "type": "string",
          "description": "Chef organization name."
        },
        "admin_user": {
          "type": "string",
          "description": "Chef organization admin user."
        },
        "admin_key": {
          "type": "string",
          "description": "Chef organization admin key."
        },
        "server_id": {
          "type": "string",
          "description": "Chef Infra Server ID."
        },
        "projects": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "List of projects this chef organization belongs to. May be empty."
        }
      }
    },
    "chef.automate.api.infra_proxy.request.UpdateServer": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "description": "Chef Infra Server ID."
        },
        "name": {
          "type": "string",
          "description": "Chef Infra Server name."
        },
        "fqdn": {
          "type": "string",
          "description": "Chef Infra Server FQDN."
        },
        "ip_address": {
          "type": "string",
          "description": "Chef Infra Server IP address."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.AffectedNodes": {
      "type": "object",
      "properties": {
        "nodes": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.infra_proxy.response.NodeAttribute"
          },
          "description": "List of the nodes which are affected by the chef object."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.Client": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        },
        "client_name": {
          "type": "string"
        },
        "org_name": {
          "type": "string"
        },
        "admin": {
          "type": "boolean",
          "format": "boolean"
        },
        "validator": {
          "type": "boolean",
          "format": "boolean"
        },
        "certificate": {
          "type": "string"
        },
        "public_key": {
          "type": "string"
        },
        "private_key": {
          "type": "string"
        },
        "uri": {
          "type": "string"
        },
        "json_class": {
          "type": "string"
        },
        "chef_type": {
          "type": "string"
        }
      }
    },
    "chef.automate.api.infra_proxy.response.ClientListItem": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        }
      }
    },
    "chef.automate.api.infra_proxy.response.Clients": {
      "type": "object",
      "properties": {
        "clients": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.infra_proxy.response.ClientListItem"
          }
        }
      }
    },
    "chef.automate.api.infra_proxy.response.Cookbook": {
      "type": "object",
      "properties": {
        "cookbook_name": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "version": {
          "type": "string"
        },
        "chef_type": {
          "type": "string"
        },
        "frozen": {
          "type": "boolean",
          "format": "boolean"
        },
        "json_class": {
          "type": "string"
        },
        "files": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.infra_proxy.response.CookbookItem"
          }
        },
        "templates": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.infra_proxy.response.CookbookItem"
          }
        },
        "attributes": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.infra_proxy.response.CookbookItem"
          }
        },
        "recipes": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.infra_proxy.response.CookbookItem"
          }
        },
        "definitions": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.infra_proxy.response.CookbookItem"
          }
        },
        "libraries": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.infra_proxy.response.CookbookItem"
          }
        },
        "providers": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.infra_proxy.response.CookbookItem"
          }
        },
        "resources": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.infra_proxy.response.CookbookItem"
          }
        },
        "root_files": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.infra_proxy.response.CookbookItem"
          }
        },
        "metadata": {
          "$ref": "#/definitions/chef.automate.api.infra_proxy.response.CookbookMeta"
        },
        "access": {
          "$ref": "#/definitions/chef.automate.api.infra_proxy.response.CookbookAccess"
        }
      }
    },
    "chef.automate.api.infra_proxy.response.CookbookAccess": {
      "type": "object",
      "properties": {
        "read": {
          "type": "boolean",
          "format": "boolean"
        },
        "create": {
          "type": "boolean",
          "format": "boolean"
        },
        "grant": {
          "type": "boolean",
          "format": "boolean"
        },
        "update": {
          "type": "boolean",
          "format": "boolean"
        },
        "delete": {
          "type": "boolean",
          "format": "boolean"
        }
      }
    },
    "chef.automate.api.infra_proxy.response.CookbookFileContent": {
      "type": "object",
      "properties": {
        "content": {
          "type": "string",
          "description": "Cookbook data file content."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.CookbookItem": {
      "type": "object",
      "properties": {
        "url": {
          "type": "string"
        },
        "path": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "checksum": {
          "type": "string"
        },
        "specificity": {
          "type": "string"
        }
      }
    },
    "chef.automate.api.infra_proxy.response.CookbookLock": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "description": "Cookbook name."
        },
        "version": {
          "type": "string",
          "description": "Cookbook version."
        },
        "identifier": {
          "type": "string",
          "description": "Cookbook identifier."
        },
        "dotted_identifier": {
          "type": "string",
          "description": "Cookbook decimal number identifier."
        },
        "source": {
          "type": "string",
          "description": "Cookbook source."
        },
        "cache_key": {
          "type": "string",
          "description": "Cookbook cache key."
        },
        "SCMDetail": {
          "$ref": "#/definitions/chef.automate.api.infra_proxy.response.SCMDetail",
          "description": "SCM detail."
        },
        "source_options": {
          "$ref": "#/definitions/chef.automate.api.infra_proxy.response.SourceOptions",
          "description": "Cookbook source path."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.CookbookMeta": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        },
        "version": {
          "type": "string"
        },
        "description": {
          "type": "string"
        },
        "long_description": {
          "type": "string"
        },
        "maintainer": {
          "type": "string"
        },
        "maintainer_email": {
          "type": "string"
        },
        "license": {
          "type": "string"
        }
      }
    },
    "chef.automate.api.infra_proxy.response.CookbookVersion": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "description": "Name of the cookbook."
        },
        "version": {
          "type": "string",
          "description": "Version of the cookbook."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.CookbookVersions": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "description": "Name of the cookbook."
        },
        "versions": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "List of all versions avaiable for cookbook."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.Cookbooks": {
      "type": "object",
      "properties": {
        "cookbooks": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.infra_proxy.response.CookbookVersion"
          },
          "description": "List of cookbooks with name and version."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.CreateOrg": {
      "type": "object",
      "properties": {
        "org": {
          "$ref": "#/definitions/chef.automate.api.infra_proxy.response.Org",
          "description": "Chef organization."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.CreateServer": {
      "type": "object",
      "properties": {
        "server": {
          "$ref": "#/definitions/chef.automate.api.infra_proxy.response.Server",
          "description": "Chef Infra Server."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.DataBag": {
      "type": "object",
      "properties": {
        "data": {
          "type": "string",
          "description": "Stringified json of data bag item."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.DataBagListItem": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "description": "Data bag item name."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.DataBags": {
      "type": "object",
      "properties": {
        "data_bags": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.infra_proxy.response.DataBagListItem"
          },
          "description": "Data bags item list."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.DeleteOrg": {
      "type": "object",
      "properties": {
        "org": {
          "$ref": "#/definitions/chef.automate.api.infra_proxy.response.Org",
          "description": "Chef organization."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.DeleteServer": {
      "type": "object",
      "properties": {
        "server": {
          "$ref": "#/definitions/chef.automate.api.infra_proxy.response.Server",
          "description": "Chef Infra Server."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.Environment": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "description": "Environment name."
        },
        "chef_type": {
          "type": "string",
          "description": "Object type."
        },
        "description": {
          "type": "string",
          "description": "Node description."
        },
        "json_class": {
          "type": "string",
          "description": "Class name."
        },
        "cookbook_versions": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "Versioned cookbooks list."
        },
        "default_attributes": {
          "type": "string",
          "description": "Stringified json of the default attributes."
        },
        "override_attributes": {
          "type": "string",
          "description": "Stringified json of the override attributes."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.EnvironmentListItem": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "description": "Environment name."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.Environments": {
      "type": "object",
      "properties": {
        "environments": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.infra_proxy.response.EnvironmentListItem"
          },
          "description": "Environments list."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.ExpandedRunList": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "description": "ID of the run list collection."
        },
        "run_list": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.infra_proxy.response.RunList"
          },
          "description": "List of the run list."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.GetOrg": {
      "type": "object",
      "properties": {
        "org": {
          "$ref": "#/definitions/chef.automate.api.infra_proxy.response.Org",
          "description": "Chef organization."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.GetOrgs": {
      "type": "object",
      "properties": {
        "orgs": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.infra_proxy.response.OrgListItem"
          },
          "description": "Chef organization list."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.GetServer": {
      "type": "object",
      "properties": {
        "server": {
          "$ref": "#/definitions/chef.automate.api.infra_proxy.response.Server",
          "description": "Chef Infra Server."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.GetServers": {
      "type": "object",
      "properties": {
        "servers": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.infra_proxy.response.Server"
          },
          "description": "Chef Infra Servers."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.IncludedPolicyLock": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "description": "Included Policyfile name."
        },
        "revision_id": {
          "type": "string",
          "description": "Policyfile revision ID."
        },
        "source_options": {
          "$ref": "#/definitions/chef.automate.api.infra_proxy.response.SourceOptions",
          "description": "Included policyfile source options."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.NamedRunList": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "description": "Run list name."
        },
        "run_list": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "Run list associated with the policy."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.NodeAttribute": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "description": "ID of the node."
        },
        "name": {
          "type": "string",
          "description": "Name of the node."
        },
        "check_in": {
          "type": "string",
          "description": "Last checked in of the node."
        },
        "uptime": {
          "type": "string",
          "description": "Uptime of the node."
        },
        "platform": {
          "type": "string",
          "description": "Name of the platform of the node."
        },
        "environment": {
          "type": "string",
          "description": "Environment name of the node."
        },
        "policy_group": {
          "type": "string",
          "description": "Policy group name of the node."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.Org": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "description": "Chef organization ID."
        },
        "name": {
          "type": "string",
          "description": "Chef organization name."
        },
        "admin_user": {
          "type": "string",
          "description": "Chef organization admin user."
        },
        "credential_id": {
          "type": "string",
          "description": "Chef organization credential ID."
        },
        "server_id": {
          "type": "string",
          "description": "Chef Infra Server ID."
        },
        "projects": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "List of projects this chef organization belongs to. May be empty."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.OrgListItem": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "description": "Chef organization ID."
        },
        "name": {
          "type": "string",
          "description": "Chef organization name."
        },
        "admin_user": {
          "type": "string",
          "description": "Chef organization admin user."
        },
        "server_id": {
          "type": "string",
          "description": "Chef Infra Server ID."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.Policyfile": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "description": "Policyfile name."
        },
        "policy_group": {
          "type": "string",
          "description": "Policy group name."
        },
        "revision_id": {
          "type": "string",
          "description": "Policy revision ID."
        },
        "run_list": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "Run-list associated with the policy."
        },
        "named_run_list": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.infra_proxy.response.NamedRunList"
          },
          "description": "Named run-list associated with the policy."
        },
        "included_policy_locks": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.infra_proxy.response.IncludedPolicyLock"
          },
          "description": "Included policy locks files."
        },
        "cookbook_locks": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.infra_proxy.response.CookbookLock"
          },
          "description": "List of cookbook locks under this policy."
        },
        "default_attributes": {
          "type": "string",
          "description": "Stringified JSON of the default attributes."
        },
        "override_attributes": {
          "type": "string",
          "description": "Stringified JSON of the override attributes."
        },
        "expanded_run_list": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.infra_proxy.response.ExpandedRunList"
          },
          "description": "Expanded run-list associated with the policy."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.PolicyfileListItem": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "description": "Policyfile name."
        },
        "revision_id": {
          "type": "string",
          "description": "Policyfile Revision ID."
        },
        "policy_group": {
          "type": "string",
          "description": "Policyfile policy group."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.Policyfiles": {
      "type": "object",
      "properties": {
        "policies": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.infra_proxy.response.PolicyfileListItem"
          },
          "description": "Policyfiles list."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.Role": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "description": "Name of the role."
        },
        "chef_type": {
          "type": "string",
          "description": "Type of the chef object."
        },
        "description": {
          "type": "string",
          "description": "Descrption of the role."
        },
        "default_attributes": {
          "type": "string",
          "description": "Stringified json of the default attributes."
        },
        "override_attributes": {
          "type": "string",
          "description": "Stringified json of the override attributes."
        },
        "json_class": {
          "type": "string",
          "description": "Json class name."
        },
        "run_list": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "Run list for the role."
        },
        "expanded_run_list": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.infra_proxy.response.ExpandedRunList"
          },
          "description": "List of expanded run list for the role."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.RoleListItem": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "description": "Name of the role."
        },
        "description": {
          "type": "string",
          "description": "Desscription of the role."
        },
        "environments": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "Environment for the role."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.Roles": {
      "type": "object",
      "properties": {
        "roles": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.infra_proxy.response.RoleListItem"
          },
          "description": "List of the roles item."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.RunList": {
      "type": "object",
      "properties": {
        "type": {
          "type": "string",
          "description": "Type of run list item (e.g. 'recipe')."
        },
        "name": {
          "type": "string",
          "description": "Name of run list item."
        },
        "version": {
          "type": "string",
          "description": "Version of run list item."
        },
        "skipped": {
          "type": "boolean",
          "format": "boolean",
          "description": "Boolean denoting whether or not the run list item was skipped."
        },
        "children": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.infra_proxy.response.RunList"
          },
          "description": "List of the run list."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.SCMDetail": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "description": "SCM name."
        },
        "remote": {
          "type": "string",
          "description": "SCM remote location."
        },
        "revision": {
          "type": "string",
          "description": "SCM revision detail."
        },
        "working_tree_clean": {
          "type": "boolean",
          "format": "boolean",
          "description": "Boolean that denotes if the working tree is clean or not."
        },
        "published": {
          "type": "boolean",
          "format": "boolean",
          "description": "Source's published information."
        },
        "synchronized_remote_branches": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "Synchronized remote branches list."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.Server": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "description": "Chef Infra Server ID."
        },
        "name": {
          "type": "string",
          "description": "Chef Infra Server name."
        },
        "fqdn": {
          "type": "string",
          "description": "Chef Infra Server FQDN."
        },
        "ip_address": {
          "type": "string",
          "description": "Chef Infra Server IP address."
        },
        "orgs_count": {
          "type": "integer",
          "format": "int32",
          "description": "Chef organizations count associated with Chef Infra Server."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.SourceOptions": {
      "type": "object",
      "properties": {
        "path": {
          "type": "string",
          "description": "Source options path."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.UpdateOrg": {
      "type": "object",
      "properties": {
        "org": {
          "$ref": "#/definitions/chef.automate.api.infra_proxy.response.Org",
          "description": "Chef organization."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.UpdateServer": {
      "type": "object",
      "properties": {
        "server": {
          "$ref": "#/definitions/chef.automate.api.infra_proxy.response.Server",
          "description": "Chef Infra Server."
        }
      }
    }
  }
}
`)
}
