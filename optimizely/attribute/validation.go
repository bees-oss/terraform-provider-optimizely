package attribute

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

func ValidateKey(v interface{}, p cty.Path) diag.Diagnostics {
	value := v.(string)
	var diags diag.Diagnostics
	re := regexp.MustCompile(`^[a-zA-Z0-9_\- .]+$`)

	if len(value) > 64 {
		return diag.FromErr(errors.New(fmt.Sprintf("The key has %d characters", len(value))))
	}

	if !re.MatchString(value) {
		return diag.FromErr(errors.New("The key has invalid characters"))
	}

	return diags
}
