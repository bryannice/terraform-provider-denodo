package denodo

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDerivedView() *schema.Resource {
	return &schema.Resource{
		CreateContext: createDerivedView,
		DeleteContext: deleteDerivedView,
		ReadContext:   readDerivedView,
		UpdateContext: updateDerivedView,
		Schema: map[string]*schema.Schema{
			"create_date": &schema.Schema{
				Computed:    true,
				Description: "Date when the element was created.",
				Type:        schema.TypeString,
			},
			"database": &schema.Schema{
				Description: "Database where the base view will reside.",
				Required:    true,
				Type:        schema.TypeString,
			},
			"description": &schema.Schema{
				Computed:    true,
				Description: "Description of the element.",
				Type:        schema.TypeString,
			},
			"folder": &schema.Schema{
				Computed:    true,
				Description: "Folder of the element in lowercase. If the element is not in any folder, the value is /.",
				Type:        schema.TypeString,
			},
			"last_modification_date": &schema.Schema{
				Computed:    true,
				Description: "Date when the element was modified for the last time. If the element was never modified, the value is the same as create_date",
				Type:        schema.TypeString,
			},
			"last_user_modifier": &schema.Schema{
				Computed:    true,
				Description: "User that modified the element for the last time. If the element was never modified, the value is the same as user_creator.",
				Type:        schema.TypeString,
			},
			"name": &schema.Schema{
				Description: "Name of the view.",
				Required:    true,
				Type:        schema.TypeString,
			},
			"user_creator": &schema.Schema{
				Computed:    true,
				Description: "Owner of the element.",
				Type:        schema.TypeString,
			},
			"vql": &schema.Schema{
				Description: "VQL selection statement used to create or replace a dervived view.",
				Required:    true,
				Type:        schema.TypeString,
			},
		},
	}
}

func createDerivedView(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var client *Client
	var diags diag.Diagnostics
	var err error
	var vql string

	vql = d.Get("vql").(string)

	client = meta.(*Client)
	err = client.ExecuteSQL(&vql)
	if err != nil {
		return diag.FromErr(err)
	}

	diags = readDerivedView(ctx, d, meta)

	return diags
}

func deleteDerivedView(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var client *Client
	var database string
	var diags diag.Diagnostics
	var err error
	var name string
	var vql string

	database = d.Get("database").(string)
	name = d.Get("name").(string)

	vql = fmt.Sprintf(
		`CONNECT DATABASE %s;
DROP VIEW IF EXISTS %s CASCADE;`,
		database,
		name,
	)
	client = meta.(*Client)
	err = client.ExecuteSQL(&vql)
	if err != nil {
		return diag.FromErr(err)
	}

	if err != nil {
		return diag.FromErr(err)
	} else {
		d.SetId("")
	}

	return diags
}

func readDerivedView(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var client *Client
	var database string
	var diags diag.Diagnostics
	var err error
	var name string
	var resultSet [][]string
	var sqlStmt string

	database = d.Get("database").(string)
	name = d.Get("name").(string)

	sqlStmt = fmt.Sprintf(
		`
CONNECT DATABASE %s;
SELECT
  internal_id,
  database_name,
  name,
  user_creator,
  last_user_modifier,
  create_date,
  last_modification_date,
  description,
  folder
FROM GET_ELEMENTS()
WHERE name = '%s'
  AND database_name = '%s'
  AND type = 'view'
  AND subtype = 'derived';`,
		name,
		database,
	)

	client = meta.(*Client)

	resultSet, err = client.ResultSet(&sqlStmt)
	if err != nil {
		return diag.FromErr(err)
	}

	if err = d.Set("id", resultSet[0][0]); err != nil {
		diags = diag.FromErr(err)
	}
	if err = d.Set("database", resultSet[0][1]); err != nil {
		diags = diag.FromErr(err)
	}
	if err = d.Set("name", resultSet[0][2]); err != nil {
		diags = diag.FromErr(err)
	}
	if err = d.Set("user_creator", resultSet[0][3]); err != nil {
		diags = diag.FromErr(err)
	}
	if err = d.Set("last_user_modifier", resultSet[0][4]); err != nil {
		diags = diag.FromErr(err)
	}
	if err = d.Set("create_date", resultSet[0][5]); err != nil {
		diags = diag.FromErr(err)
	}
	if err = d.Set("last_modification_date", resultSet[0][6]); err != nil {
		diags = diag.FromErr(err)
	}
	if err = d.Set("description", resultSet[0][7]); err != nil {
		diags = diag.FromErr(err)
	}
	if err = d.Set("folder", resultSet[0][8]); err != nil {
		diags = diag.FromErr(err)
	}

	return diags
}

func updateDerivedView(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	diags = createDerivedView(ctx, d, meta)
	return diags
}
