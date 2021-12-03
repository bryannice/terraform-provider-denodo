package denodo

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDatabaseRole() *schema.Resource {
	return &schema.Resource{
		CreateContext: createDatabaseRole,
		DeleteContext: deleteDatabaseRole,
		ReadContext:   readDatabaseRole,
		UpdateContext: updateDatabaseRole,
		Schema: map[string]*schema.Schema{
			"admin": &schema.Schema{
				Default:     false,
				Description: "Admin privilege over a database.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"all_privileges": &schema.Schema{
				Default:     false,
				Description: "All privileges CONNECT, CREATE, CREATE_DATA_SOURCE, CREATE_VIEW, CREATE_DATA_SERVICE, CREATE_FOLDER, EXECUTE, METADATA, WRITE, and FILE.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"all_users": &schema.Schema{
				Default:     false,
				Description: "granted by default to new local users. Additionally, you can automatically grant it to all users that connect to Virtual DataPort using Kerberos authentication, SAML 2.0 or to a database with LDAP authentication.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"assign_privileges": &schema.Schema{
				Default:     false,
				Description: "grants the privilege of granting/revoking privileges to other users.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"connect": &schema.Schema{
				Default:     false,
				Description: "Connect privilege to the database.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"create": &schema.Schema{
				Default:     false,
				Description: "Create privilege for all objects in a database.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"create_data_service": &schema.Schema{
				Default:     false,
				Description: "Create data service privilege in a database.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"create_data_source": &schema.Schema{
				Default:     false,
				Description: "Create data source privilege in a database.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"create_folder": &schema.Schema{
				Default:     false,
				Description: "Create folder privilege in a database.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"create_role": &schema.Schema{
				Default:     false,
				Description: "create_user and create_role: grant the privilege of creating users and roles.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"create_temporary_table": &schema.Schema{
				Default:     false,
				Description: "grants the privilege of creating temporary tables. This is useful to allow a user account to create temporary tables but do not want to grant the privilege CREATE nor CREATE VIEW because it would allow the user to create other types of elements.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"create_user": &schema.Schema{
				Default:     false,
				Description: "create_user and create_role: grant the privilege of creating users and roles.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"create_view": &schema.Schema{
				Default:     false,
				Description: "Create view privilege in a database.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"data_catalog_admin": &schema.Schema{
				Default:     false,
				Description: "Create/Delete categories, Edit categories, Assign categories, Create/Delete tags, Edit tags, Assign tags, Create/Delete custom elements, Edit custom elements, Assign custom elements, Edit elements, Synchronize, Import/Export, Servers, Personalize, Content, Permissions, Create endorsements, Edit endorsements, Delete endorsements, Create warnings, Edit warnings, Delete warnings, Create deprecations, Edit deprecations, and Delete deprecations",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"data_catalog_classifier": &schema.Schema{
				Default:     false,
				Description: "Personalize, and Content",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"data_catalog_content_admin": &schema.Schema{
				Default:     false,
				Description: "can change it Assign categories, Assign tags, and Assign custom elements",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"data_catalog_editor": &schema.Schema{
				Default:     false,
				Description: "Edit categories, Assign categories, Edit tags, Assign tags, Edit custom elements, Assign custom elements, Edit elements, Edit endorsements, and Edit warnings",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"data_catalog_exporter": &schema.Schema{
				Default:     false,
				Description: "In the Export dialog you can configure that these users are the only ones authorized to export the query results to specific formats.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"data_catalog_manager": &schema.Schema{
				Default:     false,
				Description: "Create/Delete categories, Edit categories, Assign categories, Create/Delete tags, Edit tags, Assign tags, Create/Delete custom elements, Edit custom elements, Assign custom elements, Edit elements, Create endorsements, Edit endorsements, Delete endorsements, Create warnings, Edit warnings, Delete warnings, Create deprecations, Edit deprecations, and Delete deprecations",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"database_name": &schema.Schema{
				Description: "Name of the database the role belongs too.",
				ForceNew:    true,
				Required:    true,
				Type:        schema.TypeString,
			},
			"delete": &schema.Schema{
				Default:     false,
				Description: "delete privileges on a database.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"diagnostic_monitoring_tool_admin": &schema.Schema{
				Default:     false,
				Description: "role to be able to change the server configuration of the Diagnostic & Monitoring Tool.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"diagnostic_monitoring_tool_create_diagnostic": &schema.Schema{
				Default:     false,
				Description: "role to create a new diagnostic.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"disable_cache_query": &schema.Schema{
				Default:     false,
				Description: "grants the privilege to execute queries disabling the cache of views over which the user does not have WRITE privleges. The cache can be disabled using the context clause 'cache'='off'.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"execute": &schema.Schema{
				Default:     false,
				Description: "Execute privilege on objects in a database.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"file": &schema.Schema{
				Default:     false,
				Description: "File privilege in a database to browse through directories.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"grant": &schema.Schema{
				Default:     true,
				Description: "Grant privileges on a database.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"impersonator": &schema.Schema{
				Default:     true,
				Description: "when users with this role publish REST web services, these services can impersonate other users.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"insert": &schema.Schema{
				Default:     false,
				Description: "insert privileges on a database.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"meta_data": &schema.Schema{
				Default:     false,
				Description: "Metadata privilege to get view information in the database.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"monitor_admin": &schema.Schema{
				Default:     false,
				Description: "grants the privilege of connecting to the monitoring interface of Virtual DataPort. This interface uses the JMX protocol (Java Management Extensions).",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"name": &schema.Schema{
				Description: "Name of the role being created.",
				ForceNew:    true,
				Required:    true,
				Type:        schema.TypeString,
			},
			"revoke": &schema.Schema{
				Default:     false,
				Description: "Revoke privileges on a database.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"scheduler_admin": &schema.Schema{
				Default:     false,
				Description: "Used by the Scheduler Administration Tool. The users that have this role assigned can perform any task in the Scheduler Administration Tool.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"server_admin": &schema.Schema{
				Default:     false,
				Description: "equivalent to being an administrator user of Virtual DataPort, except that it does not grant the privilege of connecting to Virtual DataPort via JMX. That is, a user with this role can manage databases, change settings of the Server, etc. A user with this role also needs the role 'assignprivileges' (see below) to manage the privileges of users and roles.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"update": &schema.Schema{
				Default:     false,
				Description: "update privileges on a database.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"write": &schema.Schema{
				Default:     false,
				Description: "Write privileges on elements in a database.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
		},
	}
}

func createDatabaseRole(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var admin bool
	var allPrivileges bool
	var allUsers bool
	var assignPrivileges bool
	var client *Client
	var connect bool
	var create bool
	var createDataService bool
	var createDataSource bool
	var createFolder bool
	var createRole bool
	var createTemporaryTable bool
	var createUser bool
	var createView bool
	var dataCatalogAdmin bool
	var dataCatalogClassifier bool
	var dataCatalogContentAdmin bool
	var dataCatalogEditor bool
	var dataCatalogExporter bool
	var dataCatalogManager bool
	var databaseName string
	var delete bool
	var diagnosticMonitoringToolAdmin bool
	var diagnosticMonitoringToolCreateDiagnostic bool
	var disableCacheQuery bool
	var diags diag.Diagnostics
	var err error
	var execute bool
	var file bool
	var impersonator bool
	var insert bool
	var grantClause []string
	var name string
	var metaData bool
	var monitorAdmin bool
	var roles []string
	var schedulerAdmin bool
	var sqlStmt string
	var serverAdmin bool
	var update bool
	var write bool

	admin = d.Get("admin").(bool)
	allPrivileges = d.Get("all_privileges").(bool)
	allUsers = d.Get("all_users").(bool)
	assignPrivileges = d.Get("assign_privileges").(bool)
	connect = d.Get("connect").(bool)
	create = d.Get("create").(bool)
	createDataService = d.Get("create_data_service").(bool)
	createDataSource = d.Get("create_data_source").(bool)
	createFolder = d.Get("create_folder").(bool)
	createRole = d.Get("create_role").(bool)
	createTemporaryTable = d.Get("create_temporary_table").(bool)
	createUser = d.Get("create_user").(bool)
	createView = d.Get("create_view").(bool)
	dataCatalogAdmin = d.Get("data_catalog_admin").(bool)
	dataCatalogClassifier = d.Get("data_catalog_classifier").(bool)
	dataCatalogContentAdmin = d.Get("data_catalog_content_admin").(bool)
	dataCatalogEditor = d.Get("data_catalog_editor").(bool)
	dataCatalogExporter = d.Get("data_catalog_exporter").(bool)
	dataCatalogManager = d.Get("data_catalog_manager").(bool)
	databaseName = d.Get("database_name").(string)
	delete = d.Get("delete").(bool)
	diagnosticMonitoringToolAdmin = d.Get("diagnostic_monitoring_tool_admin").(bool)
	diagnosticMonitoringToolCreateDiagnostic = d.Get("diagnostic_monitoring_tool_create_diagnostic").(bool)
	disableCacheQuery = d.Get("disable_cache_query").(bool)
	execute = d.Get("execute").(bool)
	file = d.Get("file").(bool)
	impersonator = d.Get("impersonator").(bool)
	insert = d.Get("insert").(bool)
	metaData = d.Get("meta_data").(bool)
	monitorAdmin = d.Get("monitor_admin").(bool)
	name = d.Get("name").(string)
	schedulerAdmin = d.Get("scheduler_admin").(bool)
	serverAdmin = d.Get("server_admin").(bool)
	update = d.Get("update").(bool)
	write = d.Get("write").(bool)

	sqlStmt = fmt.Sprintf(
		`
CREATE ROLE %s
GRANT `,
		name,
	)

	if admin {
		grantClause = append(grantClause, "ADMIN")
	}
	if allPrivileges {
		grantClause = append(grantClause, "ALL PRIVILEGES")
	}
	if connect {
		grantClause = append(grantClause, "CONNECT")
	}
	if create {
		grantClause = append(grantClause, "CREATE")
	}
	if createDataService {
		grantClause = append(grantClause, "CREATE_DATA_SERVICE")
	}
	if createDataSource {
		grantClause = append(grantClause, "CREATE_DATA_SOURCE")
	}
	if createFolder {
		grantClause = append(grantClause, "CREATE_FOLDER")
	}
	if createView {
		grantClause = append(grantClause, "CREATE_VIEW")
	}
	if delete {
		grantClause = append(grantClause, "DELETE")
	}
	if execute {
		grantClause = append(grantClause, "EXECUTE")
	}
	if file {
		grantClause = append(grantClause, "FILE")
	}
	if insert {
		grantClause = append(grantClause, "INSERT")
	}
	if metaData {
		grantClause = append(grantClause, "METADATA")
	}
	if update {
		grantClause = append(grantClause, "UPDATE")
	}
	if write {
		grantClause = append(grantClause, "WRITE")
	}

	sqlStmt += fmt.Sprintf(
		"%s ON %s;",
		strings.Join(grantClause, ", "),
		databaseName,
	)

	if assignPrivileges ||
		allUsers ||
		createRole ||
		createTemporaryTable ||
		createUser ||
		dataCatalogAdmin ||
		dataCatalogClassifier ||
		dataCatalogContentAdmin ||
		dataCatalogEditor ||
		dataCatalogExporter ||
		dataCatalogManager ||
		diagnosticMonitoringToolAdmin ||
		diagnosticMonitoringToolCreateDiagnostic ||
		disableCacheQuery ||
		impersonator ||
		monitorAdmin ||
		schedulerAdmin ||
		serverAdmin {
		sqlStmt += fmt.Sprintf(
			`
			ALTER ROLE %s`,
			name,
		)
		if assignPrivileges {
			grantClause = append(grantClause, "assignprivileges")
		}
		if allUsers {
			grantClause = append(grantClause, "allusers")
		}
		if createRole {
			grantClause = append(grantClause, "create_role")
		}
		if createTemporaryTable {
			grantClause = append(grantClause, "create_temporary_table")
		}
		if createUser {
			grantClause = append(grantClause, "create_user")
		}
		if dataCatalogAdmin {
			grantClause = append(grantClause, "data_catalog_admin")
		}
		if dataCatalogClassifier {
			grantClause = append(grantClause, "data_catalog_classifier")
		}
		if dataCatalogContentAdmin {
			grantClause = append(grantClause, "data_catalog_content_admin")
		}
		if dataCatalogEditor {
			grantClause = append(grantClause, "data_catalog_editor")
		}
		if dataCatalogExporter {
			grantClause = append(grantClause, "data_catalog_exporter")
		}
		if dataCatalogManager {
			grantClause = append(grantClause, "data_catalog_manager")
		}
		if diagnosticMonitoringToolAdmin {
			grantClause = append(grantClause, "diagnostic_monitoring_tool_admin")
		}
		if diagnosticMonitoringToolCreateDiagnostic {
			grantClause = append(grantClause, "diagnostic_monitoring_tool_create_diagnostic")
		}
		if disableCacheQuery {
			grantClause = append(grantClause, "disable_cache_query")
		}
		if impersonator {
			grantClause = append(grantClause, "impersonator")
		}
		if monitorAdmin {
			grantClause = append(roles, "monitor_admin")
		}
		if schedulerAdmin {
			grantClause = append(roles, "scheduler_admin")
		}
		if serverAdmin {
			grantClause = append(roles, "serveradmin")
		}
		sqlStmt += fmt.Sprintf(
			`
			GRANT ROLE %s;`,
			strings.Join(roles, ", "),
		)
	}

	client = meta.(*Client)

	err = client.ExecuteSQL(&sqlStmt)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(name)

	diags = readDatabaseRole(ctx, d, meta)

	return diags
}

func deleteDatabaseRole(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var client *Client
	var diags diag.Diagnostics
	var err error
	var name string
	var sqlStmt string

	name = d.Id()
	sqlStmt = fmt.Sprintf(
		"DROP ROLE IF EXISTS %s;",
		name,
	)
	client = meta.(*Client)

	err = client.ExecuteSQL(&sqlStmt)
	if err != nil {
		return diag.FromErr(err)
	} else {
		d.SetId("")
	}

	return diags
}

func readDatabaseRole(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var client *Client
	var databaseName string
	var diags diag.Diagnostics
	var err error
	var name string
	var resultSet [][]string
	var sqlStmt string

	databaseName = d.Get("database_name").(string)
	name = d.Id()
	sqlStmt = fmt.Sprintf(
		`
CONNECT DATABASE %s;
DESC ROLE %s;`,
		databaseName,
		name,
	)

	client = meta.(*Client)

	resultSet, err = client.ResultSet(&sqlStmt)
	if err != nil {
		return diag.FromErr(err)
	}

	if len(resultSet) != 0 {
		d.Set("name", name)
		d.Set("database_name", databaseName)
	}

	return diags
}

func updateDatabaseRole(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var admin bool
	var allPrivileges bool
	var allUsers bool
	var assignPrivileges bool
	var client *Client
	var connect bool
	var create bool
	var createDataService bool
	var createDataSource bool
	var createFolder bool
	var createRole bool
	var createTemporaryTable bool
	var createUser bool
	var createView bool
	var dataCatalogAdmin bool
	var dataCatalogClassifier bool
	var dataCatalogContentAdmin bool
	var dataCatalogEditor bool
	var dataCatalogExporter bool
	var dataCatalogManager bool
	var databaseName string
	var delete bool
	var diagnosticMonitoringToolAdmin bool
	var diagnosticMonitoringToolCreateDiagnostic bool
	var disableCacheQuery bool
	var diags diag.Diagnostics
	var err error
	var execute bool
	var file bool
	var grant bool
	var grantClause []string
	var impersonator bool
	var insert bool
	var name string
	var metaData bool
	var monitorAdmin bool
	var revoke bool
	var schedulerAdmin bool
	var sqlStmt string
	var serverAdmin bool
	var update bool
	var write bool

	admin = d.Get("admin").(bool)
	allPrivileges = d.Get("all_privileges").(bool)
	allUsers = d.Get("all_users").(bool)
	assignPrivileges = d.Get("assign_privileges").(bool)
	connect = d.Get("connect").(bool)
	create = d.Get("create").(bool)
	createDataService = d.Get("create_data_service").(bool)
	createDataSource = d.Get("create_data_source").(bool)
	createFolder = d.Get("create_folder").(bool)
	createRole = d.Get("create_role").(bool)
	createTemporaryTable = d.Get("create_temporary_table").(bool)
	createUser = d.Get("create_user").(bool)
	createView = d.Get("create_view").(bool)
	dataCatalogAdmin = d.Get("data_catalog_admin").(bool)
	dataCatalogClassifier = d.Get("data_catalog_classifier").(bool)
	dataCatalogContentAdmin = d.Get("data_catalog_content_admin").(bool)
	dataCatalogEditor = d.Get("data_catalog_editor").(bool)
	dataCatalogExporter = d.Get("data_catalog_exporter").(bool)
	dataCatalogManager = d.Get("data_catalog_manager").(bool)
	databaseName = d.Get("database_name").(string)
	delete = d.Get("delete").(bool)
	diagnosticMonitoringToolAdmin = d.Get("diagnostic_monitoring_tool_admin").(bool)
	diagnosticMonitoringToolCreateDiagnostic = d.Get("diagnostic_monitoring_tool_create_diagnostic").(bool)
	disableCacheQuery = d.Get("disable_cache_query").(bool)
	execute = d.Get("execute").(bool)
	file = d.Get("file").(bool)
	grant = d.Get("grant").(bool)
	impersonator = d.Get("impersonator").(bool)
	insert = d.Get("insert").(bool)
	metaData = d.Get("meta_data").(bool)
	monitorAdmin = d.Get("monitor_admin").(bool)
	name = d.Get("name").(string)
	revoke = d.Get("revoke").(bool)
	schedulerAdmin = d.Get("scheduler_admin").(bool)
	serverAdmin = d.Get("server_admin").(bool)
	update = d.Get("update").(bool)
	write = d.Get("write").(bool)

	sqlStmt = fmt.Sprintf(
		"ALTER ROLE %s\n",
		name,
	)

	if grant {
		sqlStmt += "GRANT "
	}
	if revoke {
		sqlStmt += "REVOKE "
	}

	if !assignPrivileges &&
		!allUsers &&
		!createRole &&
		!createTemporaryTable &&
		!createUser &&
		!dataCatalogAdmin &&
		!dataCatalogClassifier &&
		!dataCatalogContentAdmin &&
		!dataCatalogEditor &&
		!dataCatalogExporter &&
		!dataCatalogManager &&
		!diagnosticMonitoringToolAdmin &&
		!diagnosticMonitoringToolCreateDiagnostic &&
		!disableCacheQuery &&
		!impersonator &&
		!monitorAdmin &&
		!schedulerAdmin &&
		serverAdmin {

		if admin {
			grantClause = append(grantClause, "ADMIN")
		}
		if allPrivileges {
			grantClause = append(grantClause, "ALL PRIVILEGES")
		}
		if connect {
			grantClause = append(grantClause, "CONNECT")
		}
		if create {
			grantClause = append(grantClause, "CREATE")
		}
		if createDataService {
			grantClause = append(grantClause, "CREATE_DATA_SERVICE")
		}
		if createDataSource {
			grantClause = append(grantClause, "CREATE_DATA_SOURCE")
		}
		if createFolder {
			grantClause = append(grantClause, "CREATE_FOLDER")
		}
		if createView {
			grantClause = append(grantClause, "CREATE_VIEW")
		}
		if delete {
			grantClause = append(grantClause, "DELETE")
		}
		if execute {
			grantClause = append(grantClause, "EXECUTE")
		}
		if file {
			grantClause = append(grantClause, "FILE")
		}
		if insert {
			grantClause = append(grantClause, "INSERT")
		}
		if metaData {
			grantClause = append(grantClause, "METADATA")
		}
		if update {
			grantClause = append(grantClause, "UPDATE")
		}
		if write {
			grantClause = append(grantClause, "WRITE")
		}
		sqlStmt += fmt.Sprintf(
			"%s ON %s",
			strings.Join(grantClause, ", "),
			databaseName,
		)
	} else {
		if assignPrivileges {
			grantClause = append(grantClause, "assignprivileges")
		}
		if allUsers {
			grantClause = append(grantClause, "allusers")
		}
		if createRole {
			grantClause = append(grantClause, "create_role")
		}
		if createTemporaryTable {
			grantClause = append(grantClause, "create_temporary_table")
		}
		if createUser {
			grantClause = append(grantClause, "create_user")
		}
		if dataCatalogAdmin {
			grantClause = append(grantClause, "data_catalog_admin")
		}
		if dataCatalogClassifier {
			grantClause = append(grantClause, "data_catalog_classifier")
		}
		if dataCatalogContentAdmin {
			grantClause = append(grantClause, "data_catalog_content_admin")
		}
		if dataCatalogEditor {
			grantClause = append(grantClause, "data_catalog_editor")
		}
		if dataCatalogExporter {
			grantClause = append(grantClause, "data_catalog_exporter")
		}
		if dataCatalogManager {
			grantClause = append(grantClause, "data_catalog_manager")
		}
		if diagnosticMonitoringToolAdmin {
			grantClause = append(grantClause, "diagnostic_monitoring_tool_admin")
		}
		if diagnosticMonitoringToolCreateDiagnostic {
			grantClause = append(grantClause, "diagnostic_monitoring_tool_create_diagnostic")
		}
		if disableCacheQuery {
			grantClause = append(grantClause, "disable_cache_query")
		}
		if impersonator {
			grantClause = append(grantClause, "impersonator")
		}
		if monitorAdmin {
			grantClause = append(grantClause, "monitor_admin")
		}
		if schedulerAdmin {
			grantClause = append(grantClause, "scheduler_admin")
		}
		if serverAdmin {
			grantClause = append(grantClause, "serveradmin")
		}
		sqlStmt += fmt.Sprintf(
			"ROLE %s",
			strings.Join(grantClause, ", "),
		)
	}

	sqlStmt += ";"
	client = meta.(*Client)

	err = client.ExecuteSQL(&sqlStmt)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(d.Get("name").(string))

	diags = readDatabaseRole(ctx, d, meta)

	return diags
}
