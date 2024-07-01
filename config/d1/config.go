package d1

import "github.com/crossplane/upjet/pkg/config"

// Configure adds configurations for d1 group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_d1_database", func(r *config.Resource) {
		r.References["account_id"] = config.Reference{
			Type: "github.com/ecoupal-believe/provider-cloudflare/apis/account/v1alpha1.Account",
		}
	})
}
