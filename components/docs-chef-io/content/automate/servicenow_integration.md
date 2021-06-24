+++
title = "ServiceNow Integration"

draft = false

gh_repo = "automate"

[menu]
  [menu.automate]
    title = "ServiceNow Integration"
    parent = "automate/reference"
    identifier = "automate/reference/servicenow_integration.md ServiceNow Integration"
    weight = 40
+++
## Chef Automate Scoped Certified Application
This guide aids you to set up a customer’s instance for enabling the population of CMDB [Configuration Management Database] data from each node being managed by Chef Automate in the customer’s estate. It also includes configuration management status and compliance management status.
## ServiceNow Integration
The integration between a Chef Automate server and a ServiceNow instance requires the following:
- Chef Automate Scoped Application
- Chef Automate server
The Chef Automate application is a ServiceNow certified scoped application available from the ServiceNow store. The [Chef Automate server](https://www.chef.io/automate/) provides a full suite of enterprise capabilities for workflow, node visibility, and compliance, and integrates with the Chef Automate infrastructure and compliance functionality. This server sends HTTPS JSON data feeds to the Chef Automate application in a ServiceNow instance to create and update the _ServiceNow CMDB_ tables.
{{< figure src="/images/automate/snow_integration_dataflow_diagram.png" alt="Data Flow Diagram">}}
### Prerequisites
- The ServiceNow instance must be publicly reachable on https port 443.
- [Chef automate server installation](https://docs.chef.io/chef_automate.html).
- Service now package - System Import Sets com. `glide.system\_import\_set`, min version 1.0.0.
- Service now package - Configuration Management (CMDB) 1.1.
- Service now plugin - Configuration Management For Scoped Apps [com.snc.cmdb.scoped] 1.0.0.
## Installation
### Install Chef Automate application in ServiceNow
The Chef Automate application exposes the REST API endpoint that facilitates the communication between Chef Automate and the ServiceNow instance.
- Visit the ServiceNow store at <https://store.servicenow.com>.
- Get the Chef Automate application.
- In the ServiceNow instance, navigate to the **System Applications** > **Applications** menu.
- From the **Downloads** tab, install the **Chef Automate** application.
### Create Application Users
The application provides several roles appropriate for integration, which can be assigned to existing or new ServiceNow users. The roles are as follows:
- x_chef_automate.admin
- x_chef_automate.user
- x_chef_automate.api
These roles are part of the package. Users can create their requisite roles and controls if any further restrictions are required.
#### Role x_chef_automate.admin
This role can be assigned to a user other than a Systems administrator to allow administration of the application properties and logs. Thus, administration can be carried out by a user who is not the system administrator. 
Note that a systems administrator can perform all tasks that an `x_chef_automate.admin` role can. In addition, the Admin role is added by default in **ServiceNow** and is not part of the app.
The **admin** role grants a user access to the:
- Asset Imports module
- Properties
- Automate instances.
- Servers
- Chef infra client runs
- Compliance reports profiles
- Compliance report results
- Compliance Profiles
- Compliance Profiles results
- Support
- Events 
- Script includes
- Transform maps
- Chef servers
#### Role x_chef_automate.user
The role is suitable for users that require application access without administration access. The role grants a user access to the:
- Chef Automate menu
- Servers module
- Client runs module.
- CI cookbooks module
- Compliance reports module
- Compliance reports profiles module
- Compliance reports module
- Compliance profiles module
- Compliance profiles controls module
- Logs module
- Properties module
- Support module
Note that for integration with CMDB data in ServiceNow, you need to assign the _OOB ITIL_ role.
#### Role x_chef_automate.api
This role should be assigned to a user who is responsible for integrating the Chef automate data into the application. It is recommended that a new user is created specifically for this. This user’s credentials are required to configure the Chef automate server for communication with the application. Note that for integration with CMDB data you need to assign the _OOB ITIL_ role. You need to select Web service access only for this user.
#### References
- [Creating users in ServiceNow](https://docs.servicenow.com/bundle/madrid-platform-administration/page/administer/users-and-groups/task/t_CreateAUser.html)
- [Assigning roles in ServiceNow](https://docs.servicenow.com/bundle/madrid-platform-administration/page/administer/users-and-groups/task/t_AssignARoleToAUser.html)
### Application properties
The application properties can be configured by users with `admin` or `x_chef_automate.admin` roles. Select the **Chef Automate** > **Properties** menu item to navigate to the **Properties** configuration UI.
{{< figure src="/images/automate/snow_integration_appproperties.png" alt="Chef Automate Properties">}}

The application system properties are:
- x_chef_automate.Rest.api
- x_chef_automate.insert_manufacturer
- x_chef_automate.insert_model
- x_chef_automate.client_runs_per_node
- x_chef_automate.compliance_reports_per_node
- x_chef_automate.chef.user
- x_chef_automate.logging.enabled
- x_chef_automate.chef.default.status
- x_chef_automate.logging.verbosity
- X_chef_automate.discovery source 
#### Property  x_chef_automate.Rest.api
The `x_chef_automate.Rest.api` property enables the Chef Automate API from ServiceNow when Turn on REST API is set to **Yes** by default. The possible values are:
- Yes (default)
- No
#### Property x_chef_automate.insert_manufacturer
The `x_chef_automate.insert_manufacturer` property inserts the new record during the import if a model is not found in the _core\_company_ table by setting the property to **Yes**.
#### Property x_chef_automate.insert_model
The `x_chef_automate.insert_model` property inserts the new record during the import if a model is not found in the _cmdb\_model_ table by setting the property to **Yes**.
#### Property x_chef_automate.client_runs_per_node
These 2 properties are used to set a maximum number of clients runs and compliance reports for a node. The default value is **5**.
#### Property x_chef_automate.compliance_reports_per_node
These 2 properties are used to set a maximum number of clients runs and compliance reports for a node. The default value is **5**.
#### Property  x_chef_automate.chef.user
The `x_chef_automate.chef.user` property is used to set up user’s ID or sys ID for inserting the data.
#### Property x_chef_automate. logging.enabled
The `x_chef_automate. logging.enabled` property is used for flagging the enable or disable logging. The default value is **No**.
#### Property x_chef_automate.chef.default.status
The `x_chef_automate.chef.default. status` property is used to set up the status of the service record as **inserted** or **updated**.
#### Property x_chef_automate.logging.verbosity
The `x_chef_automate.logging.verbosity` property debugs the data in ServiceNow. The possible values are:
- Debug
- Warn
- Info
- Error
It enables the selected logging level and is visible in logs.
#### Property x_chef_automate.discovery.source
The `x_chef_automate.discovery.source` property sets the value for the discovery source. It takes import set when left blank by default.
### Scripted REST API
The _Scripted REST API_ is used to establish a connection between ServiceNow and Chef Automate with authentication.
{{< figure src="/images/automate/snow_integration_scripted_restapi_1.png" alt="Scripted REST Service]">}}

{{< figure src="/images/automate/snow_integration_scripted_restapi_2.png" alt="Scripted REST Resource">}}
#### Event
- The `x_chef_automate.chef.process.data` event is triggered whenever ServiceNow receives the node data from either terminal or through app.
{{< figure src="/images/automate/snow_integration_event_registration.png" alt="Event Registration">}}

- The `x_chef_automate.chef.process.installation` event is triggered whenever asset import table receives software installation details to map in the _CMDB server_ table.
{{< figure src="/images/automate/snow_integration_event_registration1.png" alt="Event Registration">}}

#### Script Action
The **Asset Process** script action processes the event and updates the node data into the _asset import_ table.
{{< figure src="/images/automate/snow_integration_script_action.png" alt="Script Action]">}}

The **Process Software Instance** script action processes the event and updates the software installation data in _CMDB_ sever table.
{{< figure src="/images/automate/snow_integration_script_action1.png" alt="Script Action]">}}

#### Transform Map
Transform map maps the source table, asset import to the target table, and CMDB Server.
#### Field Maps
Field maps establish a relationship between a field in an import set table and a field in the target table.
| Source field | Target field |
 :---        |    :--- |
| asset_tag | asset_tag |
| attributes | attributes |
| cpu_core_count | cpu_core_count |
| cpu_count | cpu_count |
| cpu_manufacturer | cpu_manufacturer |
| cpu_speed | cpu_speed |
| cpu_type | cpu_type |
| disk_space | disk_space |
| dns_domain | dns_domain |
| host_name | host_name |
| ip_address | ip_address |
| mac_address | mac_address |
| manufacturer | manufacturer |
| model | model_id |
| model_category_name | x_chef_automate_model_categoryname |
| name | name |
| os | os |
| os_service_pack | os_service_pack |
| os_version | os_version |
| ram | ram |
| serial_number | serial_number |
| short_description | short_description |
| system_class_name | sys_class_name |
| virtual | virtual |
| [Script] | install_status |

{{< figure src="/images/automate/snow_integration_field_map.png" alt="Table Transaction Map">}}
#### Transform Map Scripts
The **Transformation** events occur during the process of transforming an import set table onto a target table. You can navigate to this script by selecting **Importing Data** >**Transform Event Scripts**.
- **on Before:** Mapping manufacturer, CPU manufacturer, model category, and model by using this script.
- **on After:** 
  - First on after script used to update **node id**, **node name**, **organization**, **IP address** and insert the record into chef automate client run cookbooks, chef automate client run recipes, chef automate client run lists.
  - Second on after script used to update software, name and installed on in _CMDB software instance_ table.
  - The third one is used to updating data in _cmdb\_ci\_spkg_ table and _cmdb\_software\_instance_ table.
  - The fourth one is used to inserting or updating data in _cmdb\_ci\_file\_system_ table.
  - The fifth one is used to inserting or updating data in _cmdb\_ci\_network\_adapter_. 
{{< figure src="/images/automate/snow_integration_transform_map_scripts.png" alt="Transform Scripts">}}
##### *Script Includes:*
- **BufferToImportUtil:** Updates the CPU speed in Linux server, CMDB server, CMDI OSX server, and CMDB win server.
- **Logging:** Updates logging status as **enabled** or **disabled**.
- **JsonUtil:** Updates the JSON objects.
- **AutomateApi:** Calls the Chef automate API and checks the status of the response.
- **AutomateApiClient:** Updates the token and password in the Chef Automate instance.
- **PropertiesUtil:** Updates the role in the property.
- **ImportUtil:** Updates the compliance import data.
- **Util:** Updates manufacturer data in the core company table and updates the category and manufacturer data in the _CMDB model_ table.

{{< figure src="/images/automate/snow_integration_transform_map_scripts1.png" alt="Script Includes">}}
## Chef Automate Configuration
Chef Automate can be configured in two ways:
- From the ServiceNow application
- From the Chef Automate UI

**Precaution** - It is recommended to take a cautionary approach to set up the data imports using Chef Automate Data Feed to avoid unnecessary performance impacts to your ServiceNow infrastructure. In addition, it is recommended that larger time intervals for the **feed_interval** setting, such as 4 or 8 hours, are used initially and with a smaller **node_batch_size** of **50**. Testing this on your development infrastructure is recommended and making use of the CIDR filter functionality on Chef Automate Data Feed may make sense for your production environment, depending on the size of your estate and your Production ServiceNow setup.
### Configuring Chef Automate from ServiceNow
Configuration from a ServiceNow production instance requires the Chef Automate instance to be configured with a valid signed SSL certificate.
Follow these steps to configure the Chef Automate integration from ServiceNow:
1. Navigate to the **Chef Automate** menu.
2. Select the **Automate instances** module.
1. Click the **Submit** button to create a new Automate instance.
1. Enter the following details:
   - Name - a unique name for the integration.
   - URL - URL of the Chef Automate instance.
   - Chef automate API token - an API token generated for Chef Automate with data-feed-service authorization.
   - ServiceNow user - a ServiceNow user with **ITIL** and `x_chef_automate.api` roles.
   - ServiceNow password - password for the ServiceNow user being set.
1. Click the **Test Connectivity** button. This will check that the values are correct and that there is connectivity between Chef Automate and the application. A successful test shows the **service-now automate connectivity passed** message. If there is an error, a message is displayed that will help resolve any connectivity or credentials issues. 
1. Click **Submit**. The Automate instance details are saved.
### Configuring Chef Automate from Chef Automate UI
Follow these steps to configure the Chef Automate integration from the Chef Automate UI :
1. Navigate to the **Settings** menu.
1. Click the **Data Feeds** link from the **General** settings.
1. Click the **Create Data Feed** button to create a Data Feed.
1. Enter the following details:
   - Name - a unique name for the integration.
   - Data Feed URL – URL of the ServiceNow application Datafeed API ending `api/x_chef_automate/asset`.
   - Chef Automate API token - an API token generated for Chef Automate with data-feed-service authorization.
   - ServiceNow user - a ServiceNow user with **ITIL** and `x_chef_automate.api` roles.
   - ServiceNow password - Password for the ServiceNow user being set.<br  />
   Note that the Data Feed URL is the ServiceNow application, `FQDN/api/x_chef_automate/asset`. For example,  <https://venxxx.service-now.com/api/x_chef_automate/asset>.
1. Click the **Test Data Feed** button. This will check that the values are correct and that there is connectivity between Chef Automate and the ServiceNow application. A successful test shows the **service-now automate connectivity passed** message. If there is an error, a message is displayed that will help resolve any connectivity or credentials issues. 
1. Click **Create Data Feed**. The configuration is saved.
   
{{< figure src="/images/automate/snow_integration_create_data_feed.png" alt="Create Data Feed">}}
## Navigation
In ServiceNow, the navigation of the application is through the **Chef Automate** menu.
{{< figure src="/images/automate/snow_integration_navigation.png" alt="Navigation">}}

The **Automate Instances** module allows the user to configure the integration with Chef Automate. The main module for navigation is the **Servers** module. The **Server** module displays a list of servers in the _CMDB_ module. The Chef Automate integration augments the existing CMDB servers and inserts new servers into CMDB. The application uses the ServiceNow discovery IRE (Identification and Reconciliation Engine) when inserting or updating servers.<br  />
In addition, the application updates the CMDB file systems and software installed tables, and adds related information on the servers form with related data from Chef Automate:
- Client Runs
- Attributes
- Compliance Reports
### Client Runs
From a server record, the user can drill down into **Client Run** detail by clicking the name on an individual Client Run. Client run record displays related information for:
- Client run cookbooks - cookbooks executed during the client run.
- Client run lists - run lists executed during the client run.
- Client run recipes - recipes executed during the client run.<br  />
Full Chef client run details are available for each server. Client runs are also available from the **Client Runs** module.
### Attributes
From a server record, the user can drill down into the current server attributes detail by clicking on the attributes record. Full OHAI attributes are available for each server.
### Compliance Reports
From a server record, the user can drill down into **Compliance** report detail by clicking the name on an individual **Compliance** report. Compliance report record displays related information for:
- Compliance report profiles - all profiles executed during the compliance scan.
- Compliance report results - all results from the compliance scan.
In addition, the user can drill down into each **Compliance** report profile to view the individual results for each profile.<br  /> Full Chef compliance report details are available for each server. Compliance reports are also available from the **Compliance** report module.
## Uninstalling
To uninstall the application:
- In the ServiceNow instance, navigate to the **System Applications** > **Applications** menu.
- From the **Downloads** tab, click the **Chef Automate** link.
- In the **Related Links** section, select **Uninstall**. 