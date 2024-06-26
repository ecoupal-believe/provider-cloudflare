package magic

import "github.com/crossplane/upjet/pkg/config"

const (
	shortGroupName = "Magic"
)

// Configure adds configurations for magic group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_gre_tunnel", func(r *config.Resource) {
		r.ShortGroup = shortGroupName
		r.Kind = "GRETunnel"
		r.References["account_id"] = config.Reference{
			Type: "github.com/ecoupal-believe/provider-cloudflare/apis/account/v1alpha1.Account",
		}
	})
	p.AddResourceConfigurator("cloudflare_ipsec_tunnel", func(r *config.Resource) {
		r.ShortGroup = shortGroupName
		r.Kind = "IPsecTunnel"
		r.References["account_id"] = config.Reference{
			Type: "github.com/ecoupal-believe/provider-cloudflare/apis/account/v1alpha1.Account",
		}
	})
	p.AddResourceConfigurator("cloudflare_magic_firewall_ruleset", func(r *config.Resource) {
		r.ShortGroup = shortGroupName
		r.Kind = "FirewallRuleset"
		r.References["account_id"] = config.Reference{
			Type: "github.com/ecoupal-believe/provider-cloudflare/apis/account/v1alpha1.Account",
		}
	})
	p.AddResourceConfigurator("cloudflare_static_route", func(r *config.Resource) {
		r.ShortGroup = shortGroupName
		r.Kind = "StaticRoute"
		r.References["account_id"] = config.Reference{
			Type: "github.com/ecoupal-believe/provider-cloudflare/apis/account/v1alpha1.Account",
		}
	})
}
