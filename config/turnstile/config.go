package turnstile

import "github.com/crossplane/upjet/pkg/config"

// Configure adds configurations for worker group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_turnstile_widget", func(r *config.Resource) {
		r.References["account_id"] = config.Reference{
			Type: "github.com/ecoupal-believe/provider-cloudflare/apis/account/v1alpha1.Account",
		}
	})
}
