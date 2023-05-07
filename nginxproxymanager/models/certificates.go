package models

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/sander0542/terraform-provider-nginxproxymanager/client/models"
	"github.com/sander0542/terraform-provider-nginxproxymanager/nginxproxymanager/common"
)

type Certificates struct {
	common.IModel[models.CertificateResourceCollection]
	Certificates []Certificate `tfsdk:"certificates"`
}

func (m *Certificates) Load(ctx context.Context, resource *models.CertificateResourceCollection) diag.Diagnostics {
	diags := diag.Diagnostics{}
	m.Certificates = make([]Certificate, len(*resource))
	for i, certificate := range *resource {
		diags.Append(m.Certificates[i].Load(ctx, &certificate)...)
	}

	return diags
}
