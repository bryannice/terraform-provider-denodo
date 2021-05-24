package denodo

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceUser() *schema.Resource {
	return &schema.Resource{
		Create: CreateUser,
		Delete: DeleteUser,
		Exists: ExistsUser,
		Read:   ReadUser,
		Update: UpdateUser,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"description" : {
				Default: "user",
				Description: "The description of the user.",
				Optimal: true,
				Type: schema.TypeString,
			},
			"password": {
				Default:     nil,
				Description: "Password associated to the user.",
				ForceNew: true,
				Required: true,
				Type:        schema.TypeString,
			},
			"roles": {
				Default:     nil,
				Description: "Password associated to the user.",
				ForceNew: true,
				Required: true,
				Type:        schema.TypeString,
			},
			"username": {
				Default:     nil,
				Description: "Username to be created.",
				ForceNew: true,
				Required: true,
				Type:        schema.TypeString,
			},
		}
	}
}

func CreateUser(d *schema.ResourceData, meta interface{}) error {
	description := d.Get("description").(string)
	password := d.Get("password").(string)
	roles := d.Get("roles").(string)
	username := d.Get("username").(string)

	var encryptPassword string
	err := db.QueryRow(
		fmt.Sprintf(
			"ENCRYPT_PASSWORD '%s';",
			password,
		),
	).Scan(&encryptPassword)
	if err != nil {
		fmt.Println(err)
	}

	sqlStmt := fmt.Sprintf(
		`
CREATE USER %s %s ENCRYPTED TRANSFER
%s`,
		username,
		encryptPassword,
		description,
	)

	if roles != nil {
		sqlStmt += fmt.Sprintf(
			"\nGRANT ROLE %s",
			roles,
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

	d.SetId(d.Get("username").(string))

	return ReadDatabase(d, meta)
}

func DeleteUser(d *schema.ResourceData, meta interface{}) error {
	username := d.Get("username").(string)

	sqlStmt := fmt.Sprintf(
		"DROP USER %s;",
		username,
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

func ExistsUser(d *schema.ResourceData, meta interface{}) error {

}

func ReadUser(d *schema.ResourceData, meta interface{}) error {

}

func UpdateUser(d *schema.ResourceData, meta interface{}) error {

}