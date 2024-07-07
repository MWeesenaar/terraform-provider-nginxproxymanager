package nginxproxymanager

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/mweesenaar/terraform-provider-nginxproxymanager/client"
	"github.com/mweesenaar/terraform-provider-nginxproxymanager/client/inputs"
	attributes "github.com/mweesenaar/terraform-provider-nginxproxymanager/nginxproxymanager/attributes/resources"
	"github.com/mweesenaar/terraform-provider-nginxproxymanager/nginxproxymanager/common"
	"github.com/mweesenaar/terraform-provider-nginxproxymanager/nginxproxymanager/models"
	"strconv"
)

var (
	_ common.IResource                    = &certificateLetsEncryptResource{}
	_ resource.ResourceWithConfigure      = &certificateLetsEncryptResource{}
	_ resource.ResourceWithValidateConfig = &certificateLetsEncryptResource{}
	_ resource.ResourceWithImportState    = &certificateLetsEncryptResource{}
)

func NewCertificateLetsEncryptResource() resource.Resource {
	b := &common.Resource{Name: "certificate_letsencrypt"}
	r := &certificateLetsEncryptResource{b, nil}
	b.IResource = r
	return r
}

type certificateLetsEncryptResource struct {
	*common.Resource
	client *client.Client
}

func (r *certificateLetsEncryptResource) SchemaImpl(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "SSL Certificates --- Manage a LetsEncrypt certificate.",
		Attributes:  attributes.CertificateCustom,
	}
}

func (r *certificateLetsEncryptResource) Configure(ctx context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*client.Client)
}

func (r *certificateLetsEncryptResource) CreateImpl(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan models.CertificateCustom
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var item inputs.CertificateCustom
	diags.Append(plan.Save(ctx, &item)...)

	certificate, err := r.client.CreateCertificateCustom(ctx, &item)
	if err != nil {
		resp.Diagnostics.AddError("Error creating certificate letsencrypt", "Could not create certificate letsencrypt, unexpected error: "+err.Error())
		return
	}

	resp.Diagnostics.Append(plan.Load(ctx, certificate)...)

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *certificateLetsEncryptResource) ReadImpl(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state *models.CertificateCustom
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	certificate, err := r.client.GetCertificate(ctx, state.ID.ValueInt64Pointer())
	if err != nil {
		resp.Diagnostics.AddError("Error reading certificate letsencrypt", "Could not read certificate letsencrypt, unexpected error: "+err.Error())
		return
	}
	if certificate == nil {
		state = nil
	} else if certificate.Provider != "other" {
		resp.Diagnostics.AddError("Error reading certificate letsencrypt", "Certificate is not a letsencrypt certificate.")
	} else {
		resp.Diagnostics.Append(state.Load(ctx, certificate)...)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *certificateLetsEncryptResource) UpdateImpl(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// There is no update method for certificates, so we delete and recreate
}

func (r *certificateLetsEncryptResource) DeleteImpl(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state *models.CertificateCustom
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.DeleteCertificate(ctx, state.ID.ValueInt64Pointer())
	if err != nil {
		resp.Diagnostics.AddError("Error deleting certificate custom", "Could not delete certificate custom, unexpected error: "+err.Error())
		return
	}
}

func (r *certificateLetsEncryptResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.CertificateCustom

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *certificateLetsEncryptResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	id, err := strconv.ParseInt(req.ID, 10, 64)
	if err != nil {
		resp.Diagnostics.AddError("Error importing certificate custom", "Could not import certificate custom, unexpected error: "+err.Error())
		return
	}

	diags := resp.State.SetAttribute(ctx, path.Root("id"), types.Int64Value(id))
	resp.Diagnostics.Append(diags...)
}
