package filesystem

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFile() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceFileRead,
		Schema: map[string]*schema.Schema{
			"path": {
				Type:     schema.TypeString,
				Required: true,
			},
			"filename": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"size": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"update_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"content": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceFileRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	fm := m.(*FileManager)
	var diags diag.Diagnostics

	path := d.Get("path").(string)
	fi, err := fm.ReadFile(path)
	if err != nil {
		return diag.FromErr(err)
	}
	d.Set("filename", fi.Filename)
	d.Set("size", strconv.FormatInt(fi.FileSize, 10))
	d.Set("update_time", fi.FileUpdateTime)
	d.Set("content", fi.FileContent)

	d.SetId(path) 
	return diags
}
