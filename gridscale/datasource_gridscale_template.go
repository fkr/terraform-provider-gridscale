package gridscale

import (
	"bitbucket.org/gridscale/gsclient-go"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	"log"
)

func dataSourceGridscaleTemplate() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGridscaleTemplateRead,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				Description:  "name of the domain",
				ValidateFunc: validation.NoZeroValues,
			},
		},
	}
}

func dataSourceGridscaleTemplateRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*gsclient.Client)

	name := d.Get("name").(string)

	template, err := client.GetTemplateByName(name)

	if err == nil {
		d.SetId(template.Properties.ObjectUuid)
		log.Printf("Found template with key: %v", template.Properties.ObjectUuid)
	}

	return err
}
