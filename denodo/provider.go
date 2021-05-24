package denodo

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/url"

	_ "github.com/lib/pq"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"database": {
				Type: schema.TypeString,
				Required: true,
				DefaultFunc: schema.EnvDefaultFunc("DENODO_DATABASE", "admin"),
			},
			"host": {
				Type: schema.TypeString,
				Required: true,
				DefaultFunc: schema.EnvDefaultFunc("DENODO_HOST", nil),
			},
			"password": {
				Type: schema.TypeString,
				Required: true,
				DefaultFunc: schema.EnvDefaultFunc("DENODO_PASSWORD", "admin"),
			},
			"port": {
				Type: schema.TypeInt,
				Required: true,
				DefaultFunc: schema.EnvDefaultFunc("DENODO_PORT",9996),
			},
			"sslmode": {
				Type: schema.TypeString,
				Required: true,
				DefaultFunc: schema.EnvDefaultFunc("DENODO_SSL_MODE", "perfer"),
			}
			"username": {
				Type: schema.TypeString,
				Required: true,
				DefaultFunc: schema.EnvDefaultFunc("DENODO_USERNAME", "admin"),
			},
		},
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	denodoConnUrl := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?%s",
		d.Get("username").(string),
		url.QueryEscape(d.Get("password").(string)),
		d.Get("host").(string),
		d.Get("port").(string),
		d.Get("database").(string),
		d.Get("connOpts").(string),
	)
	return sql.Open("pgx", denodoConnUrl)
}
