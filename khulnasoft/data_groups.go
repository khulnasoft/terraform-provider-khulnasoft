package khulnasoft

import (
	"log"

	"github.com/khulnasoft/terraform-provider-khulnasoft/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGroups() *schema.Resource {
	return &schema.Resource{
		Description: "The data source `khulnasoft_groups` provides a method to query all groups within the Khulnasoft CSPM" +
			"group database. The fields returned from this query are detailed in the Schema section below.",
		Read: dataGroupRead,
		Schema: map[string]*schema.Schema{
			"groups": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"group_id": {
							Type:     schema.TypeString,
							Description: "The ID of the created group.",
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Description: "The desired name of the group.",
							Computed: true,
						},
						"created": {
							Type:     schema.TypeString,
							Description: "The creation date of the group.",
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataGroupRead(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG]  inside dataGroup")
	c := m.(*client.Client)
	result, err := c.GetGroups()
	if err == nil {
		groups, id := flattenGroupsData(&result)
		d.SetId(id)
		if err := d.Set("groups", groups); err != nil {
			return err
		}
	} else {
		return err
	}

	return nil
}
