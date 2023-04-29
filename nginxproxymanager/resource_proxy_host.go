package nginxproxymanager

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/sander0542/terraform-provider-nginxproxymanager/client"
	"github.com/sander0542/terraform-provider-nginxproxymanager/client/models"
)

var (
	_ resource.Resource                   = &proxyHostResource{}
	_ resource.ResourceWithConfigure      = &proxyHostResource{}
	_ resource.ResourceWithValidateConfig = &proxyHostResource{}
	_ resource.ResourceWithImportState    = &proxyHostResource{}
)

func NewProxyHostResource() resource.Resource {
	return &proxyHostResource{}
}

type proxyHostResource struct {
	client *client.Client
}

type proxyHostResourceModel struct {
	ID                    types.Int64              `tfsdk:"id"`
	CreatedOn             types.String             `tfsdk:"created_on"`
	ModifiedOn            types.String             `tfsdk:"modified_on"`
	DomainNames           []types.String           `tfsdk:"domain_names"`
	ForwardScheme         types.String             `tfsdk:"forward_scheme"`
	ForwardHost           types.String             `tfsdk:"forward_host"`
	ForwardPort           types.Int64              `tfsdk:"forward_port"`
	CertificateID         types.String             `tfsdk:"certificate_id"`
	SSLForced             types.Bool               `tfsdk:"ssl_forced"`
	HSTSEnabled           types.Bool               `tfsdk:"hsts_enabled"`
	HSTSSubdomains        types.Bool               `tfsdk:"hsts_subdomains"`
	HTTP2Support          types.Bool               `tfsdk:"http2_support"`
	BlockExploits         types.Bool               `tfsdk:"block_exploits"`
	CachingEnabled        types.Bool               `tfsdk:"caching_enabled"`
	AllowWebsocketUpgrade types.Bool               `tfsdk:"allow_websocket_upgrade"`
	AccessListID          types.Int64              `tfsdk:"access_list_id"`
	AdvancedConfig        types.String             `tfsdk:"advanced_config"`
	Enabled               types.Bool               `tfsdk:"enabled"`
	Meta                  types.Map                `tfsdk:"meta"`
	Locations             []proxyHostLocationModel `tfsdk:"location"`
}

type proxyHostLocationModel struct {
	Path           types.String `tfsdk:"path"`
	ForwardScheme  types.String `tfsdk:"forward_scheme"`
	ForwardHost    types.String `tfsdk:"forward_host"`
	ForwardPort    types.Int64  `tfsdk:"forward_port"`
	AdvancedConfig types.String `tfsdk:"advanced_config"`
}

func (r *proxyHostResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_proxy_host"
}

func (r *proxyHostResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Proxy Host data source",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"created_on": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"modified_on": schema.StringAttribute{
				Computed: true,
			},
			"domain_names": schema.ListAttribute{
				Required:    true,
				ElementType: types.StringType,
				Validators: []validator.List{
					listvalidator.SizeAtLeast(1),
				},
			},
			"forward_scheme": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.OneOf("http", "https"),
				},
			},
			"forward_host": schema.StringAttribute{
				Required: true,
			},
			"forward_port": schema.Int64Attribute{
				Required: true,
				Validators: []validator.Int64{
					int64validator.Between(1, 65535),
				},
			},
			"certificate_id": schema.StringAttribute{
				Computed: true,
				Optional: true,
				Default:  stringdefault.StaticString("0"),
			},
			"ssl_forced": schema.BoolAttribute{
				Computed: true,
				Optional: true,
				Default:  booldefault.StaticBool(false),
			},
			"hsts_enabled": schema.BoolAttribute{
				Computed: true,
				Optional: true,
				Default:  booldefault.StaticBool(false),
			},
			"hsts_subdomains": schema.BoolAttribute{
				Computed: true,
				Optional: true,
				Default:  booldefault.StaticBool(false),
			},
			"http2_support": schema.BoolAttribute{
				Computed: true,
				Optional: true,
				Default:  booldefault.StaticBool(false),
			},
			"block_exploits": schema.BoolAttribute{
				Computed: true,
				Optional: true,
				Default:  booldefault.StaticBool(false),
			},
			"caching_enabled": schema.BoolAttribute{
				Computed: true,
				Optional: true,
				Default:  booldefault.StaticBool(false),
			},
			"allow_websocket_upgrade": schema.BoolAttribute{
				Computed: true,
				Optional: true,
				Default:  booldefault.StaticBool(true),
			},
			"access_list_id": schema.Int64Attribute{
				Computed: true,
				Optional: true,
				Default:  int64default.StaticInt64(0),
			},
			"advanced_config": schema.StringAttribute{
				Computed: true,
				Optional: true,
				Default:  stringdefault.StaticString(""),
			},
			"enabled": schema.BoolAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"meta": schema.MapAttribute{
				ElementType: types.StringType,
				Computed:    true,
				PlanModifiers: []planmodifier.Map{
					mapplanmodifier.UseStateForUnknown(),
				},
			},
		},
		Blocks: map[string]schema.Block{
			"location": schema.ListNestedBlock{
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"path": schema.StringAttribute{
							Required: true,
						},
						"forward_scheme": schema.StringAttribute{
							Required: true,
							Validators: []validator.String{
								stringvalidator.OneOf("http", "https"),
							},
						},
						"forward_host": schema.StringAttribute{
							Required: true,
						},
						"forward_port": schema.Int64Attribute{
							Required: true,
							Validators: []validator.Int64{
								int64validator.Between(1, 65535),
							},
						},
						"advanced_config": schema.StringAttribute{
							Computed: true,
							Default:  stringdefault.StaticString(""),
						},
					},
				},
			},
		},
	}
}

func (r *proxyHostResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*client.Client)
}

func (r *proxyHostResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan proxyHostResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var item models.ProxyHostCreate

	item.DomainNames = []string{}
	for _, domainName := range plan.DomainNames {
		item.DomainNames = append(item.DomainNames, domainName.ValueString())
	}
	item.ForwardScheme = plan.ForwardScheme.ValueString()
	item.ForwardHost = plan.ForwardHost.ValueString()
	item.ForwardPort = uint16(plan.ForwardPort.ValueInt64())
	item.CertificateID = plan.CertificateID.ValueString()
	item.SSLForced = plan.SSLForced.ValueBool()
	item.HSTSEnabled = plan.HSTSEnabled.ValueBool()
	item.HSTSSubdomains = plan.HSTSSubdomains.ValueBool()
	item.HTTP2Support = plan.HTTP2Support.ValueBool()
	item.BlockExploits = plan.BlockExploits.ValueBool()
	item.CachingEnabled = plan.CachingEnabled.ValueBool()
	item.AllowWebsocketUpgrade = plan.AllowWebsocketUpgrade.ValueBool()
	item.AccessListID = plan.AccessListID.ValueInt64()
	item.AdvancedConfig = plan.AdvancedConfig.ValueString()
	item.Meta = map[string]string{}
	item.Locations = []models.ProxyHostLocation{}
	for _, location := range plan.Locations {
		item.Locations = append(item.Locations, models.ProxyHostLocation{
			Path:           location.Path.ValueString(),
			ForwardScheme:  location.ForwardScheme.ValueString(),
			ForwardHost:    location.ForwardHost.ValueString(),
			ForwardPort:    uint16(location.ForwardPort.ValueInt64()),
			AdvancedConfig: location.AdvancedConfig.ValueString(),
		})
	}

	proxyHost, err := r.client.CreateProxyHost(&item)
	if err != nil {
		resp.Diagnostics.AddError("Error creating proxy host", "Could not create proxy host, unexpected error: "+err.Error())
		return
	}

	meta, diags := types.MapValueFrom(ctx, types.StringType, proxyHost.Meta.Map())
	resp.Diagnostics.Append(diags...)

	plan.ID = types.Int64Value(proxyHost.ID)
	plan.CreatedOn = types.StringValue(proxyHost.CreatedOn)
	plan.ModifiedOn = types.StringValue(proxyHost.ModifiedOn)
	plan.Enabled = types.BoolValue(proxyHost.Enabled.Bool())
	plan.Meta = meta

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *proxyHostResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state *proxyHostResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	proxyHost, err := r.client.GetProxyHost(state.ID.ValueInt64Pointer())
	if err != nil {
		resp.Diagnostics.AddError("Error reading proxy host", "Could not read proxy host, unexpected error: "+err.Error())
		return
	}
	if proxyHost == nil {
		state = nil
	} else {
		state.DomainNames = []types.String{}
		for _, domainName := range proxyHost.DomainNames {
			state.DomainNames = append(state.DomainNames, types.StringValue(domainName))
		}
		state.Locations = []proxyHostLocationModel{}
		for _, location := range proxyHost.Locations {
			state.Locations = append(state.Locations, proxyHostLocationModel{
				Path:           types.StringValue(location.Path),
				ForwardScheme:  types.StringValue(location.ForwardScheme),
				ForwardHost:    types.StringValue(location.ForwardHost),
				ForwardPort:    types.Int64Value(int64(location.ForwardPort)),
				AdvancedConfig: types.StringValue(location.AdvancedConfig),
			})
		}
		meta, diags := types.MapValueFrom(ctx, types.StringType, proxyHost.Meta.Map())
		resp.Diagnostics.Append(diags...)

		state.ID = types.Int64Value(proxyHost.ID)
		state.CreatedOn = types.StringValue(proxyHost.CreatedOn)
		state.ModifiedOn = types.StringValue(proxyHost.ModifiedOn)
		state.ForwardScheme = types.StringValue(proxyHost.ForwardScheme)
		state.ForwardHost = types.StringValue(proxyHost.ForwardHost)
		state.ForwardPort = types.Int64Value(int64(proxyHost.ForwardPort))
		state.CertificateID = types.StringValue(string(proxyHost.CertificateID))
		state.SSLForced = types.BoolValue(proxyHost.SSLForced.Bool())
		state.HSTSEnabled = types.BoolValue(proxyHost.HSTSEnabled.Bool())
		state.HSTSSubdomains = types.BoolValue(proxyHost.HSTSSubdomains.Bool())
		state.HTTP2Support = types.BoolValue(proxyHost.HTTP2Support.Bool())
		state.BlockExploits = types.BoolValue(proxyHost.BlockExploits.Bool())
		state.CachingEnabled = types.BoolValue(proxyHost.CachingEnabled.Bool())
		state.AllowWebsocketUpgrade = types.BoolValue(proxyHost.AllowWebsocketUpgrade.Bool())
		state.AccessListID = types.Int64Value(proxyHost.AccessListID)
		state.AdvancedConfig = types.StringValue(proxyHost.AdvancedConfig)
		state.Enabled = types.BoolValue(proxyHost.Enabled.Bool())
		state.Meta = meta
	}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *proxyHostResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan proxyHostResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var item models.ProxyHostUpdate

	item.DomainNames = []string{}
	for _, domainName := range plan.DomainNames {
		item.DomainNames = append(item.DomainNames, domainName.ValueString())
	}
	item.ForwardScheme = plan.ForwardScheme.ValueString()
	item.ForwardHost = plan.ForwardHost.ValueString()
	item.ForwardPort = uint16(plan.ForwardPort.ValueInt64())
	item.CertificateID = plan.CertificateID.ValueString()
	item.SSLForced = plan.SSLForced.ValueBool()
	item.HSTSEnabled = plan.HSTSEnabled.ValueBool()
	item.HSTSSubdomains = plan.HSTSSubdomains.ValueBool()
	item.HTTP2Support = plan.HTTP2Support.ValueBool()
	item.BlockExploits = plan.BlockExploits.ValueBool()
	item.CachingEnabled = plan.CachingEnabled.ValueBool()
	item.AllowWebsocketUpgrade = plan.AllowWebsocketUpgrade.ValueBool()
	item.AccessListID = plan.AccessListID.ValueInt64()
	item.AdvancedConfig = plan.AdvancedConfig.ValueString()
	item.Meta = map[string]string{}
	item.Locations = []models.ProxyHostLocation{}
	for _, location := range plan.Locations {
		item.Locations = append(item.Locations, models.ProxyHostLocation{
			Path:           location.Path.ValueString(),
			ForwardScheme:  location.ForwardScheme.ValueString(),
			ForwardHost:    location.ForwardHost.ValueString(),
			ForwardPort:    uint16(location.ForwardPort.ValueInt64()),
			AdvancedConfig: location.AdvancedConfig.ValueString(),
		})
	}

	proxyHost, err := r.client.UpdateProxyHost(plan.ID.ValueInt64Pointer(), &item)
	if err != nil {
		resp.Diagnostics.AddError("Error updating proxy host", "Could not update proxy host, unexpected error: "+err.Error())
		return
	}

	meta, diags := types.MapValueFrom(ctx, types.StringType, proxyHost.Meta.Map())
	resp.Diagnostics.Append(diags...)

	plan.ID = types.Int64Value(proxyHost.ID)
	plan.DomainNames = []types.String{}
	for _, domainName := range proxyHost.DomainNames {
		plan.DomainNames = append(plan.DomainNames, types.StringValue(domainName))
	}
	plan.Locations = []proxyHostLocationModel{}
	for _, location := range proxyHost.Locations {
		plan.Locations = append(plan.Locations, proxyHostLocationModel{
			Path:           types.StringValue(location.Path),
			ForwardScheme:  types.StringValue(location.ForwardScheme),
			ForwardHost:    types.StringValue(location.ForwardHost),
			ForwardPort:    types.Int64Value(int64(location.ForwardPort)),
			AdvancedConfig: types.StringValue(location.AdvancedConfig),
		})
	}
	plan.CreatedOn = types.StringValue(proxyHost.CreatedOn)
	plan.ModifiedOn = types.StringValue(proxyHost.ModifiedOn)
	plan.ForwardScheme = types.StringValue(proxyHost.ForwardScheme)
	plan.ForwardHost = types.StringValue(proxyHost.ForwardHost)
	plan.ForwardPort = types.Int64Value(int64(proxyHost.ForwardPort))
	plan.CertificateID = types.StringValue(string(proxyHost.CertificateID))
	plan.SSLForced = types.BoolValue(proxyHost.SSLForced.Bool())
	plan.HSTSEnabled = types.BoolValue(proxyHost.HSTSEnabled.Bool())
	plan.HSTSSubdomains = types.BoolValue(proxyHost.HSTSSubdomains.Bool())
	plan.HTTP2Support = types.BoolValue(proxyHost.HTTP2Support.Bool())
	plan.BlockExploits = types.BoolValue(proxyHost.BlockExploits.Bool())
	plan.CachingEnabled = types.BoolValue(proxyHost.CachingEnabled.Bool())
	plan.AllowWebsocketUpgrade = types.BoolValue(proxyHost.AllowWebsocketUpgrade.Bool())
	plan.AccessListID = types.Int64Value(proxyHost.AccessListID)
	plan.AdvancedConfig = types.StringValue(proxyHost.AdvancedConfig)
	plan.Enabled = types.BoolValue(proxyHost.Enabled.Bool())
	plan.Meta = meta

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *proxyHostResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state proxyHostResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.DeleteProxyHost(state.ID.ValueInt64Pointer())
	if err != nil {
		resp.Diagnostics.AddError("Error deleting proxy host", "Could not delete proxy host, unexpected error: "+err.Error())
		return
	}
}

func (r *proxyHostResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data proxyHostResourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if data.SSLForced.ValueBool() && data.CertificateID.ValueString() == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("ssl_forced"),
			"Certificate ID is required when SSL is forced",
			"Certificate ID is required when SSL is forced")
	}

	if data.HTTP2Support.ValueBool() && data.CertificateID.ValueString() == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("http2_support"),
			"Certificate ID is required when HTTP/2 Support is enabled",
			"Certificate ID is required when HTTP/2 Support is enabled")
	}

	if data.HSTSEnabled.ValueBool() && !data.SSLForced.ValueBool() {
		resp.Diagnostics.AddAttributeError(
			path.Root("hsts_enabled"),
			"SSL must be forced when HSTS is enabled",
			"SSL must be forced when HSTS is enabled")
	}

	if data.HSTSSubdomains.ValueBool() && !data.HSTSEnabled.ValueBool() {
		resp.Diagnostics.AddAttributeError(
			path.Root("hsts_subdomains"),
			"HSTS must be enabled when HSTS Subdomains is enabled",
			"HSTS must be enabled when HSTS Subdomains is enabled")
	}
}

func (r *proxyHostResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}