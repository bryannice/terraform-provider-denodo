package denodo

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceJDBCDataSourceObject() *schema.Resource {
	return &schema.Resource{
		ReadContext: readJDBCDataSourceObject,
		Schema: map[string]*schema.Schema{
			"catalog_name": &schema.Schema{
				Default:     "NULL",
				Description: "Name of the catalog for which you want to get the list of tables. If the data source does not support catalogs, set to null. If null and the data source does support catalogs, the procedure will return all the matching tables across all catalogs.",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"database": &schema.Schema{
				Description: "Name of the database to which the data source belongs. If null, the procedure will use the current database.",
				Required:    true,
				Type:        schema.TypeString,
			},
			"name": &schema.Schema{
				Description: "Name of the data source for which you want to get the list of tables.",
				Required:    true,
				Type:        schema.TypeString,
			},
			"schema_name": &schema.Schema{
				Default:     "NULL",
				Description: "When the data source has to insert several rows into the database of this data source, it can insert them in batches. This number sets the number of queries per batch.",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"objects": &schema.Schema{
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"catalog_name": &schema.Schema{
							Computed:    true,
							Description: "Catalog name returned from data source.",
							Type:        schema.TypeString,
						},
						"schema_name": &schema.Schema{
							Computed:    true,
							Description: "Schema name returned from data source.",
							Type:        schema.TypeString,
						},
						"object_name": &schema.Schema{
							Computed:    true,
							Description: "Object name returned from data source.",
							Type:        schema.TypeString,
						},
						"object_type": &schema.Schema{
							Computed:    true,
							Description: "Object type returned from data source.",
							Type:        schema.TypeString,
						},
					},
				},
				Type: schema.TypeList,
			},
		},
	}
}

func readJDBCDataSourceObject(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var catalogName string
	var client *Client
	var database string
	var diags diag.Diagnostics
	var resultSet [][]string
	var err error
	var name string
	var records []interface{}
	var schemaName string
	var sqlStmt string

	catalogName = d.Get("catalog_name").(string)
	database = d.Get("database").(string)
	name = d.Get("name").(string)
	schemaName = d.Get("schema_name").(string)

	client = meta.(*Client)

	sqlStmt = fmt.Sprintf(
		`
CONNECT DATABASE %s;
CALL GET_JDBC_DATASOURCE_TABLES(
    '%s',
    %s,
    %s
);`,
		database,
		name,
		TenaryString(catalogName == "NULL", catalogName, fmt.Sprintf("'%s'", catalogName)),
		TenaryString(schemaName == "NULL", schemaName, fmt.Sprintf("'%s'", schemaName)),
	)

	resultSet, err = client.ResultSet(&sqlStmt)
	if err != nil {
		return diag.FromErr(err)
	}

	for _, tuple := range resultSet {

		records = append(
			records,
			map[string]interface{}{
				"catalog_name": tuple[0],
				"schema_name":  tuple[1],
				"object_name":  tuple[2],
				"object_type":  tuple[3],
			},
		)
	}

	if err = d.Set("objects", records); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(name)

	return diags
}
