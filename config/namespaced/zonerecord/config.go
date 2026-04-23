package zonerecord

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("dnsimple_zone_record", func(r *config.Resource) {
		r.ShortGroup = "zonerecord"
		r.Version = "v1beta1"
	})
}
