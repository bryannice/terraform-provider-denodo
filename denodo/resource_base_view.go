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
				Required:    true,
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
				Required:    true,
				Type:        schema.TypeString,
			},
			"data_source_table_name": &schema.Schema{
				Description: "Name of the table/view over which to create the base view.",
				Required:    true,
				Type:        schema.TypeString,
			},
			"folder": &schema.Schema{
				Default:     "NULL",
				Description: "Folder in which to place the created base view. The result will include the VQL statements to create this folder(s). If null, the VQL will not specify a folder.",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"name": &schema.Schema{
				Description: "Name of the base view to be created. If null, the name will be auto-generated.",
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

	database = d.Get("database").(string)
	dataSourceName = d.Get("data_source_name").(string)
	dataSourceCatalogName = d.Get("data_source_catalog_name").(string)
	dataSourceDatabase = d.Get("data_source_database").(string)
	dataSourceSchemaName = d.Get("data_source_schema_name").(string)
	dataSourceTableName = d.Get("data_source_table_name").(string)
	folder = d.Get("folder").(string)
	name = d.Get("name").(string)

	sqlStmt = fmt.Sprintf(
		`
CONNECT DATABASE %s;
CALL GENERATE_VQL_TO_CREATE_JDBC_BASE_VIEW(
	'%s',
	%s,
	%s,
	'%s',
	'%s',
	%s,
	NULL
);
`,
		dataSourceDatabase,
		dataSourceName,
		TenaryString(dataSourceCatalogName == "NULL", dataSourceCatalogName, fmt.Sprintf("'%s'", dataSourceCatalogName)),
		TenaryString(dataSourceSchemaName == "NULL", dataSourceSchemaName, fmt.Sprintf("'%s'", dataSourceSchemaName)),
		dataSourceTableName,
		name,
		TenaryString(folder == "NULL", folder, fmt.Sprintf("'%s'", folder)),
	)

	client = meta.(*Client)

	resultSet, err = client.ResultSet(&sqlStmt)
	if err != nil {
		return diag.FromErr(err)
	}

	for _, generatedVql := range resultSet {
		createStmt := fmt.Sprintf(
			`
CONNECT DATABASE %s;
%s`,
			database,
			generatedVql[0],
		)
		err = client.ExecuteSQL(&createStmt)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	d.SetId(d.Get("name").(string))

	diags = readBaseView(ctx, d, meta)

	return diags
}

func deleteBaseView(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var client *Client
	var dataSourceDatabase string
	var diags diag.Diagnostics
	var err error
	var name string
	var sqlStmt string

	dataSourceDatabase = d.Get("data_source_database").(string)
	name = d.Id()

	sqlStmt = fmt.Sprintf(
		`
CONNECT DATABASE %s;
DROP VIEW IF EXISTS %s CASCADE;`,
		dataSourceDatabase,
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
	var dataSourceDatabase string
	var diags diag.Diagnostics
	var err error
	var name string
	var resultSet [][]string
	var sqlStmt string

	dataSourceDatabase = d.Get("data_source_database").(string)
	name = d.Id()

	sqlStmt = fmt.Sprintf(
		`
CONNECT DATABASE %s;
DESC VIEW %s;`,
		dataSourceDatabase,
		name,
	)

	client = meta.(*Client)

	resultSet, err = client.ResultSet(&sqlStmt)
	if err != nil {
		return diag.FromErr(err)
	}

	if len(resultSet) != 0 {
		d.Set("name", name)
		d.Set("database", dataSourceDatabase)
	}

	return diags
}

func updateBaseView(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	// ToDo: Add logic for Alter table to modify Base View
	return diags
}
