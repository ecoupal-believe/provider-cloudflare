package custom

import "github.com/crossplane/upjet/pkg/config"

// Configure adds configurations for custom group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_custom_pages", func(r *config.Resource) {
		r.References["account_id"] = config.Reference{
			Type: "github.com/ecoupal-believe/provider-cloudflare/apis/account/v1alpha1.Account",
		}
		r.References["zone_id"] = config.Reference{
			Type: "github.com/ecoupal-believe/provider-cloudflare/apis/zone/v1alpha1.Zone",
		}
	})
	p.AddResourceConfigurator("cloudflare_custom_ssl", func(r *config.Resource) {
		r.Kind = "SSL"
		r.References["zone_id"] = config.Reference{
			Type: "github.com/ecoupal-believe/provider-cloudflare/apis/zone/v1alpha1.Zone",
		}
	})
}
