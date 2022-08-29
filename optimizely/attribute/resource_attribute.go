package attribute

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceAttribute() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"archived": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether or not the Attribute has been archived.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A short description of the Attribute.",
			},
			"key": {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "Unique string identifier for this Attribute within the project.",
				ValidateDiagFunc: ValidateKey,
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the Attribute. For Full Stack projects, the name will be set to the value of the key.",
			},
			"project_id": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of the project the Attribute belongs to.",
			},
		},
		CreateContext: createAttributeContext,
		ReadContext:   readAttributeContext,
		UpdateContext: updateAttributeContext,
		DeleteContext: deleteAttributeContext,
	}
}
