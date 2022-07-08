package filesystem

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"secret_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("FILESYSTEM_SECRET_KEY", nil),
				Description: "The secret key will be used to encrypt the file.",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"filesystem_file": resourceFile(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"filesystem_file": dataSourceFile(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	fm := &FileManager{}
	var diags diag.Diagnostics

	secretKey := d.Get("secret_key").(string)
	if secretKey != "" {
		fm.SetSecretKey(secretKey)
	}
	return fm, diags
}
