package denodo

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceObject() *schema.Resource {
	return &schema.Resource{
		ReadContext: readObject,
		Schema: map[string]*schema.Schema{
			"database": &schema.Schema{
				Description: "Database name.",
				Required:    true,
				Type:        schema.TypeString,
			},
			"object_name": &schema.Schema{
				Description: "Object name.",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"object_type": &schema.Schema{
				Default:     "Views",
				Description: "Type of object to retrieve. Folders, DataSources, StoredProcedures, Wrappers, Views, WebServices, Widgets, Associations, JMSListeners",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"objects": &schema.Schema{
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"catalog_id": &schema.Schema{
							Computed:    true,
							Description: "Catalog id of the object.",
							Type:        schema.TypeString,
						},
						"create_date": &schema.Schema{
							Computed:    true,
							Description: "Date when the element was created.",
							Type:        schema.TypeString,
						},
						"database_name": &schema.Schema{
							Computed:    true,
							Description: "Name of database where the element belongs to.",
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
						"object_name": &schema.Schema{
							Computed:    true,
							Description: "Name of the element.",
							Type:        schema.TypeString,
						},
						"sub_type": &schema.Schema{
							Computed:    true,
							Description: "subtype: subtype of the element or an empty string if the element does not have a subtype. Elements that have a subtype and what subtypes they can have. View: base, derived, interface or materialized. Datasource: custom, df, essbase, jdbc, json, ldap, odbc, olap, salesforce, sapbwbapi, saperp, ws or xml. Wrapper: custom, df, essbase, html, jdbc, json, ldap, odbc, olap, salesforce, sapbwbapi, saperp, ws or xml.",
							Type:        schema.TypeString,
						},
						"type": &schema.Schema{
							Computed:    true,
							Description: "type: type of the element. The values can be association, datasource, folder, storedProcedure, type, view, webService, widget, wrapper.",
							Type:        schema.TypeString,
						},
						"user_creator": &schema.Schema{
							Computed:    true,
							Description: "Owner of the element.",
							Type:        schema.TypeString,
						},
					},
				},
				Type: schema.TypeList,
			},
		},
	}
}

func readObject(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var client *Client
	var database string
	var diags diag.Diagnostics
	var resultSet [][]string
	var objectName string
	var objectType string
	var err error
	var records []interface{}
	var sqlStmt string

	database = d.Get("database").(string)
	objectName = d.Get("object_name").(string)
	objectType = d.Get("object_type").(string)

	sqlStmt = fmt.Sprintf(
		`
CALL GET_ELEMENTS(
    '%s',
    %s,
    '%s'
);`,
		database,
		TenaryString(objectName == "", "NULL", fmt.Sprintf("'%s'", objectName)),
		objectType,
	)

	client = meta.(*Client)

	resultSet, err = client.ResultSet(&sqlStmt)
	if err != nil {
		return diag.FromErr(err)
	}

	for _, tuple := range resultSet {

		records = append(
			records,
			map[string]interface{}{
				"database_name":          tuple[0],
				"object_name":            tuple[1],
				"type":                   tuple[2],
				"sub_type":               tuple[3],
				"user_creator":           tuple[4],
				"last_user_modifier":     tuple[5],
				"create_date":            tuple[6],
				"last_modification_date": tuple[7],
				"description":            tuple[8],
				"folder":                 tuple[9],
				"catalog_id":             tuple[11],
			},
		)
	}

	if err = d.Set("objects", records); err != nil {
		diags = diag.FromErr(err)
	}

	d.SetId(database)

	return diags
}
