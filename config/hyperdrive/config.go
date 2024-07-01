package hyperdrive

import "github.com/crossplane/upjet/pkg/config"

// Configure adds configurations for lists group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_hyperdrive_config", func(r *config.Resource) {
		r.References["account_id"] = config.Reference{
			Type: "github.com/ecoupal-believe/provider-cloudflare/apis/account/v1alpha1.Account",
		}
	})
}
