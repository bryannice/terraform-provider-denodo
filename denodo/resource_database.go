package denodo

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDatabase() *schema.Resource {
	return &schema.Resource{
		CreateContext: createDatabase,
		DeleteContext: deleteDatabase,
		ReadContext:   readDatabase,
		UpdateContext: updateDatabase,
		Schema: map[string]*schema.Schema{
			"authentication": &schema.Schema{
				Default:     "LOCAL",
				Description: "Authenication method the database will use.",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"char_set": &schema.Schema{
				Default:     "DEFAULT",
				Description: "Setting the charset. Valid values are UNICODE, RESTRICTED, or DEFAULT.",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"check_view_restrictions": &schema.Schema{
				Default:     "DEFAULT",
				Description: "The section Column Privileges, Row Restrictions and Custom Policies Are Always Propagated of the Upgrade Guide explains the implications of changing this parameter.",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"cost_optimization": &schema.Schema{
				Default:     "DEFAULT",
				Description: "Enables or disables the cost-based optimization on this database.",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"description": &schema.Schema{
				Default:     "database",
				Description: "The description of the database.",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"name": &schema.Schema{
				Description: "Name of the database to be created.",
				ForceNew:    true,
				Required:    true,
				Type:        schema.TypeString,
			},
			"odbc_authentication": &schema.Schema{
				Default:     "NORMAL",
				Description: "ODBC Authenication method it will use. Valid values are NORMAL or KERBEROS.",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"query_simplification": &schema.Schema{
				Default:     "DEFAULT",
				Description: "Enables or disables automatic simplification of queries on the database. Valid values are ON, OFF, or DEFAULT.",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"summary_rewrite": &schema.Schema{
				Default:     "DEFAULT",
				Description: "Enables or disables the summary rewrite optimization on the database. Valid values are ON, OFF, or DEFAULT.",
				Optional:    true,
				Type:        schema.TypeString,
			},
		},
	}
}

func createDatabase(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var authentication string
	var charSet string
	var checkViewRestrictions string
	var client *Client
	var costOptimization string
	var description string
	var diags diag.Diagnostics
	var err error
	var name string
	var odbcAuthentication string
	var querySimplification string
	var sqlStmt string
	var summaryRewrite string

	authentication = d.Get("authentication").(string)
	charSet = d.Get("char_set").(string)
	checkViewRestrictions = d.Get("check_view_restrictions").(string)
	costOptimization = d.Get("cost_optimization").(string)
	description = d.Get("description").(string)
	name = d.Get("name").(string)
	odbcAuthentication = d.Get("odbc_authentication").(string)
	querySimplification = d.Get("query_simplification").(string)
	summaryRewrite = d.Get("summary_rewrite").(string)

	client = meta.(*Client)

	sqlStmt = fmt.Sprintf(
		`
CREATE DATABASE %s
'%s'
CHARSET %s
AUTHENTICATION %s
ODBC AUTHENTICATION %s
CHECK_VIEW_RESTRICTIONS %s;
ALTER DATABASE %s
COST OPTIMIZATION %s
QUERY SIMPLIFICATION %s
SUMMARY REWRITE %s;`,
		name,
		description,
		charSet,
		authentication,
		odbcAuthentication,
		checkViewRestrictions,
		name,
		costOptimization,
		querySimplification,
		summaryRewrite,
	)

	err = client.ExecuteSQL(&sqlStmt)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(d.Get("name").(string))

	diags = readDatabase(ctx, d, meta)

	return diags
}

func deleteDatabase(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var client *Client
	var diags diag.Diagnostics
	var err error
	var name string
	var sqlStmt string

	name = d.Id()

	client = meta.(*Client)

	sqlStmt = fmt.Sprintf(
		"DROP DATABASE IF EXISTS %s CASCADE;",
		name,
	)

	err = client.ExecuteSQL(&sqlStmt)
	if err != nil {
		return diag.FromErr(err)
	} else {
		d.SetId("")
	}

	return diags
}

func readDatabase(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var client *Client
	var diags diag.Diagnostics
	var err error
	var name string
	var resultSet [][]string
	var sqlStmt string

	name = d.Id()

	client = meta.(*Client)

	sqlStmt = fmt.Sprintf(
		"DESC DATABASE %s;",
		name,
	)

	resultSet, err = client.ResultSet(&sqlStmt)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("name", resultSet[0][0])
	d.Set("description", resultSet[0][1])

	return diags
}

func updateDatabase(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var authentication string
	var charSet string
	var client *Client
	var costOptimization string
	var diags diag.Diagnostics
	var err error
	var name string
	var odbcAuthentication string
	var querySimplification string
	var sqlStmt string
	var summaryRewrite string

	authentication = d.Get("authentication").(string)
	charSet = d.Get("char_set").(string)
	costOptimization = d.Get("cost_optimization").(string)
	name = d.Get("name").(string)
	odbcAuthentication = d.Get("odbc_authentication").(string)
	querySimplification = d.Get("query_simplification").(string)
	summaryRewrite = d.Get("summary_rewrite").(string)

	client = meta.(*Client)

	sqlStmt = fmt.Sprintf(
		"ALTER DATABASE %s",
		name,
	)
	if authentication != "" {
		sqlStmt += fmt.Sprintf(
			"\nAUTHENTICATION %s",
			authentication,
		)
	}
	if charSet != "" {
		sqlStmt += fmt.Sprintf(
			"\nCHARSET %s",
			charSet,
		)
	}
	if costOptimization != "" {
		sqlStmt += fmt.Sprintf(
			"\nCOST OPTIMIZATION %s",
			costOptimization,
		)
	}
	if odbcAuthentication != "" {
		sqlStmt += fmt.Sprintf(
			"\nODBC AUTHENTICATION %s",
			odbcAuthentication,
		)
	}
	if summaryRewrite != "" {
		sqlStmt += fmt.Sprintf(
			"\nSUMMARY REWRITE %s",
			summaryRewrite,
		)
	}
	if querySimplification != "" {
		sqlStmt += fmt.Sprintf(
			"\nQUERY SIMPLIFICATION %s",
			querySimplification,
		)
	}

	err = client.ExecuteSQL(&sqlStmt)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(d.Get("name").(string))

	diags = readDatabase(ctx, d, meta)

	return diags
}
