package lke

import "github.com/crossplane-contrib/terrajet/pkg/config"

// Customize configures individual resources by adding custom ResourceConfigurators.
func Customize(p *config.Provider) {
	p.AddResourceConfigurator("linode_lke_cluster", func(r *config.Resource) {

		// we need to override the default group that terrajet generated for
		// this resource, which would be "github"
		r.ShortGroup = ""

		// this property is generated by the Linode provider and must not be set in this provider
		r.ExternalName = config.IdentifierFromProvider
	})
}
