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
			"allPrivilege": &schema.Schema{
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
			"createDataService": &schema.Schema{
				Default:     false,
				Description: "Create data service privilege in a database.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"createDataSource": &schema.Schema{
				Default:     false,
				Description: "Create data source privilege in a database.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"createFolder": &schema.Schema{
				Default:     false,
				Description: "Create folder privilege in a database.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"createView": &schema.Schema{
				Default:     false,
				Description: "Create view privilege in a database.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"databaseName": &schema.Schema{
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
			"metaData": &schema.Schema{
				Default:     false,
				Description: "Metadata privilege to get view information in the database.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"monitorAdmin": &schema.Schema{
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
			"schedulerAdmin": &schema.Schema{
				Description: "Scheduling admin role on the database.",
				ForceNew:    true,
				Required:    true,
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

	databaseName = d.Get("databaseName").(string)
	name = d.Get("name").(string)

	sqlStmt = fmt.Sprintf(
		"CREATE ROLE %s\nGRANT ",
		name,
	)

	if d.Get("admin").(bool) {
		grantClause = append(grantClause, "ADMIN")
	}
	if d.Get("allPrivilege").(bool) {
		grantClause = append(grantClause, "ALL PRIVILEGES")
	}
	if d.Get("connect").(bool) {
		grantClause = append(grantClause, "CONNECT")
	}
	if d.Get("create").(bool) {
		grantClause = append(grantClause, "CREATE")
	}
	if d.Get("createDataService").(bool) {
		grantClause = append(grantClause, "CREATE_DATA_SERVICE")
	}
	if d.Get("createDataSource").(bool) {
		grantClause = append(grantClause, "CREATE_DATA_SOURCE")
	}
	if d.Get("createFolder").(bool) {
		grantClause = append(grantClause, "CREATE_FOLDER")
	}
	if d.Get("createView").(bool) {
		grantClause = append(grantClause, "CREATE_VIEW")
	}
	if d.Get("execute").(bool) {
		grantClause = append(grantClause, "EXECUTE")
	}
	if d.Get("file").(bool) {
		grantClause = append(grantClause, "FILE")
	}
	if d.Get("metaData").(bool) {
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

	d.SetId(d.Get("name").(string))

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
	var diags diag.Diagnostics
	var err error
	var name string
	var resultSet [][]string
	var sqlStmt string

	name = d.Id()
	sqlStmt = fmt.Sprintf(
		`
SELECT DISTINCT
	dbname AS db_name,
	rolename AS role_name,
	CASE
		WHEN dbadmin like '%true%'
		THEN 1
		ELSE 0
	END AS db_admin,
	CASE
		WHEN dbconnect like '%true%'
		THEN 1
		ELSE 0
	END AS db_connect,
	CASE
		WHEN dbcreate like '%true%'
		THEN 1
		ELSE 0
	END AS db_create,
	CASE
		WHEN dbcreatedataservice like '%true%'
		THEN 1
		ELSE 0
	END AS db_create_data_service,
	CASE
		WHEN dbcreatedatasource like '%true%'
		THEN 1
		ELSE 0
	END AS db_create_data_source,
	CASE
		WHEN dbcreatefolder like '%true%'
		THEN 1
		ELSE 0
	END AS db_create_folder,
	CASE
		WHEN dbcreateview like '%true%'
		THEN 1
		ELSE 0
	END AS db_create_view,
	CASE
		WHEN dbexecute like '%true%'
		THEN 1
		ELSE 0
	END AS db_execute,
	CASE
		WHEN dbfile like '%true%'
		THEN 1
		ELSE 0
	END AS db_file,
	CASE
		WHEN dbmetadata like '%true%'
		THEN 1
		ELSE 0
	END AS db_meta_data,
	CASE
		WHEN dbwrite like '%true%'
		THEN 1
		ELSE 0
	END AS db_write
FROM
	CATALOG_PERMISSIONS()
WHERE
	rolename = '%s'`,
		name,
	)

	client = meta.(*Client)

	resultSet, err = client.ResultSet(&sqlStmt)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("databaseName", resultSet[0][0])
	d.Set("name", resultSet[0][1])
	d.Set("admin", resultSet[0][2])
	d.Set("connect", resultSet[0][3])
	d.Set("create", resultSet[0][4])
	d.Set("createDataService", resultSet[0][5])
	d.Set("createDataSource", resultSet[0][6])
	d.Set("createFolder", resultSet[0][7])
	d.Set("createView", resultSet[0][8])
	d.Set("execute", resultSet[0][9])
	d.Set("file", resultSet[0][10])
	d.Set("metaData", resultSet[0][11])
	d.Set("write", resultSet[0][12])

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

	databaseName = d.Get("databaseName").(string)
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

	if !d.Get("monitorAdmin").(bool) && !d.Get("schedulerAdmin").(bool) {
		if d.Get("admin").(bool) {
			grantClause = append(grantClause, "ADMIN")
		}
		if d.Get("allPrivilege").(bool) {
			grantClause = append(grantClause, "ALL PRIVILEGES")
		}
		if d.Get("connect").(bool) {
			grantClause = append(grantClause, "CONNECT")
		}
		if d.Get("create").(bool) {
			grantClause = append(grantClause, "CREATE")
		}
		if d.Get("createDataService").(bool) {
			grantClause = append(grantClause, "CREATE_DATA_SERVICE")
		}
		if d.Get("createDataSource").(bool) {
			grantClause = append(grantClause, "CREATE_DATA_SOURCE")
		}
		if d.Get("createFolder").(bool) {
			grantClause = append(grantClause, "CREATE_FOLDER")
		}
		if d.Get("createView").(bool) {
			grantClause = append(grantClause, "CREATE_VIEW")
		}
		if d.Get("execute").(bool) {
			grantClause = append(grantClause, "EXECUTE")
		}
		if d.Get("file").(bool) {
			grantClause = append(grantClause, "FILE")
		}
		if d.Get("metaData").(bool) {
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
		if d.Get("monitorAdmin").(bool) {
			grantClause = append(grantClause, "monitor_admin")
		}
		if d.Get("schedulerAdmin").(bool) {
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
