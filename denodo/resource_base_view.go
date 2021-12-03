package denodo

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceBaseView() *schema.Resource {
	return &schema.Resource{
		CreateContext: createBaseView,
		DeleteContext: deleteBaseView,
		ReadContext:   readBaseView,
		UpdateContext: updateBaseView,
		Schema: map[string]*schema.Schema{
			"database": &schema.Schema{
				Description: "Database where the base view will reside.",
				Required:    true,
				Type:        schema.TypeString,
			},
			"data_source_name": &schema.Schema{
				Description: "Name of the data source.",
				Required:    true,
				Type:        schema.TypeString,
			},
			"data_source_catalog_name": &schema.Schema{
				Default:     "NULL",
				Description: "Name of the catalog in the source database that contains the table from which to create the base view. Pass null if the database does not support catalogs.",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"data_source_database": &schema.Schema{
				Description: "Name of the database to which the JDBC data source belongs.",
				Required:    true,
				Type:        schema.TypeString,
			},
			"data_source_schema_name": &schema.Schema{
				Default:     "NULL",
				Description: "Name of the schema in the source database that contains the table from which to create the base view. Pass null if the database does not support schemas.",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"data_source_table_name": &schema.Schema{
				Description: "Name of the table/view over which to create the base view.",
				Required:    true,
				Type:        schema.TypeString,
			},
			"description": &schema.Schema{
				Computed:    true,
				Description: "Description of the element.",
				Type:        schema.TypeString,
			},
			"folder": &schema.Schema{
				Default:     "NULL",
				Description: "Folder in which to place the created base view. The result will include the VQL statements to create this folder(s). If null, the VQL will not specify a folder.",
				Optional:    true,
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
				Description: "Name of the base view to be created. If null, the name will be auto-generated.",
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

func createBaseView(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var client *Client
	var database string
	var dataSourceName string
	var dataSourceCatalogName string
	var dataSourceDatabase string
	var dataSourceSchemaName string
	var dataSourceTableName string
	var diags diag.Diagnostics
	var err error
	var folder string
	var name string
	var resultSet [][]string
	var sqlStmt string
	var vql string

	database = d.Get("database").(string)
	dataSourceName = d.Get("data_source_name").(string)
	dataSourceCatalogName = d.Get("data_source_catalog_name").(string)
	dataSourceDatabase = d.Get("data_source_database").(string)
	dataSourceSchemaName = d.Get("data_source_schema_name").(string)
	dataSourceTableName = d.Get("data_source_table_name").(string)
	folder = d.Get("folder").(string)
	name = d.Get("name").(string)
	vql = d.Get("vql").(string)

	sqlStmt = fmt.Sprintf(
		`CONNECT DATABASE %s;
`,
		database,
	)

	if vql == "" {
		sqlStmt += fmt.Sprintf(
			`CALL GENERATE_VQL_TO_CREATE_JDBC_BASE_VIEW(
	'%s',
	%s,
	%s,
	'%s',
	'%s',
	%s,
	NULL,
	'%s'
);`,
			dataSourceName,
			TenaryString(dataSourceCatalogName == "NULL", dataSourceCatalogName, fmt.Sprintf("'%s'", dataSourceCatalogName)),
			TenaryString(dataSourceSchemaName == "NULL", dataSourceSchemaName, fmt.Sprintf("'%s'", dataSourceSchemaName)),
			dataSourceTableName,
			name,
			TenaryString(folder == "NULL", folder, fmt.Sprintf("'%s'", folder)),
			dataSourceDatabase,
		)

		client = meta.(*Client)

		resultSet, err = client.ResultSet(&sqlStmt)
		if err != nil {
			return diag.FromErr(err)
		}

		for _, generatedVql := range resultSet {
			createStmt := fmt.Sprintf(
				`CONNECT DATABASE %s;
%s`,
				database,
				generatedVql[0],
			)
			err = client.ExecuteSQL(&createStmt)
			if err != nil {
				return diag.FromErr(err)
			}
		}
	} else {
		sqlStmt += vql
		err = client.ExecuteSQL(&sqlStmt)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	diags = readBaseView(ctx, d, meta)

	return diags
}

func deleteBaseView(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var client *Client
	var database string
	var diags diag.Diagnostics
	var err error
	var name string
	var sqlStmt string

	database = d.Get("database").(string)
	name = d.Id()

	sqlStmt = fmt.Sprintf(
		`
CONNECT DATABASE %s;
DROP VIEW IF EXISTS %s CASCADE;`,
		database,
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

func readBaseView(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
  AND subtype = 'base';`,
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

func updateBaseView(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	// ToDo: Add logic for Alter table to modify Base View
	return diags
}
