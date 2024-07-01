// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"context"
	// Note(ezgidemirel): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/pkg/config"
	"github.com/crossplane/upjet/pkg/config/conversion"
	"github.com/crossplane/upjet/pkg/registry/reference"
	"github.com/crossplane/upjet/pkg/schema/traverser"
	conversiontfjson "github.com/crossplane/upjet/pkg/types/conversion/tfjson"
	tfjson "github.com/hashicorp/terraform-json"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/ecoupal-believe/provider-cloudflare/config/access"
	"github.com/ecoupal-believe/provider-cloudflare/config/account"
	"github.com/ecoupal-believe/provider-cloudflare/config/address"
	"github.com/ecoupal-believe/provider-cloudflare/config/apishield"
	"github.com/ecoupal-believe/provider-cloudflare/config/argo"
	"github.com/ecoupal-believe/provider-cloudflare/config/authenticatedoriginpulls"
	"github.com/ecoupal-believe/provider-cloudflare/config/bot"
	"github.com/ecoupal-believe/provider-cloudflare/config/byoip"
	"github.com/ecoupal-believe/provider-cloudflare/config/certificate"
	"github.com/ecoupal-believe/provider-cloudflare/config/custom"
	"github.com/ecoupal-believe/provider-cloudflare/config/customhostname"
	"github.com/ecoupal-believe/provider-cloudflare/config/d1"
	"github.com/ecoupal-believe/provider-cloudflare/config/dlp"
	"github.com/ecoupal-believe/provider-cloudflare/config/dns"
	"github.com/ecoupal-believe/provider-cloudflare/config/emailrouting"
	"github.com/ecoupal-believe/provider-cloudflare/config/filters"
	"github.com/ecoupal-believe/provider-cloudflare/config/firewall"
	"github.com/ecoupal-believe/provider-cloudflare/config/lists"
	"github.com/ecoupal-believe/provider-cloudflare/config/loadbalancer"
	"github.com/ecoupal-believe/provider-cloudflare/config/logpush"
	"github.com/ecoupal-believe/provider-cloudflare/config/magic"
	"github.com/ecoupal-believe/provider-cloudflare/config/notification"
	"github.com/ecoupal-believe/provider-cloudflare/config/observatory"
	"github.com/ecoupal-believe/provider-cloudflare/config/originca"
	"github.com/ecoupal-believe/provider-cloudflare/config/page"
	"github.com/ecoupal-believe/provider-cloudflare/config/pages"
	"github.com/ecoupal-believe/provider-cloudflare/config/r2"
	"github.com/ecoupal-believe/provider-cloudflare/config/ruleset"
	"github.com/ecoupal-believe/provider-cloudflare/config/spectrum"
	"github.com/ecoupal-believe/provider-cloudflare/config/teams"
	"github.com/ecoupal-believe/provider-cloudflare/config/tunnel"
	"github.com/ecoupal-believe/provider-cloudflare/config/turnstile"
	"github.com/ecoupal-believe/provider-cloudflare/config/waitingroom"
	"github.com/ecoupal-believe/provider-cloudflare/config/warp"
	"github.com/ecoupal-believe/provider-cloudflare/config/web3"
	"github.com/ecoupal-believe/provider-cloudflare/config/webanalytics"
	"github.com/ecoupal-believe/provider-cloudflare/config/worker"
	"github.com/ecoupal-believe/provider-cloudflare/config/zone"
)

const (
	resourcePrefix = "cloudflare"
	modulePath     = "github.com/believe/provider-cloudflare"
)

var (
	//go:embed schema.json
	providerSchema string

	//go:embed provider-metadata.yaml
	providerMetadata []byte
)

var skipList = []string{}

// workaround for the TF Google v4.77.0-based no-fork release: We would like to
// keep the types in the generated CRDs intact
// (prevent number->int type replacements).
func getProviderSchema(s string) (*schema.Provider, error) {
	ps := tfjson.ProviderSchemas{}
	if err := ps.UnmarshalJSON([]byte(s)); err != nil {
		panic(err)
	}
	if len(ps.Schemas) != 1 {
		return nil, errors.Errorf("there should exactly be 1 provider schema but there are %d", len(ps.Schemas))
	}
	var rs map[string]*tfjson.Schema
	for _, v := range ps.Schemas {
		rs = v.ResourceSchemas
		break
	}
	return &schema.Provider{
		ResourcesMap: conversiontfjson.GetV2ResourceMap(rs),
	}, nil
}

// GetProvider returns provider configuration
func GetProvider(_ context.Context, generationProvider bool) (*ujconfig.Provider, error) {
	sdkProvider := provider.Provider()

	if generationProvider {
		p, err := getProviderSchema(providerSchema)
		if err != nil {
			return nil, errors.Wrap(err, "cannot read the Terraform SDK provider from the JSON schema for code generation")
		}
		if err := traverser.TFResourceSchema(sdkProvider.ResourcesMap).Traverse(traverser.NewMaxItemsSync(p.ResourcesMap)); err != nil {
			return nil, errors.Wrap(err, "cannot sync the MaxItems constraints between the Go schema and the JSON schema")
		}
		// use the JSON schema to temporarily prevent float64->int64
		// conversions in the CRD APIs.
		// We would like to convert to int64s with the next major release of
		// the provider.
		sdkProvider = p
	}

	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, providerMetadata,
		ujconfig.WithDefaultResourceOptions(
			groupOverrides(),
			externalNameConfig(),
			defaultVersion(),
			resourceConfigurator(),
			descriptionOverrides(),
		),
		ujconfig.WithRootGroup("gcp.upbound.io"),
		ujconfig.WithShortName("gcp"),
		// Comment out the following line to generate all resources.
		ujconfig.WithIncludeList(resourceList(cliReconciledExternalNameConfigs)),
		ujconfig.WithTerraformPluginSDKIncludeList(resourceList(terraformPluginSDKExternalNameConfigs)),
		ujconfig.WithReferenceInjectors([]ujconfig.ReferenceInjector{reference.NewInjector(modulePath)}),
		ujconfig.WithSkipList(skipList),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithMainTemplate(hack.MainTemplate),
		ujconfig.WithTerraformProvider(sdkProvider),
		ujconfig.WithSchemaTraversers(&ujconfig.SingletonListEmbedder{}),
	)

	bumpVersionsWithEmbeddedLists(pc)
	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		access.Configure,
		account.Configure,
		address.Configure,
		apishield.Configure,
		argo.Configure,
		authenticatedoriginpulls.Configure,
		bot.Configure,
		byoip.Configure,
		certificate.Configure,
		custom.Configure,
		customhostname.Configure,
		d1.Configure,
		dlp.Configure,
		dns.Configure,
		emailrouting.Configure,
		filters.Configure,
		firewall.Configure,
		//hyperdrive.Configure,
		lists.Configure,
		loadbalancer.Configure,
		logpush.Configure,
		magic.Configure,
		notification.Configure,
		observatory.Configure,
		originca.Configure,
		page.Configure,
		pages.Configure,
		r2.Configure,
		ruleset.Configure,
		spectrum.Configure,
		teams.Configure,
		tunnel.Configure,
		turnstile.Configure,
		waitingroom.Configure,
		warp.Configure,
		web3.Configure,
		webanalytics.Configure,
		worker.Configure,
		zone.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc, nil
}

// resourceList returns the list of resources that have external
// name configured in the specified table.
func resourceList(t map[string]ujconfig.ExternalName) []string {
	l := make([]string, len(t))
	i := 0
	for n := range t {
		// Expected format is regex and we'd like to have exact matches.
		l[i] = n + "$"
		i++
	}
	return l
}

func bumpVersionsWithEmbeddedLists(pc *ujconfig.Provider) {
	for n, r := range pc.Resources {
		r := r
		// nothing to do if no singleton list has been converted to
		// an embedded object
		if len(r.CRDListConversionPaths()) == 0 {
			continue
		}
		r.Version = "v1beta2"
		r.PreviousVersions = []string{VersionV1Beta1}
		// we would like to set the storage version to v1beta1 to facilitate
		// downgrades.
		r.SetCRDStorageVersion("v1beta1")
		r.ControllerReconcileVersion = "v1beta1"
		r.Conversions = []conversion.Conversion{
			conversion.NewIdentityConversionExpandPaths(conversion.AllVersions, conversion.AllVersions, conversion.DefaultPathPrefixes(), r.CRDListConversionPaths()...),
			conversion.NewSingletonListConversion("v1beta1", "v1beta2", conversion.DefaultPathPrefixes(), r.CRDListConversionPaths(), conversion.ToEmbeddedObject),
			conversion.NewSingletonListConversion("v1beta2", "v1beta1", conversion.DefaultPathPrefixes(), r.CRDListConversionPaths(), conversion.ToSingletonList)}
		pc.Resources[n] = r
	}
}

func init() {
}
