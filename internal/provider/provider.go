// Generated by LIBLAB | https://liblab.com

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/liblab-sdk/pkg/client"
	"github.com/liblab-sdk/pkg/clientconfig"

	"github.com/terraform-provider-liblab/internal/provider/token"
)

// Ensure Provider satisfies various provider interfaces.
var _ provider.Provider = &Provider{}

type Provider struct {
	version string
}

type liblabProviderModel struct {
	Host      types.String `tfsdk:"host"`
	AuthToken types.String `tfsdk:"auth_token"`
}

func (p *Provider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "liblab"
	resp.Version = "0.0.1"
}

func (p *Provider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"host": schema.StringAttribute{
				Required:    true,
				Sensitive:   false,
				Description: "The API host.",
			},
			"auth_token": schema.StringAttribute{
				Required:    true,
				Sensitive:   true,
				Description: "The authentication token.",
			},
		},
	}
}

func (p *Provider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var dataModel liblabProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &dataModel)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if dataModel.Host.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("host"),
			"Unknown Host",
			"Cannot create API client without host.",
		)
		return
	}

	if dataModel.AuthToken.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("auth_token"),
			"Unknown AuthToken",
			"Cannot create API client without auth_token.",
		)
		return
	}

	config := clientconfig.NewConfig()
	config.SetBaseUrl(dataModel.Host.ValueString())
	config.SetAccessToken(dataModel.AuthToken.ValueString())
	apiClient := client.NewClient(config)

	// Example of setting the client in resp
	resp.DataSourceData = apiClient
	resp.ResourceData = apiClient
}

func (p *Provider) Resources(ctx context.Context) []func() resource.Resource {
	resources := []func() resource.Resource{}
	resources = append(resources, token.NewTokenResource)
	return resources
}

func (p *Provider) DataSources(ctx context.Context) []func() datasource.DataSource {
	dataSources := []func() datasource.DataSource{}
	return dataSources
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &Provider{
			version: version,
		}
	}
}