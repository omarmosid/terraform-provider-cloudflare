package zaraz

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func (r *ZarazConfigResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: heredoc.Doc(`
			The [Zaraz Config](https://developers.cloudflare.com/zaraz/) resource allows you to manage your Cloudflare Zaraz config.
	`),

		Attributes: map[string]schema.Attribute{
			consts.AccountIDSchemaKey: schema.StringAttribute{
				MarkdownDescription: consts.AccountIDSchemaDescription,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.ConflictsWith(
						path.Expression(path.MatchRoot((consts.ZoneIDSchemaKey))),
					),
				},
			},
			consts.ZoneIDSchemaKey: schema.StringAttribute{
				MarkdownDescription: consts.ZoneIDSchemaDescription,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.ConflictsWith(
						path.Expression(path.MatchRoot((consts.AccountIDSchemaKey))),
					),
				},
			},
		},
		Blocks: map[string]schema.Block{
			"config": schema.ListNestedBlock{
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"debug_key": schema.StringAttribute{
							Required: true,
						},
						"tools": schema.MapNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"enabled": schema.BoolAttribute{
										Required: true,
										// ... potentially other fields ...
									},
								},
							},
							Required: true,
							// ... potentially other fields ...
						},
					},
				},
			},
		},
	}
}

// func configSchema() schema.MapAttribute {
// 	return schema.MapAttribute{
// 		ElementType: types.StringType,
// 		Required:    true,
// 		// ... potentially other fields ...
// 	}
// }

// func configToolSchema() schema.MapNestedAttribute {
// 	return schema.MapNestedAttribute{
// 		NestedObject: schema.NestedAttributeObject{
// 			Attributes: map[string]schema.Attribute{
// 				"enabled": schema.BoolAttribute{
// 					Required: true,
// 				},
// 				"name": schema.StringAttribute{
// 					Required: true,
// 				},
// 			},
// 		},
// 		Optional: true,
// 	}
// }
