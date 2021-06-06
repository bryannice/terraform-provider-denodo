package denodo

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceUser() *schema.Resource {
	return &schema.Resource{
		CreateContext: createUser,
		DeleteContext: deleteUser,
		ReadContext:   readUser,
		UpdateContext: updateUser,
		Schema: map[string]*schema.Schema{
			"admin": &schema.Schema{
				Default:     false,
				Description: "User type admin.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"description": &schema.Schema{
				Default:     "user",
				Description: "The description of the user.",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"grant": &schema.Schema{
				Default:     true,
				Description: "Grant privileges on a role.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"password": &schema.Schema{
				Default:     nil,
				Description: "Password associated to the user.",
				ForceNew:    true,
				Required:    true,
				Type:        schema.TypeString,
			},
			"roles": &schema.Schema{
				Default:     nil,
				Description: "Password associated to the user.",
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
			"username": &schema.Schema{
				Default:     nil,
				Description: "Username to be created.",
				ForceNew:    true,
				Required:    true,
				Type:        schema.TypeString,
			},
		},
	}
}

func createUser(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var admin bool
	var client *Client
	var description string
	var diags diag.Diagnostics
	var encryptPassword string
	var encryptPasswordCommand string
	var err error
	var password string
	var resultSet [][]string
	var roles string
	var sqlStmt string
	var username string

	admin = d.Get("admin").(bool)
	description = d.Get("description").(string)
	password = d.Get("password").(string)
	roles = d.Get("roles").(string)
	username = d.Get("username").(string)

	client = meta.(*Client)

	encryptPasswordCommand = fmt.Sprintf(
		"ENCRYPT_PASSWORD '%s';",
		&password,
	)

	resultSet, err = client.ResultSet(&encryptPasswordCommand)
	if err != nil {
		return diag.FromErr(err)
	}

	encryptPassword = resultSet[0][0]

	sqlStmt = "CREATE USER "

	if admin {
		sqlStmt += "ADMIN "
	}

	sqlStmt += fmt.Sprintf(
		"%s %s ENCRYPTED TRANSFER\n%s",
		username,
		encryptPassword,
		description,
	)

	if roles != "" {
		sqlStmt += fmt.Sprintf(
			"\nGRANT ROLE %s",
			roles,
		)
	}
	sqlStmt += ";"
	client = meta.(*Client)

	err = client.ExecuteSQL(&sqlStmt)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(d.Get("username").(string))

	diags = readUser(ctx, d, meta)

	return diags
}

func deleteUser(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var client *Client
	var diags diag.Diagnostics
	var err error
	var sqlStmt string
	var username string

	username = d.Get("username").(string)
	sqlStmt = fmt.Sprintf(
		"DROP USER %s;",
		username,
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

func readUser(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var client *Client
	var diags diag.Diagnostics
	var err error
	var name string
	var resultSet [][]string
	var sqlStmt string

	name = d.Id()
	sqlStmt = fmt.Sprintf(
		"DESC USER %s;",
		name,
	)

	client = meta.(*Client)

	resultSet, err = client.ResultSet(&sqlStmt)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("admin", resultSet[0][0])
	d.Set("description", resultSet[0][1])
	d.Set("username", resultSet[0][2])

	return diags
}

func updateUser(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var client *Client
	var diags diag.Diagnostics
	var err error
	var grant bool
	var revoke bool
	var roles string
	var sqlStmt string
	var username string

	grant = d.Get("grant").(bool)
	username = d.Get("username").(string)
	revoke = d.Get("revoke").(bool)
	roles = d.Get("roles").(string)
	sqlStmt = fmt.Sprintf(
		"ALTER USER %s\n",
		username,
	)

	if grant {
		sqlStmt += "GRANT "
	}
	if revoke {
		sqlStmt += "REVOKE "
	}

	sqlStmt += fmt.Sprintf(
		"ROLE %s;",
		roles,
	)

	client = meta.(*Client)

	err = client.ExecuteSQL(&sqlStmt)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(d.Get("username").(string))

	diags = readUser(ctx, d, meta)

	return diags
}
