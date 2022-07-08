package filesystem

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFile() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceFileCreate,
		ReadContext:   resourceFileRead,
		UpdateContext: resourceFileUpdate,
		DeleteContext: resourceFileDelete,
		Schema: map[string]*schema.Schema{
			"path": {
				Type:        schema.TypeString,
				ForceNew:    true,
				Required:    true,
				Description: "The path to save the file.",
			},
			"content": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The content of the file.",
			},
		},
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceFileCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	fm := m.(*FileManager)
	var diags diag.Diagnostics

	path := d.Get("path").(string)
	content := d.Get("content").(string)
	_, err := fm.CreateFile(path, content)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(path)
	resourceFileRead(ctx, d, m)
	return diags
}

func resourceFileRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	fm := m.(*FileManager)
	var diags diag.Diagnostics

	path := d.Id()
	fi, err := fm.ReadFile(path)
	if err != nil {
		return diag.FromErr(err)
	}
	d.Set("content", fi.FileContent)
	return diags
}

func resourceFileUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	fm := m.(*FileManager)

	path := d.Id()
	content := d.Get("content").(string)
	_, err := fm.UpdateFile(path, content)
	if err != nil {
		return diag.FromErr(err)
	}
	return resourceFileRead(ctx, d, m)
}

func resourceFileDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	fm := m.(*FileManager)
	var diags diag.Diagnostics

	path := d.Id()
	err := fm.DeleteFile(path)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")
	return diags
}
