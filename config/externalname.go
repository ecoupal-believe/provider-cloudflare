// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package config

import "github.com/crossplane/upjet/pkg/config"

// terraformPluginSDKExternalNameConfigs contains all external name configurations
// belonging to Terraform resources to be reconciled under the no-fork
// architecture for this provider.
var terraformPluginSDKExternalNameConfigs = map[string]config.ExternalName{
	"cloudflare_access_application":                              config.IdentifierFromProvider,
	"cloudflare_access_ca_certificate":                           config.IdentifierFromProvider,
	"cloudflare_access_custom_page":                              config.IdentifierFromProvider,
	"cloudflare_access_group":                                    config.IdentifierFromProvider,
	"cloudflare_access_identity_provider":                        config.IdentifierFromProvider,
	"cloudflare_access_keys_configuration":                       config.IdentifierFromProvider,
	"cloudflare_access_mutual_tls_certificate":                   config.IdentifierFromProvider,
	"cloudflare_access_mutual_tls_hostname_settings":             config.IdentifierFromProvider,
	"cloudflare_access_organization":                             config.IdentifierFromProvider,
	"cloudflare_access_policy":                                   config.IdentifierFromProvider,
	"cloudflare_access_rule":                                     config.IdentifierFromProvider,
	"cloudflare_access_rule_migrate":                             config.IdentifierFromProvider,
	"cloudflare_access_service_token":                            config.IdentifierFromProvider,
	"cloudflare_access_tag":                                      config.IdentifierFromProvider,
	"cloudflare_account":                                         config.IdentifierFromProvider,
	"cloudflare_account_member":                                  config.IdentifierFromProvider,
	"cloudflare_api_token":                                       config.IdentifierFromProvider,
	"cloudflare_address_map":                                     config.IdentifierFromProvider,
	"cloudflare_api_shield":                                      config.IdentifierFromProvider,
	"cloudflare_api_shield_operation":                            config.IdentifierFromProvider,
	"cloudflare_api_shield_operation_schema_validation_settings": config.IdentifierFromProvider,
	"cloudflare_api_shield_operation_schema":                     config.IdentifierFromProvider,
	"cloudflare_api_shield_schema_validation_settings.":          config.IdentifierFromProvider,
	"cloudflare_argo":                                            config.IdentifierFromProvider,
	"cloudflare_authenticated_origin_pulls_certificate":          config.IdentifierFromProvider,
	"cloudflare_authenticated_origin_pulls":                      config.IdentifierFromProvider,
	"cloudflare_bot_management":                                  config.IdentifierFromProvider,
	"cloudflare_byo_ip_prefix":                                   config.IdentifierFromProvider,
	"cloudflare_certificate_pack":                                config.IdentifierFromProvider,
	"cloudflare_custom_hostname_fallback_origin":                 config.IdentifierFromProvider,
	"cloudflare_custom_hostname":                                 config.IdentifierFromProvider,
	"cloudflare_custom_pages":                                    config.IdentifierFromProvider,
	"cloudflare_custom_ssl":                                      config.IdentifierFromProvider,
	"cloudflare_d1_database":                                     config.IdentifierFromProvider,
	"cloudflare_device_dex":                                      config.IdentifierFromProvider,
	"cloudflare_device_managed_networks":                         config.IdentifierFromProvider,
	"cloudflare_device_policy_certificates":                      config.IdentifierFromProvider,
	"cloudflare_device_posture_integration":                      config.IdentifierFromProvider,
	"cloudflare_device_posture_rule":                             config.IdentifierFromProvider,
	"cloudflare_device_settings_policy":                          config.IdentifierFromProvider,
	"cloudflare_dlp_profile":                                     config.IdentifierFromProvider,
	"cloudflare_email_routing_catch_all":                         config.IdentifierFromProvider,
	"cloudflare_email_routing_settings":                          config.IdentifierFromProvider,
	"cloudflare_fallback_domain":                                 config.IdentifierFromProvider,
	"cloudflare_filter":                                          config.IdentifierFromProvider,
	"cloudflare_firewall_rule":                                   config.IdentifierFromProvider,
	"cloudflare_gre_tunnel":                                      config.IdentifierFromProvider,
	"cloudflare_healthcheck":                                     config.IdentifierFromProvider,
	"cloudflare_hostname_tls_setting":                            config.IdentifierFromProvider,
	"cloudflare_hostname_tls_setting_ciphers":                    config.IdentifierFromProvider,
	//"cloudflare_hyperdrive_config":                               config.IdentifierFromProvider,
	"cloudflare_ipsec_tunnel":                    config.IdentifierFromProvider,
	"cloudflare_keyless_certificate":             config.IdentifierFromProvider,
	"cloudflare_list":                            config.IdentifierFromProvider,
	"cloudflare_load_balancer":                   config.IdentifierFromProvider,
	"cloudflare_load_balancer_migrate":           config.IdentifierFromProvider,
	"cloudflare_load_balancer_monitor":           config.IdentifierFromProvider,
	"cloudflare_load_balancer_pool":              config.IdentifierFromProvider,
	"cloudflare_logpull_retention":               config.IdentifierFromProvider,
	"cloudflare_logpush_job":                     config.IdentifierFromProvider,
	"cloudflare_logpush_ownership_challenge":     config.IdentifierFromProvider,
	"cloudflare_magic_firewall_ruleset":          config.IdentifierFromProvider,
	"cloudflare_managed_headers":                 config.IdentifierFromProvider,
	"cloudflare_mtls_certificate":                config.IdentifierFromProvider,
	"cloudflare_notification_policy":             config.IdentifierFromProvider,
	"cloudflare_notification_policy_webhooks":    config.IdentifierFromProvider,
	"cloudflare_observatory_scheduled_test":      config.IdentifierFromProvider,
	"cloudflare_origin_ca_certificate":           config.IdentifierFromProvider,
	"cloudflare_page_rule":                       config.IdentifierFromProvider,
	"cloudflare_pages_domain":                    config.IdentifierFromProvider,
	"cloudflare_pages_project":                   config.IdentifierFromProvider,
	"cloudflare_queue":                           config.IdentifierFromProvider,
	"cloudflare_r2_bucket":                       config.IdentifierFromProvider,
	"cloudflare_rate_limit":                      config.IdentifierFromProvider,
	"cloudflare_record":                          config.IdentifierFromProvider,
	"cloudflare_regional_hostname":               config.IdentifierFromProvider,
	"cloudflare_regional_tiered_cache":           config.IdentifierFromProvider,
	"cloudflare_ruleset":                         config.IdentifierFromProvider,
	"cloudflare_spectrum_application":            config.IdentifierFromProvider,
	"cloudflare_split_tunnel":                    config.IdentifierFromProvider,
	"cloudflare_static_route":                    config.IdentifierFromProvider,
	"cloudflare_teams_account":                   config.IdentifierFromProvider,
	"cloudflare_teams_list":                      config.IdentifierFromProvider,
	"cloudflare_teams_location":                  config.IdentifierFromProvider,
	"cloudflare_teams_proxy_endpoint":            config.IdentifierFromProvider,
	"cloudflare_teams_rule":                      config.IdentifierFromProvider,
	"cloudflare_tiered_cache":                    config.IdentifierFromProvider,
	"cloudflare_total_tls":                       config.IdentifierFromProvider,
	"cloudflare_tunnel":                          config.IdentifierFromProvider,
	"cloudflare_tunnel_config":                   config.IdentifierFromProvider,
	"cloudflare_tunnel_route":                    config.IdentifierFromProvider,
	"cloudflare_tunnel_virtual_network":          config.IdentifierFromProvider,
	"cloudflare_turnstile_widget":                config.IdentifierFromProvider,
	"cloudflare_url_normalization_settings":      config.IdentifierFromProvider,
	"cloudflare_user_agent_blocking_rule":        config.IdentifierFromProvider,
	"cloudflare_waiting_room":                    config.IdentifierFromProvider,
	"cloudflare_waiting_room_event":              config.IdentifierFromProvider,
	"cloudflare_waiting_room_rules":              config.IdentifierFromProvider,
	"cloudflare_waiting_room_settings":           config.IdentifierFromProvider,
	"cloudflare_web3_hostname":                   config.IdentifierFromProvider,
	"cloudflare_web_analytics_rule":              config.IdentifierFromProvider,
	"cloudflare_web_analytics_site":              config.IdentifierFromProvider,
	"cloudflare_worker_cron_trigger":             config.IdentifierFromProvider,
	"cloudflare_workers_domain":                  config.IdentifierFromProvider,
	"cloudflare_workers_kv":                      config.IdentifierFromProvider,
	"cloudflare_workers_kv_namespace":            config.IdentifierFromProvider,
	"cloudflare_worker_route":                    config.IdentifierFromProvider,
	"cloudflare_worker_script":                   config.IdentifierFromProvider,
	"cloudflare_worker_secret":                   config.IdentifierFromProvider,
	"cloudflare_workers_for_platforms_namespace": config.IdentifierFromProvider,
	"cloudflare_zone":                            config.IdentifierFromProvider,
	"cloudflare_zone_cache_reserve":              config.IdentifierFromProvider,
	"cloudflare_zone_cache_variants":             config.IdentifierFromProvider,
	"cloudflare_zone_dnssec":                     config.IdentifierFromProvider,
	"cloudflare_zone_hold":                       config.IdentifierFromProvider,
	"cloudflare_zone_lockdown":                   config.IdentifierFromProvider,
	"cloudflare_zone_settings_override":          config.IdentifierFromProvider,
}

// cliReconciledExternalNameConfigs contains all external name configurations
// belonging to Terraform resources to be reconciled under the CLI-based
// architecture for this provider.
var cliReconciledExternalNameConfigs = map[string]config.ExternalName{}

// TemplatedStringAsIdentifierWithNoName uses TemplatedStringAsIdentifier but
// without the name initializer. This allows it to be used in cases where the ID
// is constructed with parameters and a provider-defined value, meaning no
// user-defined input. Since the external name is not user-defined, the name
// initializer has to be disabled.
func TemplatedStringAsIdentifierWithNoName(tmpl string) config.ExternalName {
	e := config.TemplatedStringAsIdentifier("", tmpl)
	e.DisableNameInitializer = true
	return e
}

// resourceConfigurator applies all external name configs
// listed in the table terraformPluginSDKExternalNameConfigs and
// cliReconciledExternalNameConfigs and sets the version
// of those resources to v1beta1. For those resource in
// terraformPluginSDKExternalNameConfigs, it also sets
// config.Resource.UseNoForkClient to `true`.
func resourceConfigurator() config.ResourceOption {
	return func(r *config.Resource) {
		// if configured both for the no-fork and CLI based architectures,
		// no-fork configuration prevails
		e, configured := terraformPluginSDKExternalNameConfigs[r.Name]
		if !configured {
			e, configured = cliReconciledExternalNameConfigs[r.Name]
		}
		if !configured {
			return
		}
		r.Version = common.VersionV1Beta1
		r.ExternalName = e
	}
}
