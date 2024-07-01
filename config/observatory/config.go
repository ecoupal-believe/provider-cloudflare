package observatory

import "github.com/crossplane/upjet/pkg/config"

// Configure adds configurations for observatory group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_observatory_scheduled_test ", func(r *config.Resource) {
		r.References["zone_id"] = config.Reference{
			Type: "github.com/ecoupal-believe/provider-cloudflare/apis/zone/v1alpha1.Zone",
		}
	})
}
