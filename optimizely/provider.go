package optimizely

import (
	"context"

	"github.com/bees-oss/optimizely-terraform-provider/optimizely/attribute"
	"github.com/bees-oss/optimizely-terraform-provider/optimizely/audience"
	"github.com/bees-oss/optimizely-terraform-provider/optimizely/client"
	"github.com/bees-oss/optimizely-terraform-provider/optimizely/environment"
	"github.com/bees-oss/optimizely-terraform-provider/optimizely/flag"
	"github.com/bees-oss/optimizely-terraform-provider/optimizely/project"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host": {
				Type:     schema.TypeString,
				Required: true,
			},
			"token": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"optimizely_feature":   flag.ResourceFeature(),
			"optimizely_audience":  audience.ResourceAudience(),
			"optimizely_attribute": attribute.ResourceAttribute(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"optimizely_environment": environment.DataSourceEnvironment(),
			"optimizely_project":     project.DataSourceProject(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics

	address := d.Get("host").(string)
	token := d.Get("token").(string)

	optimizelyClient := client.OptimizelyClient{
		Address: address,
		Token:   token,
	}

	return optimizelyClient, diags
}
