package denodo

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceDatabaseRole() *schema.Resource {
	return &schema.Resource{
		Create: CreateDatabaseRole,
		Delete: DeleteDatabaseRole,
		Exists: nil,
		Read: ReadDatabaseRole,
		Update: UpdateDatabaseRole,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"admin": {
				Default: false,
				Description: "Admin privilege over a database.",
				Optional: true,
				Type: schema.TypeBool,
			},
			"allPrivilege": {
				Default: false,
				Description: "All privileges CONNECT, CREATE, CREATE_DATA_SOURCE, CREATE_VIEW, CREATE_DATA_SERVICE, CREATE_FOLDER, EXECUTE, METADATA, WRITE, and FILE.",
				Optional: true,
				Type: schema.TypeBool,
			},
			"connect": {
				Default: false,
				Description: "Connect privilege to the database.",
				Optional: true,
				Type: schema.TypeBool,
			},
			"create": {
				Default: false,
				Description: "Create privilege for all objects in a database.",
				Optional: true,
				Type: schema.TypeBool,
			},
			"createDataService": {
				Default: false,
				Description: "Create data service privilege in a database.",
				Optional: true,
				Type: schema.TypeBool,
			},
			"createDataSource": {
				Default: false,
				Description: "Create data source privilege in a database.",
				Optional: true,
				Type: schema.TypeBool,
			},
			"createFolder": {
				Default: false,
				Description: "Create folder privilege in a database.",
				Optional: true,
				Type: schema.TypeBool,
			},
			"createView": {
				Default: false,
				Description: "Create view privilege in a database.",
				Optional: true,
				Type: schema.TypeBool,
			},
			"databaseName": {
				Description: "Name of the database the role belongs too.",
				ForceNew: true,
				Required: true,
				Type: schema.TypeString,
			},
			"execute": {
				Default: false,
				Description: "Execute privilege on objects in a database.",
				Optional: true,
				Type: schema.TypeBool,
			},
			"file": {
				Default: false,
				Description: "File privilege in a database to browse through directories.",
				Optional: true,
				Type: schema.TypeBool,
			},
			"grant": {
				Default: true,
				Description: "Grant privileges on a database.",
				Optional: true,
				Type: schema.TypeBool,
			},
			"metaData": {
				Default: false,
				Description: "Metadata privilege to get view information in the database.",
				Optional: true,
				Type: schema.TypeBool,
			},
			"monitorAdmin": {
				Default: false,
				Description: "Monitoring admin role on the database.",
				Optional: true,
				Type: schema.TypeBool,
			}
			"name": {
				Description: "Name of the role being created.",
				ForceNew: true,
				Required: true,
				Type: schema.TypeString,
			},
			"revoke": {
				Default: false,
				Description: "Revoke privileges on a database.",
				Optional: true,
				Type: schema.TypeBool,
			},
			"schedulerAdmin": {
				Description: "Scheduling admin role on the database.",
				ForceNew: true,
				Required: true,
				Type: schema.TypeBool,
			},
			"write": {
				Default: false,
				Description: "Write privileges on elements in a database.",
				Optional: true,
				Type: schema.TypeBool,
			},
		},
	}
}

func CreateDatabaseRole(d *schema.ResourceData, meta interface{}) error {
	var grantClause []string

	databaseName := d.Get("databaseName").(string)
	name := d.Get("name").(string)

	sqlStmt := fmt.Sprintf(
		"CREATE ROLE %s\nGRANT ",
		name,
	)

	if d.Get("admin") {
		grantClause = append(grantClause, "ADMIN")
	}
	if d.Get("allPrivilege") {
		grantClause = append(grantClause, "ALL PRIVILEGES")
	}
	if d.Get("connect") {
		grantClause = append(grantClause, "CONNECT")
	}
	if d.Get("create") {
		grantClause = append(grantClause, "CREATE")
	}
	if d.Get("createDataService") {
		grantClause = append(grantClause, "CREATE_DATA_SERVICE")
	}
	if d.Get("createDataSource") {
		grantClause = append(grantClause, "CREATE_DATA_SOURCE")
	}
	if d.Get("createFolder") {
		grantClause = append(grantClause, "CREATE_FOLDER")
	}
	if d.Get("createView") {
		grantClause = append(grantClause, "CREATE_VIEW")
	}
	if d.Get("execute") {
		grantClause = append(grantClause, "EXECUTE")
	}
	if d.Get("file") {
		grantClause = append(grantClause, "FILE")
	}
	if d.Get("metaData") {
		grantClause = append(grantClause, "METADATA")
	}
	if d.Get("write") {
		grantClause = append(grantClause, "WRITE")
	}

	sqlStmt += fmt.Sprintf(
		"%s ON %s;",
		strings.Join(grantClause, ", "),
		databaseName,
	)
	log.Println("Executing statement: ", sqlStmt)

	db, err := connectToDenodo(meta.(*DenodoConfiguration))
	if err != nul {
		return err
	}

	_, err = db.Exec(sqlStmt)
	if err != nul {
		return err
	}

	d.SetId(d.Get("name").(string))

	return ReadDatabaseRole(d, meta)
}

func DeleteDatabaseRole(d *schema.ResourceData, meta interface{}) error {
	name := d.Id()

	sqlStmt := fmt.Sprintf(
		"DROP ROLE %s;",
		name,
	)
	log.Println("Executing statement: ", sqlStmt)

	db, err := connectToDenodo(meta.(*DenodoConfiguration))
	if err != nil {
		return err
	}

	_, err = db.Exec(stmtSQL)
	if err == nil {
		d.SetId("")
	}
	return err
}

func ReadDatabaseRole(d *schema.ResourceData, meta interface{}) error {
	name := d.Id()

	sqlStmt := fmt.Sprintf(
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
	log.Println("Executing statement: ", sqlStmt)

	var dbName string
	var roleName string
	var dbAdmin bool
	var dbConnect bool
	var dbCreate bool
	var dbCreateDataService bool
	var dbCreateDataSource bool
	var dbCreateFolder bool
	var dbCreateView bool
	var dbExecute bool
	var dbFile bool
	var dbMetaData bool
	var dbWrite bool

	err = db.QueryRow(sqlStmt).Scan(
		&dbName,
		&roleName,
		&dbAdmin,
		&dbConnect,
		&dbCreate,
		&dbCreateDataService,
		&dbCreateDataSource,
		&dbCreateFolder,
		&dbCreateView,
		&dbExecute,
		&dbFile,
		&dbMetaData,
		&dbWrite,
	)
	if err != nil {
		return err
	}

	d.Set("databaseName", dbName)
	d.Set("name", roleName)
	d.Set("admin", dbAdmin)
	d.Set("connect", dbConnect)
	d.Set("create", dbCreate)
	d.Set("createDataService", dbCreateDataService)
	d.Set("createDataSource", dbCreateDataSource)
	d.Set("createFolder", dbCreateFolder)
	d.Set("createView", dbCreateView)
	d.Set("execute", dbExecute)
	d.Set("file", dbFile)
	d.Set("metaData", dbMetaData)
	d.Set("write", dbWrite)

	return nil
}

func UpdateDatabaseRole(d *schema.ResourceData, meta interface{}) error {
	var grantClause []string

	databaseName := d.Get("databaseName").(string)
	name := d.Get("name").(string)

	sqlStmt := fmt.Sprintf(
		"ALTER ROLE %s\n",
		name,
	)

	if d.Get("grant") {
		sqlStmt += "GRANT "
	}
	if d.Get("revoke") {
		sqlStmt += "REVOKE "
	}

	if ! d.Get("monitorAdmin") && ! d.Get("schedulerAdmin") {
		if d.Get("admin") {
			grantClause = append(grantClause, "ADMIN")
		}
		if d.Get("allPrivilege") {
			grantClause = append(grantClause, "ALL PRIVILEGES")
		}
		if d.Get("connect") {
			grantClause = append(grantClause, "CONNECT")
		}
		if d.Get("create") {
			grantClause = append(grantClause, "CREATE")
		}
		if d.Get("createDataService") {
			grantClause = append(grantClause, "CREATE_DATA_SERVICE")
		}
		if d.Get("createDataSource") {
			grantClause = append(grantClause, "CREATE_DATA_SOURCE")
		}
		if d.Get("createFolder") {
			grantClause = append(grantClause, "CREATE_FOLDER")
		}
		if d.Get("createView") {
			grantClause = append(grantClause, "CREATE_VIEW")
		}
		if d.Get("execute") {
			grantClause = append(grantClause, "EXECUTE")
		}
		if d.Get("file") {
			grantClause = append(grantClause, "FILE")
		}
		if d.Get("metaData") {
			grantClause = append(grantClause, "METADATA")
		}
		if d.Get("write") {
			grantClause = append(grantClause, "WRITE")
		}
		sqlStmt += fmt.Sprintf(
			"%s ON %s",
			strings.Join(grantClause, ", "),
			databaseName,
		)
	} else {
		if d.Get("monitorAdmin") {
			grantClause = append(grantClause, "monitor_admin")
		}
		if d.Get("schedulerAdmin") {
			grantClause = append(grantClause, "scheduler_admin")
		}
		sqlStmt += fmt.Sprintf(
			"ROLE %s",
			strings.Join(grantClause, ", "),
		)
	}

	sqlStmt += ";"
	log.Println("Executing statement: ", sqlStmt)

	db, err := connectToDenodo(meta.(*DenodoConfiguration))
	if err != nul {
		return err
	}

	_, err = db.Exec(sqlStmt)
	if err != nul {
		return err
	}

	d.SetId(d.Get("name").(string))

	return ReadDatabaseRole(d, meta)
}
