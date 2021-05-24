package denodo

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceDatabase() *schema.Resource {
	return &schema.Resource{
		Create: CreateDatabase,
		Delete: DeleteDatabase,
		Exists: nil,
		Read: ReadDatabase,
		Update: UpdateDatabase,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"authentication" : {
				Default: "LOCAL",
				Description: "Authenication method the database will use.",
				Optimal: true,
				Type: schema.TypeString,
			},
			"char_set" : {
				Default: "DEFAULT",
				Description: "Setting the charset. Valid values are UNICODE, RESTRICTED, or DEFAULT.",
				Optimal: true,
				Type: schema.TypeString,
			},
			"cost_optimization" : {
				Default: "DEFAULT",
				Description: "Enables or disables the cost-based optimization on this database.",
				Optimal: true,
				Type: schema.TypeString,
			},
			"description" : {
				Default: "database",
				Description: "The description of the database.",
				Optimal: true,
				Type: schema.TypeString,
			},
			"name": {
				Description: "Name of the database to be created.",
				ForceNew: true,
				Required: true,
				Type: schema.TypeString,
			},
			"odbc_authentication": {
				Default: "NORMAL",
				Description: "ODBC Authenication method it will use. Valid values are NORMAL or KERBEROS.",
				Optimal: true,
				Type: schema.TypeString,
			},
			"summary_rewrite": {
				Default: "DEFAULT",
				Description: "Enables or disables the summary rewrite optimization on the database. Valid values are ON, OFF, or DEFAULT.",
				Optimal: true,
				Type: schema.TypeString,
			},
			"query_simplification": {
				Default: "DEFAULT",
				Description: "Enables or disables automatic simplification of queries on the database. Valid values are ON, OFF, or DEFAULT.",
				Optimal: true,
				Type: schema.TypeString,
			},
		},
	}
}

func CreateDatabase(d *schema.ResourceData, meta interface{}) error {
	authentication := d.Get("authentication").(string)
	charSet := d.Get("char_set").(string)
	checkViewRestrictions := d.Get("check_view_restrictions").(string)
	description := d.Get("description").(string)
	name := d.Get("name").(string)
	odbcAuthentication := d.Get("odbc_authentication").(string)

	sqlStmt := fmt.Sprintf(
		"CREATE DATABASE %s\n%s\nCHARSET %s\nAUTHENTICATION %s\nODBC AUTHENTICATION %s\nCHECK_VIEW_RESTRICTIONS %s;",
		name,
		description,
		charSet,
		authentication,
		odbcAuthentication,
		checkViewRestrictions,
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

	return ReadDatabase(d, meta)
}

func DeleteDatabase(d *schema.ResourceData, meta interface{}) error {
	name := d.Id()

	sqlStmt := fmt.Sprintf(
		"DROP DATABASE %s;",
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

func ReadDatabase(d *schema.ResourceData, meta interface{}) error {
	name := d.Id()

	sqlStmt := fmt.Sprintf(
		"SELECT DISTINCT database_name\nFROM get_elements()\nWHERE database_name = '%s';",
		name,
	)
	log.Println("Executing statement: ", sqlStmt)

	var dbName string
	err = db.QueryRow(sqlStmt).Scan(&dbName)
	if err != nil {
		return err
	}

	d.Set("name", dbName)

	return nil
}

func UpdateDatabase(d *schema.ResourceData, meta interface{}) error {
	authentication := d.Get("authentication").(string)
	charSet := d.Get("char_set").(string)
	costOptimization := d.Get("cost_optimization").(string)
	name := d.Get("name").(string)
	odbcAuthentication := d.Get("odbc_authentication").(string)
	summaryRewrite := d.Get("summary_rewrite").(string)
	querySimplification := d.Get("query_simplification").(string)

	sqlStmt := fmt.Sprintf(
		"ALTER DATABASE %s",
		name,
	)
	if authentication != nil {
		sqlStmt += fmt.Sprintf(
			"\nAUTHENTICATION %s",
			authentication,
		)
	}
	if charSet != nil {
		sqlStmt += fmt.Sprintf(
			"\nCHARSET %s",
			charSet,
		)
	}
	if costOptimization != nil {
		sqlStmt += fmt.Sprintf(
			"\nCOST OPTIMIZATION %s",
			costOptimization,
		)
	}
	if odbcAuthentication != nil {
		sqlStmt += fmt.Sprintf(
			"\nODBC AUTHENTICATION %s",
			odbcAuthentication,
		)
	}
	if summaryRewrite != nil {
		sqlStmt += fmt.Sprintf(
			"\nSUMMARY REWRITE %s",
			summaryRewrite,
		)
	}
	if querySimplification != nil {
		sqlStmt += fmt.Sprintf(
			"\nQUERY SIMPLIFICATION %s",
			querySimplification,
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

	return ReadDatabase(d, meta)
}