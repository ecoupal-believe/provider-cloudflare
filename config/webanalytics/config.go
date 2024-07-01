package webanalytics

import "github.com/crossplane/upjet/pkg/config"

// Configure adds configurations for worker group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_web_analytics_site", func(r *config.Resource) {
		r.ShortGroup = "WebAnalytics"
		r.Kind = "Site"
		r.References["account_id"] = config.Reference{
			Type: "github.com/ecoupal-believe/provider-cloudflare/apis/account/v1alpha1.Account",
		}
	})
	p.AddResourceConfigurator("cloudflare_web_analytics_rule", func(r *config.Resource) {
		r.ShortGroup = "WebAnalytics"
		r.Kind = "Rule"
		r.References["account_id"] = config.Reference{
			Type: "github.com/ecoupal-believe/provider-cloudflare/apis/account/v1alpha1.Account",
		}
		//TODO fix me
		r.References["ruleset_id"] = config.Reference{
			Type: "Site",
		}
	})
}
