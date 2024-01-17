package zaraz

import "github.com/hashicorp/terraform-plugin-framework/types"

type ZarazConfigModel struct {
	AccountId types.String `tfsdk:"account_id"`
	ZoneID    types.String `tfsdk:"zone_id"`
	DebugKey  types.String `tfsdk:"debugKey"`
}
