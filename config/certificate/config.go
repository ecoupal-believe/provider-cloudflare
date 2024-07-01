package certificate

import "github.com/crossplane/upjet/pkg/config"

// Configure adds configurations for certificate group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_certificate_pack", func(r *config.Resource) {
		r.References["zone_id"] = config.Reference{
			Type: "github.com/ecoupal-believe/provider-cloudflare/apis/zone/v1alpha1.Zone",
		}
	})
}
