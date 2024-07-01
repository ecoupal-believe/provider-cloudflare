package worker

import "github.com/crossplane/upjet/pkg/config"

const (
	shortGroupName = "worker"
)

// Configure adds configurations for worker group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_queue", func(r *config.Resource) {
		r.ShortGroup = shortGroupName
		r.Kind = "Queue"
		r.References["account_id"] = config.Reference{
			Type: "github.com/ecoupal-believe/provider-cloudflare/apis/account/v1alpha1.Account",
		}
	})
	p.AddResourceConfigurator("cloudflare_worker_cron_trigger", func(r *config.Resource) {
		r.ShortGroup = shortGroupName
		r.References["account_id"] = config.Reference{
			Type: "github.com/ecoupal-believe/provider-cloudflare/apis/account/v1alpha1.Account",
		}
		r.References["script_name"] = config.Reference{
			Type: "Script",
		}
	})
	p.AddResourceConfigurator("cloudflare_worker_domain", func(r *config.Resource) {
		r.ShortGroup = shortGroupName
		r.References["account_id"] = config.Reference{
			Type: "github.com/ecoupal-believe/provider-cloudflare/apis/account/v1alpha1.Account",
		}
		r.References["zone_id"] = config.Reference{
			Type: "github.com/ecoupal-believe/provider-cloudflare/apis/zone/v1alpha1.Zone",
		}
	})
	p.AddResourceConfigurator("cloudflare_worker_route", func(r *config.Resource) {
		r.ShortGroup = shortGroupName
		r.References["script_name"] = config.Reference{
			Type: "Script",
		}
		r.References["zone_id"] = config.Reference{
			Type: "github.com/ecoupal-believe/provider-cloudflare/apis/zone/v1alpha1.Zone",
		}
	})
	p.AddResourceConfigurator("cloudflare_worker_script", func(r *config.Resource) {
		r.ShortGroup = shortGroupName
		r.References["account_id"] = config.Reference{
			Type: "github.com/ecoupal-believe/provider-cloudflare/apis/account/v1alpha1.Account",
		}
		r.References["kv_namespace_binding.namespace_id"] = config.Reference{
			Type: "KVNamespace",
		}
		r.References["dispatch_namespace"] = config.Reference{
			Type: "WorkersForPlatformsNamespace",
		}
	})
	p.AddResourceConfigurator("cloudflare_worker_secret", func(r *config.Resource) {
		r.ShortGroup = shortGroupName
		r.References["script_name"] = config.Reference{
			Type: "Script",
		}
		r.References["account_id"] = config.Reference{
			Type: "github.com/ecoupal-believe/provider-cloudflare/apis/account/v1alpha1.Account",
		}
	})
	p.AddResourceConfigurator("cloudflare_workers_for_platforms_namespace", func(r *config.Resource) {
		r.ShortGroup = shortGroupName
		r.Kind = "WorkersForPlatformsNamespace"
		r.References["account_id"] = config.Reference{
			Type: "github.com/ecoupal-believe/provider-cloudflare/apis/account/v1alpha1.Account",
		}
	})
	p.AddResourceConfigurator("cloudflare_workers_kv_namespace", func(r *config.Resource) {
		r.ShortGroup = shortGroupName
		r.Kind = "KVNamespace"
		r.References["account_id"] = config.Reference{
			Type: "github.com/ecoupal-believe/provider-cloudflare/apis/account/v1alpha1.Account",
		}
	})
	p.AddResourceConfigurator("cloudflare_workers_kv", func(r *config.Resource) {
		r.ShortGroup = shortGroupName
		r.Kind = "KV"
		r.References["account_id"] = config.Reference{
			Type: "github.com/ecoupal-believe/provider-cloudflare/apis/account/v1alpha1.Account",
		}
		r.References["namespace_id"] = config.Reference{
			Type: "KVNamespace",
		}
	})
}
