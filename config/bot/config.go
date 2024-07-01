package bot

import "github.com/crossplane/upjet/pkg/config"

// Configure adds configurations for bot management.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_bot_management", func(r *config.Resource) {
		r.Kind = "BotManagement"
		r.References["zone_id"] = config.Reference{
			Type: "github.com/ecoupal-believe/provider-cloudflare/apis/zone/v1alpha1.Zone",
		}
	})
}
