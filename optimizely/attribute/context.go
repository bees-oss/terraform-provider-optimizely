package attribute

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Attribute struct {
	Archived      bool   `json:"archived"`
	ConditionType string `json:"condition_type"`
	Description   string `json:"description"`
	ID            int64  `json:"id"`
	Key           string `json:"key"`
	LastModified  string `json:"last_modified"`
	Name          string `json:"name"`
	ProjectId     int    `json:"project_id"`
}

type Client interface {
	GetAttribute(id int) (*Attribute, error)
	CreateAttribute(archived bool, description, key, name string, project_id int) (string, error)
	UpdateAttribute(id int, archived bool, description, key, name string) error
	DeleteAttribute(id int) error
}

func readAttributeContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(Client)
	id, err := strconv.Atoi(d.Id())
	var diags diag.Diagnostics

	if err != nil {
		return diag.FromErr(err)
	}

	if instance, err := client.GetAttribute(id); err != nil {
		return diag.FromErr(err)
	} else {
		d.Set("archived", instance.Archived)
		d.Set("description", instance.Description)
		d.Set("key", instance.Key)
		d.Set("name", instance.Name)
		d.Set("project_id", instance.ProjectId)
	}

	return diags
}

func createAttributeContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(Client)

	archived := d.Get("archived").(bool)
	description := d.Get("description").(string)
	key := d.Get("key").(string)
	name := d.Get("name").(string)
	project_id := d.Get("project_id").(int)

	if id, err := client.CreateAttribute(archived, description, key, name, project_id); err != nil {
		return diag.FromErr(err)
	} else {
		d.SetId(id)
	}

	return readAttributeContext(ctx, d, m)
}

func updateAttributeContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(Client)

	id, err := strconv.Atoi(d.Id())

	if err != nil {
		return diag.FromErr(err)
	}

	archived := d.Get("archived").(bool)
	description := d.Get("description").(string)
	key := d.Get("key").(string)
	name := d.Get("name").(string)

	if err := client.UpdateAttribute(id, archived, description, key, name); err != nil {
		return diag.FromErr(err)
	}

	return readAttributeContext(ctx, d, m)
}

func deleteAttributeContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(Client)

	id, err := strconv.Atoi(d.Id())

	if err != nil {
		return diag.FromErr(err)
	}

	if err := client.DeleteAttribute(id); err != nil {
		return diag.FromErr(err)
	}

	return readAttributeContext(ctx, d, m)
}
