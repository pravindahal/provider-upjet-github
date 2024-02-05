// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type BranchProtectionInitParameters struct {

	// Boolean, setting this to true to allow the branch to be deleted.
	AllowsDeletions *bool `json:"allowsDeletions,omitempty" tf:"allows_deletions,omitempty"`

	// Boolean, setting this to true to allow force pushes on the branch.
	AllowsForcePushes *bool `json:"allowsForcePushes,omitempty" tf:"allows_force_pushes,omitempty"`

	// Boolean, setting this to true to block creating the branch.
	BlocksCreations *bool `json:"blocksCreations,omitempty" tf:"blocks_creations,omitempty"`

	// Boolean, setting this to true enforces status checks for repository administrators.
	EnforceAdmins *bool `json:"enforceAdmins,omitempty" tf:"enforce_admins,omitempty"`

	// Boolean, Setting this to true will make the branch read-only and preventing any pushes to it. Defaults to false
	LockBranch *bool `json:"lockBranch,omitempty" tf:"lock_branch,omitempty"`

	// Identifies the protection rule pattern.
	Pattern *string `json:"pattern,omitempty" tf:"pattern,omitempty"`

	// The list of actor Names/IDs that may push to the branch. Actor names must either begin with a "/" for users or the organization name followed by a "/" for teams.
	PushRestrictions []*string `json:"pushRestrictions,omitempty" tf:"push_restrictions,omitempty"`

	// Boolean, setting this to true requires all conversations on code must be resolved before a pull request can be merged.
	RequireConversationResolution *bool `json:"requireConversationResolution,omitempty" tf:"require_conversation_resolution,omitempty"`

	// Boolean, setting this to true requires all commits to be signed with GPG.
	RequireSignedCommits *bool `json:"requireSignedCommits,omitempty" tf:"require_signed_commits,omitempty"`

	// Boolean, setting this to true enforces a linear commit Git history, which prevents anyone from pushing merge commits to a branch
	RequiredLinearHistory *bool `json:"requiredLinearHistory,omitempty" tf:"required_linear_history,omitempty"`

	// Enforce restrictions for pull request reviews. See Required Pull Request Reviews below for details.
	RequiredPullRequestReviews []RequiredPullRequestReviewsInitParameters `json:"requiredPullRequestReviews,omitempty" tf:"required_pull_request_reviews,omitempty"`

	// Enforce restrictions for required status checks. See Required Status Checks below for details.
	RequiredStatusChecks []RequiredStatusChecksInitParameters `json:"requiredStatusChecks,omitempty" tf:"required_status_checks,omitempty"`
}

type BranchProtectionObservation struct {

	// Boolean, setting this to true to allow the branch to be deleted.
	AllowsDeletions *bool `json:"allowsDeletions,omitempty" tf:"allows_deletions,omitempty"`

	// Boolean, setting this to true to allow force pushes on the branch.
	AllowsForcePushes *bool `json:"allowsForcePushes,omitempty" tf:"allows_force_pushes,omitempty"`

	// Boolean, setting this to true to block creating the branch.
	BlocksCreations *bool `json:"blocksCreations,omitempty" tf:"blocks_creations,omitempty"`

	// Boolean, setting this to true enforces status checks for repository administrators.
	EnforceAdmins *bool `json:"enforceAdmins,omitempty" tf:"enforce_admins,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Boolean, Setting this to true will make the branch read-only and preventing any pushes to it. Defaults to false
	LockBranch *bool `json:"lockBranch,omitempty" tf:"lock_branch,omitempty"`

	// Identifies the protection rule pattern.
	Pattern *string `json:"pattern,omitempty" tf:"pattern,omitempty"`

	// The list of actor Names/IDs that may push to the branch. Actor names must either begin with a "/" for users or the organization name followed by a "/" for teams.
	PushRestrictions []*string `json:"pushRestrictions,omitempty" tf:"push_restrictions,omitempty"`

	// The name or node ID of the repository associated with this branch protection rule.
	// Node ID or name of repository
	RepositoryID *string `json:"repositoryId,omitempty" tf:"repository_id,omitempty"`

	// Boolean, setting this to true requires all conversations on code must be resolved before a pull request can be merged.
	RequireConversationResolution *bool `json:"requireConversationResolution,omitempty" tf:"require_conversation_resolution,omitempty"`

	// Boolean, setting this to true requires all commits to be signed with GPG.
	RequireSignedCommits *bool `json:"requireSignedCommits,omitempty" tf:"require_signed_commits,omitempty"`

	// Boolean, setting this to true enforces a linear commit Git history, which prevents anyone from pushing merge commits to a branch
	RequiredLinearHistory *bool `json:"requiredLinearHistory,omitempty" tf:"required_linear_history,omitempty"`

	// Enforce restrictions for pull request reviews. See Required Pull Request Reviews below for details.
	RequiredPullRequestReviews []RequiredPullRequestReviewsObservation `json:"requiredPullRequestReviews,omitempty" tf:"required_pull_request_reviews,omitempty"`

	// Enforce restrictions for required status checks. See Required Status Checks below for details.
	RequiredStatusChecks []RequiredStatusChecksObservation `json:"requiredStatusChecks,omitempty" tf:"required_status_checks,omitempty"`
}

type BranchProtectionParameters struct {

	// Boolean, setting this to true to allow the branch to be deleted.
	// +kubebuilder:validation:Optional
	AllowsDeletions *bool `json:"allowsDeletions,omitempty" tf:"allows_deletions,omitempty"`

	// Boolean, setting this to true to allow force pushes on the branch.
	// +kubebuilder:validation:Optional
	AllowsForcePushes *bool `json:"allowsForcePushes,omitempty" tf:"allows_force_pushes,omitempty"`

	// Boolean, setting this to true to block creating the branch.
	// +kubebuilder:validation:Optional
	BlocksCreations *bool `json:"blocksCreations,omitempty" tf:"blocks_creations,omitempty"`

	// Boolean, setting this to true enforces status checks for repository administrators.
	// +kubebuilder:validation:Optional
	EnforceAdmins *bool `json:"enforceAdmins,omitempty" tf:"enforce_admins,omitempty"`

	// Boolean, Setting this to true will make the branch read-only and preventing any pushes to it. Defaults to false
	// +kubebuilder:validation:Optional
	LockBranch *bool `json:"lockBranch,omitempty" tf:"lock_branch,omitempty"`

	// Identifies the protection rule pattern.
	// +kubebuilder:validation:Optional
	Pattern *string `json:"pattern,omitempty" tf:"pattern,omitempty"`

	// The list of actor Names/IDs that may push to the branch. Actor names must either begin with a "/" for users or the organization name followed by a "/" for teams.
	// +kubebuilder:validation:Optional
	PushRestrictions []*string `json:"pushRestrictions,omitempty" tf:"push_restrictions,omitempty"`

	// The name or node ID of the repository associated with this branch protection rule.
	// Node ID or name of repository
	// +crossplane:generate:reference:type=github.com/coopnorge/provider-github/apis/repo/v1alpha1.Repository
	// +kubebuilder:validation:Optional
	RepositoryID *string `json:"repositoryId,omitempty" tf:"repository_id,omitempty"`

	// Reference to a Repository in repo to populate repositoryId.
	// +kubebuilder:validation:Optional
	RepositoryIDRef *v1.Reference `json:"repositoryIdRef,omitempty" tf:"-"`

	// Selector for a Repository in repo to populate repositoryId.
	// +kubebuilder:validation:Optional
	RepositoryIDSelector *v1.Selector `json:"repositoryIdSelector,omitempty" tf:"-"`

	// Boolean, setting this to true requires all conversations on code must be resolved before a pull request can be merged.
	// +kubebuilder:validation:Optional
	RequireConversationResolution *bool `json:"requireConversationResolution,omitempty" tf:"require_conversation_resolution,omitempty"`

	// Boolean, setting this to true requires all commits to be signed with GPG.
	// +kubebuilder:validation:Optional
	RequireSignedCommits *bool `json:"requireSignedCommits,omitempty" tf:"require_signed_commits,omitempty"`

	// Boolean, setting this to true enforces a linear commit Git history, which prevents anyone from pushing merge commits to a branch
	// +kubebuilder:validation:Optional
	RequiredLinearHistory *bool `json:"requiredLinearHistory,omitempty" tf:"required_linear_history,omitempty"`

	// Enforce restrictions for pull request reviews. See Required Pull Request Reviews below for details.
	// +kubebuilder:validation:Optional
	RequiredPullRequestReviews []RequiredPullRequestReviewsParameters `json:"requiredPullRequestReviews,omitempty" tf:"required_pull_request_reviews,omitempty"`

	// Enforce restrictions for required status checks. See Required Status Checks below for details.
	// +kubebuilder:validation:Optional
	RequiredStatusChecks []RequiredStatusChecksParameters `json:"requiredStatusChecks,omitempty" tf:"required_status_checks,omitempty"`
}

type RequiredPullRequestReviewsInitParameters struct {

	// :  Dismiss approved reviews automatically when a new commit is pushed. Defaults to false.
	DismissStaleReviews *bool `json:"dismissStaleReviews,omitempty" tf:"dismiss_stale_reviews,omitempty"`

	// :  The list of actor Names/IDs with dismissal access. If not empty, restrict_dismissals is ignored. Actor names must either begin with a "/" for users or the organization name followed by a "/" for teams.
	DismissalRestrictions []*string `json:"dismissalRestrictions,omitempty" tf:"dismissal_restrictions,omitempty"`

	// :  The list of actor Names/IDs that are allowed to bypass pull request requirements. Actor names must either begin with a "/" for users or the organization name followed by a "/" for teams.
	PullRequestBypassers []*string `json:"pullRequestBypassers,omitempty" tf:"pull_request_bypassers,omitempty"`

	// :  Require an approved review in pull requests including files with a designated code owner. Defaults to false.
	RequireCodeOwnerReviews *bool `json:"requireCodeOwnerReviews,omitempty" tf:"require_code_owner_reviews,omitempty"`

	// :  Require that The most recent push must be approved by someone other than the last pusher.  Defaults to false
	RequireLastPushApproval *bool `json:"requireLastPushApproval,omitempty" tf:"require_last_push_approval,omitempty"`

	// 6. This requirement matches GitHub's API, see the upstream documentation for more information.
	// (https://developer.github.com/v3/repos/branches/#parameters-1) for more information.
	RequiredApprovingReviewCount *float64 `json:"requiredApprovingReviewCount,omitempty" tf:"required_approving_review_count,omitempty"`

	// :  Restrict pull request review dismissals.
	RestrictDismissals *bool `json:"restrictDismissals,omitempty" tf:"restrict_dismissals,omitempty"`
}

type RequiredPullRequestReviewsObservation struct {

	// :  Dismiss approved reviews automatically when a new commit is pushed. Defaults to false.
	DismissStaleReviews *bool `json:"dismissStaleReviews,omitempty" tf:"dismiss_stale_reviews,omitempty"`

	// :  The list of actor Names/IDs with dismissal access. If not empty, restrict_dismissals is ignored. Actor names must either begin with a "/" for users or the organization name followed by a "/" for teams.
	DismissalRestrictions []*string `json:"dismissalRestrictions,omitempty" tf:"dismissal_restrictions,omitempty"`

	// :  The list of actor Names/IDs that are allowed to bypass pull request requirements. Actor names must either begin with a "/" for users or the organization name followed by a "/" for teams.
	PullRequestBypassers []*string `json:"pullRequestBypassers,omitempty" tf:"pull_request_bypassers,omitempty"`

	// :  Require an approved review in pull requests including files with a designated code owner. Defaults to false.
	RequireCodeOwnerReviews *bool `json:"requireCodeOwnerReviews,omitempty" tf:"require_code_owner_reviews,omitempty"`

	// :  Require that The most recent push must be approved by someone other than the last pusher.  Defaults to false
	RequireLastPushApproval *bool `json:"requireLastPushApproval,omitempty" tf:"require_last_push_approval,omitempty"`

	// 6. This requirement matches GitHub's API, see the upstream documentation for more information.
	// (https://developer.github.com/v3/repos/branches/#parameters-1) for more information.
	RequiredApprovingReviewCount *float64 `json:"requiredApprovingReviewCount,omitempty" tf:"required_approving_review_count,omitempty"`

	// :  Restrict pull request review dismissals.
	RestrictDismissals *bool `json:"restrictDismissals,omitempty" tf:"restrict_dismissals,omitempty"`
}

type RequiredPullRequestReviewsParameters struct {

	// :  Dismiss approved reviews automatically when a new commit is pushed. Defaults to false.
	// +kubebuilder:validation:Optional
	DismissStaleReviews *bool `json:"dismissStaleReviews,omitempty" tf:"dismiss_stale_reviews,omitempty"`

	// :  The list of actor Names/IDs with dismissal access. If not empty, restrict_dismissals is ignored. Actor names must either begin with a "/" for users or the organization name followed by a "/" for teams.
	// +kubebuilder:validation:Optional
	DismissalRestrictions []*string `json:"dismissalRestrictions,omitempty" tf:"dismissal_restrictions,omitempty"`

	// :  The list of actor Names/IDs that are allowed to bypass pull request requirements. Actor names must either begin with a "/" for users or the organization name followed by a "/" for teams.
	// +kubebuilder:validation:Optional
	PullRequestBypassers []*string `json:"pullRequestBypassers,omitempty" tf:"pull_request_bypassers,omitempty"`

	// :  Require an approved review in pull requests including files with a designated code owner. Defaults to false.
	// +kubebuilder:validation:Optional
	RequireCodeOwnerReviews *bool `json:"requireCodeOwnerReviews,omitempty" tf:"require_code_owner_reviews,omitempty"`

	// :  Require that The most recent push must be approved by someone other than the last pusher.  Defaults to false
	// +kubebuilder:validation:Optional
	RequireLastPushApproval *bool `json:"requireLastPushApproval,omitempty" tf:"require_last_push_approval,omitempty"`

	// 6. This requirement matches GitHub's API, see the upstream documentation for more information.
	// (https://developer.github.com/v3/repos/branches/#parameters-1) for more information.
	// +kubebuilder:validation:Optional
	RequiredApprovingReviewCount *float64 `json:"requiredApprovingReviewCount,omitempty" tf:"required_approving_review_count,omitempty"`

	// :  Restrict pull request review dismissals.
	// +kubebuilder:validation:Optional
	RestrictDismissals *bool `json:"restrictDismissals,omitempty" tf:"restrict_dismissals,omitempty"`
}

type RequiredStatusChecksInitParameters struct {

	// :  The list of status checks to require in order to merge into this branch. No status checks are required by default.
	Contexts []*string `json:"contexts,omitempty" tf:"contexts,omitempty"`

	// :  Require branches to be up to date before merging. Defaults to false.
	Strict *bool `json:"strict,omitempty" tf:"strict,omitempty"`
}

type RequiredStatusChecksObservation struct {

	// :  The list of status checks to require in order to merge into this branch. No status checks are required by default.
	Contexts []*string `json:"contexts,omitempty" tf:"contexts,omitempty"`

	// :  Require branches to be up to date before merging. Defaults to false.
	Strict *bool `json:"strict,omitempty" tf:"strict,omitempty"`
}

type RequiredStatusChecksParameters struct {

	// :  The list of status checks to require in order to merge into this branch. No status checks are required by default.
	// +kubebuilder:validation:Optional
	Contexts []*string `json:"contexts,omitempty" tf:"contexts,omitempty"`

	// :  Require branches to be up to date before merging. Defaults to false.
	// +kubebuilder:validation:Optional
	Strict *bool `json:"strict,omitempty" tf:"strict,omitempty"`
}

// BranchProtectionSpec defines the desired state of BranchProtection
type BranchProtectionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BranchProtectionParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider BranchProtectionInitParameters `json:"initProvider,omitempty"`
}

// BranchProtectionStatus defines the observed state of BranchProtection.
type BranchProtectionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BranchProtectionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BranchProtection is the Schema for the BranchProtections API. Protects a GitHub branch.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,github}
type BranchProtection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.pattern) || (has(self.initProvider) && has(self.initProvider.pattern))",message="spec.forProvider.pattern is a required parameter"
	Spec   BranchProtectionSpec   `json:"spec"`
	Status BranchProtectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BranchProtectionList contains a list of BranchProtections
type BranchProtectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BranchProtection `json:"items"`
}

// Repository type metadata.
var (
	BranchProtection_Kind             = "BranchProtection"
	BranchProtection_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BranchProtection_Kind}.String()
	BranchProtection_KindAPIVersion   = BranchProtection_Kind + "." + CRDGroupVersion.String()
	BranchProtection_GroupVersionKind = CRDGroupVersion.WithKind(BranchProtection_Kind)
)

func init() {
	SchemeBuilder.Register(&BranchProtection{}, &BranchProtectionList{})
}
