package gridscale

import (
	"bitbucket.org/gridscale/gsclient-go"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceGridscaleNetwork() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGridscaleNetworkRead,

		Schema: map[string]*schema.Schema{},
	}
}

func dataSourceGridscaleNetworkRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*gsclient.Client)

	id := d.Get("resource_id").(string)

	storage, err := client.GetNetwork(id)

	if err == nil {
		d.SetId(storage.Properties.ObjectUuid)
	}

	return err
}
