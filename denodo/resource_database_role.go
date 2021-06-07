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
	var client *Client
	var databaseName string
	var diags diag.Diagnostics
	var err error
	var grantClause []string
	var name string
	var sqlStmt string

	databaseName = d.Get("database_name").(string)
	name = d.Get("name").(string)

	sqlStmt = fmt.Sprintf(
		`
CREATE ROLE %s
GRANT `,
		name,
	)

	if d.Get("admin").(bool) {
		grantClause = append(grantClause, "ADMIN")
	}
	if d.Get("all_privileges").(bool) {
		grantClause = append(grantClause, "ALL PRIVILEGES")
	}
	if d.Get("connect").(bool) {
		grantClause = append(grantClause, "CONNECT")
	}
	if d.Get("create").(bool) {
		grantClause = append(grantClause, "CREATE")
	}
	if d.Get("create_data_service").(bool) {
		grantClause = append(grantClause, "CREATE_DATA_SERVICE")
	}
	if d.Get("create_data_source").(bool) {
		grantClause = append(grantClause, "CREATE_DATA_SOURCE")
	}
	if d.Get("create_folder").(bool) {
		grantClause = append(grantClause, "CREATE_FOLDER")
	}
	if d.Get("create_view").(bool) {
		grantClause = append(grantClause, "CREATE_VIEW")
	}
	if d.Get("execute").(bool) {
		grantClause = append(grantClause, "EXECUTE")
	}
	if d.Get("file").(bool) {
		grantClause = append(grantClause, "FILE")
	}
	if d.Get("meta_data").(bool) {
		grantClause = append(grantClause, "METADATA")
	}
	if d.Get("write").(bool) {
		grantClause = append(grantClause, "WRITE")
	}

	sqlStmt += fmt.Sprintf(
		"%s ON %s;",
		strings.Join(grantClause, ", "),
		databaseName,
	)

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
		"DROP ROLE %s;",
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
	var client *Client
	var databaseName string
	var diags diag.Diagnostics
	var err error
	var grantClause []string
	var name string
	var sqlStmt string

	databaseName = d.Get("database_name").(string)
	name = d.Get("name").(string)
	sqlStmt = fmt.Sprintf(
		"ALTER ROLE %s\n",
		name,
	)

	if d.Get("grant").(bool) {
		sqlStmt += "GRANT "
	}
	if d.Get("revoke").(bool) {
		sqlStmt += "REVOKE "
	}

	if !d.Get("monitor_admin").(bool) && !d.Get("scheduler_admin").(bool) {
		if d.Get("admin").(bool) {
			grantClause = append(grantClause, "ADMIN")
		}
		if d.Get("all_privileges").(bool) {
			grantClause = append(grantClause, "ALL PRIVILEGES")
		}
		if d.Get("connect").(bool) {
			grantClause = append(grantClause, "CONNECT")
		}
		if d.Get("create").(bool) {
			grantClause = append(grantClause, "CREATE")
		}
		if d.Get("create_data_service").(bool) {
			grantClause = append(grantClause, "CREATE_DATA_SERVICE")
		}
		if d.Get("create_data_source").(bool) {
			grantClause = append(grantClause, "CREATE_DATA_SOURCE")
		}
		if d.Get("create_folder").(bool) {
			grantClause = append(grantClause, "CREATE_FOLDER")
		}
		if d.Get("create_view").(bool) {
			grantClause = append(grantClause, "CREATE_VIEW")
		}
		if d.Get("execute").(bool) {
			grantClause = append(grantClause, "EXECUTE")
		}
		if d.Get("file").(bool) {
			grantClause = append(grantClause, "FILE")
		}
		if d.Get("meta_data").(bool) {
			grantClause = append(grantClause, "METADATA")
		}
		if d.Get("write").(bool) {
			grantClause = append(grantClause, "WRITE")
		}
		sqlStmt += fmt.Sprintf(
			"%s ON %s",
			strings.Join(grantClause, ", "),
			databaseName,
		)
	} else {
		if d.Get("monitor_admin").(bool) {
			grantClause = append(grantClause, "monitor_admin")
		}
		if d.Get("scheduler_admin").(bool) {
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
