package khulnasoft

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/khulnasoft/terraform-provider-khulnasoft/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGroup() *schema.Resource {
	return &schema.Resource{
		Description: "The `khulnasoft_group` resource manages your groups within Khulnasoft.\n\n" +
			"The Groups created must have at least one Role that is already " +
			"present within Khulnasoft.",
		Create: resourceGroupCreate,
		Read:   resourceGroupRead,
		Update: resourceGroupUpdate,
		Delete: resourceGroupDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"group_id": {
				Type:        schema.TypeInt,
				Description: "The ID of the created group.",
				Computed:    true,
				ForceNew:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "The desired name of the group.",
				Required:    true,
			},
			"created": {
				Type:        schema.TypeString,
				Description: "The creation date of the group.",
				Computed:    true,
			},
		},
	}
}

func resourceGroupCreate(d *schema.ResourceData, m interface{}) error {
	ac := m.(*client.Client)

	group := client.Group{
		Name: d.Get("name").(string),
	}

	err := ac.CreateGroup(&group)
	if err != nil {
		return err
	}
	d.Set("group_id", group.Id)

	d.SetId(fmt.Sprintf("%v", group.Id))
	return resourceGroupRead(d, m)
}

func resourceGroupRead(d *schema.ResourceData, m interface{}) error {
	ac := m.(*client.Client)

	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return nil
	}
	r, err := ac.GetGroup(id)

	if err != nil {
		if strings.Contains(fmt.Sprintf("%s", err), "404") {
			d.SetId("")
			return nil
		}
		return err
	}

	d.Set("name", r.Name)
	d.Set("group_id", r.Id)
	d.Set("created", r.Created)
	d.SetId(fmt.Sprintf("%v", id))

	return nil
}

func resourceGroupUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(*client.Client)

	if d.HasChanges("name") {

		Group := client.Group{
			Name: d.Get("name").(string),
			Id:   d.Get("group_id").(int),
		}

		err := c.UpdateGroup(&Group)
		if err != nil {
			log.Println("[DEBUG]  error while updating Group: ", err)
			return err
		}
	}
	return nil
}

func resourceGroupDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*client.Client)
	id := d.Id()
	err := c.DeleteGroup(id)
	log.Println(err)
	if err == nil {
		d.SetId("")
	} else {
		log.Println("[DEBUG]  error deleting Group: ", err)
		return err
	}
	//d.SetId("")

	return err
}
