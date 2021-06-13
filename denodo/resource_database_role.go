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
			"create_view": &schema.Schema{
				Default:     false,
				Description: "Create view privilege in a database.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"database_name": &schema.Schema{
				Description: "Name of the database the role belongs too.",
				ForceNew:    true,
				Required:    true,
				Type:        schema.TypeString,
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
			"meta_data": &schema.Schema{
				Default:     false,
				Description: "Metadata privilege to get view information in the database.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"monitor_admin": &schema.Schema{
				Default:     false,
				Description: "Monitoring admin role on the database.",
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
				Description: "Scheduling admin role on the database.",
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
	var client *Client
	var connect bool
	var create bool
	var createDataService bool
	var createDataSource bool
	var createFolder bool
	var createView bool
	var databaseName string
	var diags diag.Diagnostics
	var err error
	var execute bool
	var file bool
	var grantClause []string
	var name string
	var metaData bool
	var monitorAdmin bool
	var roles []string
	var schedulerAdmin bool
	var sqlStmt string
	var write bool

	admin = d.Get("admin").(bool)
	allPrivileges = d.Get("all_privileges").(bool)
	connect = d.Get("connect").(bool)
	create = d.Get("create").(bool)
	createDataService = d.Get("create_data_service").(bool)
	createDataSource = d.Get("create_data_source").(bool)
	createFolder = d.Get("create_folder").(bool)
	createView = d.Get("create_view").(bool)
	databaseName = d.Get("database_name").(string)
	execute = d.Get("execute").(bool)
	file = d.Get("file").(bool)
	metaData = d.Get("meta_data").(bool)
	monitorAdmin = d.Get("monitor_admin").(bool)
	name = d.Get("name").(string)
	schedulerAdmin = d.Get("scheduler_admin").(bool)
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
	if execute {
		grantClause = append(grantClause, "EXECUTE")
	}
	if file {
		grantClause = append(grantClause, "FILE")
	}
	if metaData {
		grantClause = append(grantClause, "METADATA")
	}
	if write {
		grantClause = append(grantClause, "WRITE")
	}

	sqlStmt += fmt.Sprintf(
		"%s ON %s;",
		strings.Join(grantClause, ", "),
		databaseName,
	)

	if monitorAdmin || schedulerAdmin {
		sqlStmt += fmt.Sprintf(
			`
			ALTER ROLE %s`,
			name,
		)
		if monitorAdmin {
			roles = append(roles, "monitor_admin")
		}
		if schedulerAdmin {
			roles = append(roles, "scheduler_admin")
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
	var client *Client
	var connect bool
	var create bool
	var createDataService bool
	var createDataSource bool
	var createFolder bool
	var createView bool
	var databaseName string
	var diags diag.Diagnostics
	var err error
	var execute bool
	var file bool
	var grant bool
	var grantClause []string
	var name string
	var metaData bool
	var monitorAdmin bool
	var revoke bool
	var schedulerAdmin bool
	var sqlStmt string
	var write bool

	admin = d.Get("admin").(bool)
	allPrivileges = d.Get("all_privileges").(bool)
	connect = d.Get("connect").(bool)
	create = d.Get("create").(bool)
	createDataService = d.Get("create_data_service").(bool)
	createDataSource = d.Get("create_data_source").(bool)
	createFolder = d.Get("create_folder").(bool)
	createView = d.Get("create_view").(bool)
	databaseName = d.Get("database_name").(string)
	execute = d.Get("execute").(bool)
	file = d.Get("file").(bool)
	grant = d.Get("grant").(bool)
	metaData = d.Get("meta_data").(bool)
	monitorAdmin = d.Get("monitor_admin").(bool)
	name = d.Get("name").(string)
	revoke = d.Get("revoke").(bool)
	schedulerAdmin = d.Get("scheduler_admin").(bool)
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

	if !monitorAdmin && !schedulerAdmin {
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
		if execute {
			grantClause = append(grantClause, "EXECUTE")
		}
		if file {
			grantClause = append(grantClause, "FILE")
		}
		if metaData {
			grantClause = append(grantClause, "METADATA")
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
		if monitorAdmin {
			grantClause = append(grantClause, "monitor_admin")
		}
		if schedulerAdmin {
			grantClause = append(grantClause, "scheduler_admin")
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
