package pullrequest

import "github.com/crossplane/upjet/pkg/config"

// Configure github_repository_pull_request resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("github_repository_pull_request", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "github"
		r.Kind = "PullRequest"
		r.ShortGroup = "repo"
		// This resource need the repository in which branch would be created
		// as an input. And by defining it as a reference to Repository
		// object, we can build cross resource referencing. See
		// repositoryRef in the example in the Testing section below.
		r.References["base_repository"] = config.Reference{
			Type: "github.com/coopnorge/provider-github/apis/repo/v1alpha1.Repository",
		}
		//    r.References["base_ref"] = config.Reference{
		//			Type: "github.com/coopnorge/provider-github/apis/repo/v1alpha1.Branch",
		//		}
		r.References["head_ref"] = config.Reference{
			Type: "github.com/coopnorge/provider-github/apis/repo/v1alpha1.Branch",
		}

	})
}
