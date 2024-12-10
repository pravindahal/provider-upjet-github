/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	v1alpha1 "github.com/crossplane-contrib/provider-upjet-github/apis/repo/v1alpha1"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this ActionsSecret.
func (mg *ActionsSecret) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Repository),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RepositoryRef,
		Selector:     mg.Spec.ForProvider.RepositorySelector,
		To: reference.To{
			List:    &v1alpha1.RepositoryList{},
			Managed: &v1alpha1.Repository{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Repository")
	}
	mg.Spec.ForProvider.Repository = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RepositoryRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Repository),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.RepositoryRef,
		Selector:     mg.Spec.InitProvider.RepositorySelector,
		To: reference.To{
			List:    &v1alpha1.RepositoryList{},
			Managed: &v1alpha1.Repository{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Repository")
	}
	mg.Spec.InitProvider.Repository = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RepositoryRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ActionsVariable.
func (mg *ActionsVariable) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Repository),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RepositoryRef,
		Selector:     mg.Spec.ForProvider.RepositorySelector,
		To: reference.To{
			List:    &v1alpha1.RepositoryList{},
			Managed: &v1alpha1.Repository{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Repository")
	}
	mg.Spec.ForProvider.Repository = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RepositoryRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Repository),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.RepositoryRef,
		Selector:     mg.Spec.InitProvider.RepositorySelector,
		To: reference.To{
			List:    &v1alpha1.RepositoryList{},
			Managed: &v1alpha1.Repository{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Repository")
	}
	mg.Spec.InitProvider.Repository = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RepositoryRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this RunnerGroup.
func (mg *RunnerGroup) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Name),
		Extract:      resource.ExtractParamPath("name", false),
		Reference:    mg.Spec.ForProvider.NameRef,
		Selector:     mg.Spec.ForProvider.NameSelector,
		To: reference.To{
			List:    &v1alpha1.RepositoryList{},
			Managed: &v1alpha1.Repository{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Name")
	}
	mg.Spec.ForProvider.Name = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Name),
		Extract:      resource.ExtractParamPath("name", false),
		Reference:    mg.Spec.InitProvider.NameRef,
		Selector:     mg.Spec.InitProvider.NameSelector,
		To: reference.To{
			List:    &v1alpha1.RepositoryList{},
			Managed: &v1alpha1.Repository{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Name")
	}
	mg.Spec.InitProvider.Name = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.NameRef = rsp.ResolvedReference

	return nil
}
