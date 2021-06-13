package denodo

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ConfigureContextFunc: providerConfigure,
		DataSourcesMap: map[string]*schema.Resource{
			"denodo_jdbc_data_source_object": dataSourceJDBCDataSourceObject(),
			"denodo_object":                  dataSourceObject(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"denodo_base_view":        resourceBaseView(),
			"denodo_dervived_view":    resourceDerivedView(),
			"denodo_database":         resourceDatabase(),
			"denodo_database_folder":  resourceDatabaseFolder(),
			"denodo_database_role":    resourceDatabaseRole(),
			"denodo_jdbc_data_source": resourceJDBCDataSource(),
			"denodo_user":             resourceUser(),
		},
		Schema: map[string]*schema.Schema{
			"database": &schema.Schema{
				DefaultFunc: schema.EnvDefaultFunc("DENODO_DATABASE", ""),
				Description: "Database initial connection will be made too.",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"host": &schema.Schema{
				DefaultFunc: schema.EnvDefaultFunc("DENODO_HOST", nil),
				Description: "Host to the Denodo platform.",
				Required:    true,
				Type:        schema.TypeString,
			},
			"password": &schema.Schema{
				DefaultFunc: schema.EnvDefaultFunc("DENODO_PASSWORD", "admin"),
				Description: "Password used to connect to the Denodo platform.",
				Required:    true,
				Sensitive:   true,
				Type:        schema.TypeString,
			},
			"port": &schema.Schema{
				DefaultFunc: schema.EnvDefaultFunc("DENODO_PORT", 9996),
				Description: "Port to the Denodo platform.",
				Required:    true,
				Type:        schema.TypeInt,
			},
			"ssl_mode": &schema.Schema{
				DefaultFunc: schema.EnvDefaultFunc("DENODO_SSL_MODE", "require"),
				Description: "SSL mode used to connect to the Denodo platform. Acceptable values disable, allow, prefer, require",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"username": &schema.Schema{
				DefaultFunc: schema.EnvDefaultFunc("DENODO_USERNAME", "admin"),
				Description: "Username used to connect to the Denodo platform.",
				Required:    true,
				Sensitive:   true,
				Type:        schema.TypeString,
			},
		},
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var client *Client
	var config *Config
	var diags diag.Diagnostics
	var err error
	var retVal int

	client = new(Client)
	config = new(Config)
	config.Database = d.Get("database").(string)
	config.Host = d.Get("host").(string)
	config.Port = d.Get("port").(int)
	config.SslMode = d.Get("ssl_mode").(string)

	password := d.Get("password").(string)
	username := d.Get("username").(string)

	if (username != "") && (password != "") {
		client, err = config.NewClient(&password, &username)
		if err != nil {
			diags = append(
				diags,
				diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "Unable to create Denodo client",
					Detail:   err.Error(),
				},
			)
			return nil, diags
		}

		// Checking connection is good
		err = client.Connection.QueryRowContext(ctx, "SELECT 1 RET_VAL").Scan(
			&retVal,
		)
		if err != nil {
			diags = append(
				diags,
				diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "Unable to authenticate to Denodo client",
					Detail:   err.Error(),
				},
			)
			return nil, diags
		}

		return client, diags
	}

	return client, diags
}
