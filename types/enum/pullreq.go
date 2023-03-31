// Copyright 2022 Harness Inc. All rights reserved.
// Use of this source code is governed by the Polyform Free Trial License
// that can be found in the LICENSE.md file for this repository.

package enum

// PullReqState defines pull request state.
type PullReqState string

func (PullReqState) Enum() []interface{}                  { return toInterfaceSlice(pullReqStates) }
func (s PullReqState) Sanitize() (PullReqState, bool)     { return Sanitize(s, GetAllPullReqStates) }
func GetAllPullReqStates() ([]PullReqState, PullReqState) { return pullReqStates, "" }

// PullReqState enumeration.
const (
	PullReqStateOpen   PullReqState = "open"
	PullReqStateMerged PullReqState = "merged"
	PullReqStateClosed PullReqState = "closed"
)

var pullReqStates = sortEnum([]PullReqState{
	PullReqStateOpen,
	PullReqStateMerged,
	PullReqStateClosed,
})

// PullReqSort defines pull request attribute that can be used for sorting.
type PullReqSort string

func (PullReqSort) Enum() []interface{}                { return toInterfaceSlice(pullReqSorts) }
func (s PullReqSort) Sanitize() (PullReqSort, bool)    { return Sanitize(s, GetAllPullReqSorts) }
func GetAllPullReqSorts() ([]PullReqSort, PullReqSort) { return pullReqSorts, PullReqSortNumber }

// PullReqSort enumeration.
const (
	PullReqSortNumber  = "number"
	PullReqSortCreated = "created"
	PullReqSortEdited  = "edited"
	PullReqSortMerged  = "merged"
)

var pullReqSorts = sortEnum([]PullReqSort{
	PullReqSortNumber,
	PullReqSortCreated,
	PullReqSortEdited,
	PullReqSortMerged,
})

// PullReqActivityType defines pull request activity message type.
// Essentially, the Type determines the structure of the pull request activity's Payload structure.
type PullReqActivityType string

func (PullReqActivityType) Enum() []interface{} { return toInterfaceSlice(pullReqActivityTypes) }

func (t PullReqActivityType) Sanitize() (PullReqActivityType, bool) {
	return Sanitize(t, GetAllPullReqActivityTypes)
}

func GetAllPullReqActivityTypes() ([]PullReqActivityType, PullReqActivityType) {
	return pullReqActivityTypes, "" // No default value
}

// PullReqActivityType enumeration.
const (
	PullReqActivityTypeComment      PullReqActivityType = "comment"
	PullReqActivityTypeCodeComment  PullReqActivityType = "code-comment"
	PullReqActivityTypeTitleChange  PullReqActivityType = "title-change"
	PullReqActivityTypeStateChange  PullReqActivityType = "state-change"
	PullReqActivityTypeReviewSubmit PullReqActivityType = "review-submit"
	PullReqActivityTypeBranchUpdate PullReqActivityType = "branch-update"
	PullReqActivityTypeBranchDelete PullReqActivityType = "branch-delete"
	PullReqActivityTypeMerge        PullReqActivityType = "merge"
)

var pullReqActivityTypes = sortEnum([]PullReqActivityType{
	PullReqActivityTypeComment,
	PullReqActivityTypeCodeComment,
	PullReqActivityTypeTitleChange,
	PullReqActivityTypeStateChange,
	PullReqActivityTypeReviewSubmit,
	PullReqActivityTypeBranchUpdate,
	PullReqActivityTypeBranchDelete,
	PullReqActivityTypeMerge,
})

// PullReqActivityKind defines kind of pull request activity system message.
// Kind defines the source of the pull request activity entry:
// Whether it's generated by the system, it's a user comment or a part of code review.
type PullReqActivityKind string

func (PullReqActivityKind) Enum() []interface{} { return toInterfaceSlice(pullReqActivityKinds) }

func (k PullReqActivityKind) Sanitize() (PullReqActivityKind, bool) {
	return Sanitize(k, GetAllPullReqActivityKinds)
}

func GetAllPullReqActivityKinds() ([]PullReqActivityKind, PullReqActivityKind) {
	return pullReqActivityKinds, "" // No default value
}

// PullReqActivityKind enumeration.
const (
	PullReqActivityKindSystem        PullReqActivityKind = "system"
	PullReqActivityKindComment       PullReqActivityKind = "comment"
	PullReqActivityKindChangeComment PullReqActivityKind = "change-comment"
)

var pullReqActivityKinds = sortEnum([]PullReqActivityKind{
	PullReqActivityKindSystem,
	PullReqActivityKindComment,
	PullReqActivityKindChangeComment,
})

// PullReqReviewDecision defines state of a pull request review.
type PullReqReviewDecision string

func (PullReqReviewDecision) Enum() []interface{} {
	return toInterfaceSlice(pullReqReviewDecisions)
}

func (decision PullReqReviewDecision) Sanitize() (PullReqReviewDecision, bool) {
	return Sanitize(decision, GetAllPullReqReviewDecisions)
}

func GetAllPullReqReviewDecisions() ([]PullReqReviewDecision, PullReqReviewDecision) {
	return pullReqReviewDecisions, "" // No default value
}

// PullReqReviewDecision enumeration.
const (
	PullReqReviewDecisionPending   PullReqReviewDecision = "pending"
	PullReqReviewDecisionReviewed  PullReqReviewDecision = "reviewed"
	PullReqReviewDecisionApproved  PullReqReviewDecision = "approved"
	PullReqReviewDecisionChangeReq PullReqReviewDecision = "changereq"
)

var pullReqReviewDecisions = sortEnum([]PullReqReviewDecision{
	PullReqReviewDecisionPending,
	PullReqReviewDecisionReviewed,
	PullReqReviewDecisionApproved,
	PullReqReviewDecisionChangeReq,
})

// PullReqReviewerType defines type of a pull request reviewer.
type PullReqReviewerType string

func (PullReqReviewerType) Enum() []interface{} { return toInterfaceSlice(pullReqReviewerTypes) }

func (reviewerType PullReqReviewerType) Sanitize() (PullReqReviewerType, bool) {
	return Sanitize(reviewerType, GetAllPullReqReviewerTypes)
}

func GetAllPullReqReviewerTypes() ([]PullReqReviewerType, PullReqReviewerType) {
	return pullReqReviewerTypes, "" // No default value
}

// PullReqReviewerType enumeration.
const (
	PullReqReviewerTypeRequested    PullReqReviewerType = "requested"
	PullReqReviewerTypeAssigned     PullReqReviewerType = "assigned"
	PullReqReviewerTypeSelfAssigned PullReqReviewerType = "self_assigned"
)

var pullReqReviewerTypes = sortEnum([]PullReqReviewerType{
	PullReqReviewerTypeRequested,
	PullReqReviewerTypeAssigned,
	PullReqReviewerTypeSelfAssigned,
})

// MergeMethod represents the approach to merge commits into base branch.
type MergeMethod string

func (MergeMethod) Enum() []interface{}                { return toInterfaceSlice(mergeMethods) }
func (m MergeMethod) Sanitize() (MergeMethod, bool)    { return Sanitize(m, GetAllMergeMethods) }
func GetAllMergeMethods() ([]MergeMethod, MergeMethod) { return mergeMethods, MergeMethodMerge }

const (
	// MergeMethodMerge create merge commit.
	MergeMethodMerge MergeMethod = "merge"
	// MergeMethodSquash squash commits into single commit before merging.
	MergeMethodSquash MergeMethod = "squash"
	// MergeMethodRebase rebase before merging.
	MergeMethodRebase MergeMethod = "rebase"
)

var mergeMethods = sortEnum([]MergeMethod{
	MergeMethodMerge,
	MergeMethodSquash,
	MergeMethodRebase,
})

type MergeCheckStatus string

const (
	// MergeCheckStatusUnchecked merge status has not been checked.
	MergeCheckStatusUnchecked MergeCheckStatus = "unchecked"
	// MergeCheckStatusConflict can’t merge into the target branch due to a potential conflict.
	MergeCheckStatusConflict MergeCheckStatus = "conflict"
	// MergeCheckStatusMergeable branch can merged cleanly into the target branch.
	MergeCheckStatusMergeable MergeCheckStatus = "mergeable"
)
