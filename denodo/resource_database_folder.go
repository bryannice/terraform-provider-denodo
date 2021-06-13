package denodo

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDatabaseFolder() *schema.Resource {
	return &schema.Resource{
		CreateContext: createDatabaseFolder,
		DeleteContext: deleteDatabaseFolder,
		ReadContext:   readDatabaseFolder,
		UpdateContext: updateDatabaseFolder,
		Schema: map[string]*schema.Schema{
			"copy": &schema.Schema{
				Default:     false,
				Description: "Indicator to copy an element.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"copy_new": &schema.Schema{
				Description: "Copy to new element name.",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"copy_old": &schema.Schema{
				Description: "Copy from old elment name.",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"database": &schema.Schema{
				Description: "Database where the folder resides.",
				Required:    true,
				Type:        schema.TypeString,
			},
			"data_source_type": &schema.Schema{
				Description: "Type of data source element.",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"description": &schema.Schema{
				Default:     "Folder",
				Description: "Description of the folder.",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"element_type": &schema.Schema{
				Description: "Type of element.",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"folder_path": &schema.Schema{
				Description: "Folder path to create.",
				Required:    true,
				Type:        schema.TypeString,
			},
			"move": &schema.Schema{
				Default:     false,
				Description: "Indicator to move an element to the folder.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"rename_path": &schema.Schema{
				Description: "When the data source has to insert several rows into the database of this data source, it can insert them in batches. This number sets the number of queries per batch.",
				Optional:    true,
				Type:        schema.TypeString,
			},
		},
	}
}

func createDatabaseFolder(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var client *Client
	var database string
	var description string
	var diags diag.Diagnostics
	var err error
	var folderPath string
	var sqlStmt string

	database = d.Get("database").(string)
	description = d.Get("description").(string)
	folderPath = d.Get("folder_path").(string)

	sqlStmt = fmt.Sprintf(
		`
CONNECT DATABASE %s;
CREATE FOLDER '%s'
DESCRIPTION '%s';`,
		database,
		folderPath,
		description,
	)
	client = meta.(*Client)

	err = client.ExecuteSQL(&sqlStmt)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(d.Get("folder_path").(string))

	diags = readDatabaseFolder(ctx, d, meta)

	return diags
}

func deleteDatabaseFolder(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var client *Client
	var database string
	var diags diag.Diagnostics
	var err error
	var folderPath string
	var sqlStmt string

	database = d.Get("database").(string)
	folderPath = d.Id()
	sqlStmt = fmt.Sprintf(
		`
CONNECT DATABASE %s;
DROP FOLDER IF EXISTS '%s' CASCADE;`,
		database,
		folderPath,
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

func readDatabaseFolder(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var client *Client
	var database string
	var diags diag.Diagnostics
	var err error
	var folderPath string
	var resultSet [][]string
	var sqlStmt string

	database = d.Get("database").(string)
	folderPath = d.Id()
	sqlStmt = fmt.Sprintf(
		`
CONNECT DATABASE %s;
DESC FOLDER '%s';`,
		database,
		folderPath,
	)

	client = meta.(*Client)

	resultSet, err = client.ResultSet(&sqlStmt)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("name", resultSet[0][0])
	d.Set("folder_path", resultSet[0][1])
	d.Set("description", resultSet[0][2])

	return diags
}

func updateDatabaseFolder(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var client *Client
	var copy bool
	var copyNew string
	var copyOld string
	var database string
	var dataSourceType string
	var description string
	var diags diag.Diagnostics
	var elementType string
	var err error
	var folderPath string
	var move bool
	var renamePath string
	var sqlStmt string

	copy = d.Get("copy").(bool)
	copyNew = d.Get("copy_new").(string)
	copyOld = d.Get("copy_old").(string)
	database = d.Get("database").(string)
	dataSourceType = d.Get("data_source_type").(string)
	description = d.Get("description").(string)
	elementType = d.Get("element_type").(string)
	folderPath = d.Get("folder_path").(string)
	move = d.Get("move").(bool)
	renamePath = d.Get("rename_path").(string)
	sqlStmt = fmt.Sprintf(
		`
CONNECT DATABASE %s;
ALTER FOLDER '%s'
`,
		database,
		folderPath,
	)

	if description != "" {
		sqlStmt += fmt.Sprintf(
			"DESCRIPTION '%s';",
			description,
		)
	}

	if renamePath != "" {
		sqlStmt += fmt.Sprintf(
			"RENAME '%s';",
			renamePath,
		)
	}

	if move {
		sqlStmt += fmt.Sprintf(
			"MOVE %s %s %s;",
			elementType,
			TenaryString(dataSourceType != "", dataSourceType, ""),
			move,
		)
	}

	if copy {
		sqlStmt += fmt.Sprintf(
			"COPY %s %s %s AS %s;",
			elementType,
			TenaryString(dataSourceType != "", dataSourceType, ""),
			copyOld,
			copyNew,
		)
	}

	client = meta.(*Client)

	err = client.ExecuteSQL(&sqlStmt)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(d.Get("folder_path").(string))

	diags = readDatabaseFolder(ctx, d, meta)

	return diags
}
