package tunnel

import "github.com/crossplane/upjet/pkg/config"

// Configure adds configurations for worker group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_tunnel", func(r *config.Resource) {
		r.ShortGroup = "Tunnel"
		r.Kind = "Tunnel"
		r.References["account_id"] = config.Reference{
			Type: "github.com/ecoupal-believe/provider-cloudflare/apis/account/v1alpha1.Account",
		}
	})
	p.AddResourceConfigurator("cloudflare_tunnel_config", func(r *config.Resource) {
		r.References["account_id"] = config.Reference{
			Type: "github.com/ecoupal-believe/provider-cloudflare/apis/account/v1alpha1.Account",
		}
		r.References["tunnel_id"] = config.Reference{
			Type: "Tunnel",
		}
	})
	p.AddResourceConfigurator("cloudflare_tunnel_route", func(r *config.Resource) {
		r.References["account_id"] = config.Reference{
			Type: "github.com/ecoupal-believe/provider-cloudflare/apis/account/v1alpha1.Account",
		}
		r.References["tunnel_id"] = config.Reference{
			Type: "Tunnel",
		}
	})
	p.AddResourceConfigurator("cloudflare_tunnel_virtual_network", func(r *config.Resource) {
		r.References["account_id"] = config.Reference{
			Type: "github.com/ecoupal-believe/provider-cloudflare/apis/account/v1alpha1.Account",
		}
	})
}
