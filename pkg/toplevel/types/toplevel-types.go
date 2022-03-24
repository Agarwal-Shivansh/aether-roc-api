// Package types provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.7.0 DO NOT EDIT.
package types

import (
	"time"

	externalRef0 "github.com/onosproject/aether-roc-api/pkg/aether_2_0_0/types"
	externalRef1 "github.com/onosproject/aether-roc-api/pkg/aether_2_1_0/types"
	externalRef2 "github.com/onosproject/aether-roc-api/pkg/aether_4_0_0/types"
)

// Defines values for AbortPhaseState.
const (
	AbortPhaseStateABORTED AbortPhaseState = "ABORTED"

	AbortPhaseStateABORTING AbortPhaseState = "ABORTING"
)

// Defines values for ApplyPhaseState.
const (
	ApplyPhaseStateAPPLIED ApplyPhaseState = "APPLIED"

	ApplyPhaseStateAPPLYING ApplyPhaseState = "APPLYING"

	ApplyPhaseStateFAILED ApplyPhaseState = "FAILED"
)

// Defines values for CommitPhaseState.
const (
	CommitPhaseStateCOMMITTED CommitPhaseState = "COMMITTED"

	CommitPhaseStateCOMMITTING CommitPhaseState = "COMMITTING"
)

// Defines values for FailureType.
const (
	FailureTypeALREADYEXISTS FailureType = "ALREADY_EXISTS"

	FailureTypeCANCELED FailureType = "CANCELED"

	FailureTypeCONFLICT FailureType = "CONFLICT"

	FailureTypeFORBIDDEN FailureType = "FORBIDDEN"

	FailureTypeINTERNAL FailureType = "INTERNAL"

	FailureTypeINVALID FailureType = "INVALID"

	FailureTypeNOTFOUND FailureType = "NOT_FOUND"

	FailureTypeNOTSUPPORTED FailureType = "NOT_SUPPORTED"

	FailureTypeTIMEOUT FailureType = "TIMEOUT"

	FailureTypeUNAUTHORIZED FailureType = "UNAUTHORIZED"

	FailureTypeUNAVAILABLE FailureType = "UNAVAILABLE"

	FailureTypeUNKNOWN FailureType = "UNKNOWN"
)

// Defines values for InitializePhaseState.
const (
	InitializePhaseStateFAILED InitializePhaseState = "FAILED"

	InitializePhaseStateINITIALIZED InitializePhaseState = "INITIALIZED"

	InitializePhaseStateINITIALIZING InitializePhaseState = "INITIALIZING"
)

// Defines values for Isolation.
const (
	IsolationDEFAULT Isolation = "DEFAULT"

	IsolationSERIALIZABLE Isolation = "SERIALIZABLE"
)

// Defines values for State.
const (
	StateAPPLIED State = "APPLIED"

	StateFAILED State = "FAILED"

	StatePENDING State = "PENDING"

	StateVALIDATED State = "VALIDATED"
)

// Defines values for Synchronicity.
const (
	SynchronicityASYNCHRONOUS Synchronicity = "ASYNCHRONOUS"

	SynchronicitySYNCHRONOUS Synchronicity = "SYNCHRONOUS"
)

// Defines values for ValidatePhaseState.
const (
	ValidatePhaseStateFAILED ValidatePhaseState = "FAILED"

	ValidatePhaseStateVALIDATED ValidatePhaseState = "VALIDATED"

	ValidatePhaseStateVALIDATING ValidatePhaseState = "VALIDATING"
)

// Defines values for ValueType.
const (
	ValueTypeBOOL ValueType = "BOOL"

	ValueTypeBYTES ValueType = "BYTES"

	ValueTypeDECIMAL ValueType = "DECIMAL"

	ValueTypeEMPTY ValueType = "EMPTY"

	ValueTypeFLOAT ValueType = "FLOAT"

	ValueTypeINT ValueType = "INT"

	ValueTypeLEAFLISTBOOL ValueType = "LEAFLIST_BOOL"

	ValueTypeLEAFLISTBYTES ValueType = "LEAFLIST_BYTES"

	ValueTypeLEAFLISTDECIMAL ValueType = "LEAFLIST_DECIMAL"

	ValueTypeLEAFLISTFLOAT ValueType = "LEAFLIST_FLOAT"

	ValueTypeLEAFLISTINT ValueType = "LEAFLIST_INT"

	ValueTypeLEAFLISTSTRING ValueType = "LEAFLIST_STRING"

	ValueTypeLEAFLISTUINT ValueType = "LEAFLIST_UINT"

	ValueTypeSTRING ValueType = "STRING"

	ValueTypeUINT ValueType = "UINT"
)

// AbortPhaseState defines model for AbortPhaseState.
type AbortPhaseState string

// ApplyPhaseState defines model for ApplyPhaseState.
type ApplyPhaseState string

// ChangeTarget defines model for ChangeTarget.
type ChangeTarget struct {
	PathValues *PathValues `json:"path-values,omitempty"`
	TargetName *string     `json:"target-name,omitempty"`
}

// ChangeTransaction defines model for ChangeTransaction.
type ChangeTransaction []ChangeTarget

// CommitPhaseState defines model for CommitPhaseState.
type CommitPhaseState string

// Deleted defines model for Deleted.
type Deleted bool

// Details defines model for Details.
type Details struct {
	Change   *ChangeTransaction   `json:"change,omitempty"`
	Rollback *RollbackTransaction `json:"rollback,omitempty"`
}

// Elements defines model for Elements.
type Elements struct {

	// The top level container
	Application400 *externalRef2.Application `json:"application-4.0.0,omitempty"`

	// The top level container
	ConnectivityService400 *externalRef2.ConnectivityService `json:"connectivity-service-4.0.0,omitempty"`

	// The connectivity-services top level container
	ConnectivityServices200 *externalRef0.ConnectivityServices `json:"connectivity-services-2.0.0,omitempty"`

	// The connectivity-services top level container
	ConnectivityServices210 *externalRef1.ConnectivityServices `json:"connectivity-services-2.1.0,omitempty"`

	// The top level container
	DeviceGroup400 *externalRef2.DeviceGroup `json:"device-group-4.0.0,omitempty"`

	// The top level container
	Enterprise400 *externalRef2.Enterprise `json:"enterprise-4.0.0,omitempty"`

	// The top level enterprises container
	Enterprises200 *externalRef0.Enterprises `json:"enterprises-2.0.0,omitempty"`

	// The top level enterprises container
	Enterprises210 *externalRef1.Enterprises `json:"enterprises-2.1.0,omitempty"`

	// The top level container
	IpDomain400 *externalRef2.IpDomain `json:"ip-domain-4.0.0,omitempty"`

	// The top level container
	Site400 *externalRef2.Site `json:"site-4.0.0,omitempty"`

	// The top level container
	Template400 *externalRef2.Template `json:"template-4.0.0,omitempty"`

	// The top level container
	TrafficClass400 *externalRef2.TrafficClass `json:"traffic-class-4.0.0,omitempty"`

	// The top level container
	Upf400 *externalRef2.Upf `json:"upf-4.0.0,omitempty"`

	// The top level container
	Vcs400 *externalRef2.Vcs `json:"vcs-4.0.0,omitempty"`
}

// End defines model for End.
type End time.Time

// Failure defines model for Failure.
type Failure struct {
	Description *string `json:"description,omitempty"`

	// transaction failure type
	Type *FailureType `json:"type,omitempty"`
}

// transaction failure type
type FailureType string

// Index defines model for Index.
type Index int64

// InitializePhaseState defines model for InitializePhaseState.
type InitializePhaseState string

// Isolation defines model for Isolation.
type Isolation string

// PatchBody defines model for PatchBody.
type PatchBody struct {
	Deletes *Elements `json:"Deletes,omitempty"`

	// Model type and version of 'target' on first creation [link](https://docs.onosproject.org/onos-config/docs/gnmi_extensions/#use-of-extension-101-device-version-in-setrequest)
	Extensions *struct {
		ChangeName100   *string `json:"change-name-100,omitempty"`
		ModelType102    *string `json:"model-type-102,omitempty"`
		ModelVersion101 *string `json:"model-version-101,omitempty"`

		// Used in the responses, carries inforamtion about the transaction.
		TransactionInfo110 *struct {
			ID    *string `json:"ID,omitempty"`
			Index *int    `json:"index,omitempty"`
		} `json:"transaction-info-110,omitempty"`

		// Identifies whether a request needs to be handles Asynchronously (val: 0) or Synchronously (val: 1)
		TransactionStrategy111 *int `json:"transaction-strategy-111,omitempty"`
	} `json:"Extensions,omitempty"`
	Updates *Elements `json:"Updates,omitempty"`

	// Target (device name) to use by default if not specified on indivdual updates/deletes as an additional property
	DefaultTarget string `json:"default-target"`
}

// Path defines model for Path.
type Path string

// PathTarget defines model for PathTarget.
type PathTarget struct {
	Path *string `json:"path,omitempty"`

	// the state of a path/value in the configuration tree
	PathValue *PathValue `json:"path-value,omitempty"`
}

// the state of a path/value in the configuration tree
type PathValue struct {
	Deleted *Deleted `json:"deleted,omitempty"`
	Path    *Path    `json:"path,omitempty"`

	// value represented as a byte array
	Value *TypedValue `json:"value,omitempty"`
}

// PathValues defines model for PathValues.
type PathValues []PathTarget

// ProposalID defines model for ProposalID.
type ProposalID string

// Revision defines model for Revision.
type Revision int64

// RollbackTransaction defines model for RollbackTransaction.
type RollbackTransaction struct {
	RollbackIndex *Index `json:"rollback_index,omitempty"`
}

// Start defines model for Start.
type Start time.Time

// State defines model for State.
type State string

// Status defines model for Status.
type Status struct {
	Failure *Failure           `json:"failure,omitempty"`
	Phases  *TransactionPhases `json:"phases,omitempty"`

	// the set of proposals managed by the transaction
	Proposals *[]ProposalID `json:"proposals,omitempty"`
	State     *State        `json:"state,omitempty"`
}

// Strategy defines model for Strategy.
type Strategy struct {
	Isolation     *Isolation     `json:"isolation,omitempty"`
	Synchronicity *Synchronicity `json:"synchronicity,omitempty"`
}

// Synchronicity defines model for Synchronicity.
type Synchronicity string

// TargetName defines model for TargetName.
type TargetName struct {
	Name *string `json:"name,omitempty"`
}

// TargetsNames defines model for TargetsNames.
type TargetsNames []TargetName

// Transaction refers to a multi-target transactional change. Taken from https://github.com/onosproject/onos-api/tree/master/proto/onos/config/v2
type Transaction struct {
	Details *Details `json:"details,omitempty"`

	// the unique identifier of the transaction
	Id string `json:"id"`

	// a monotonically increasing, globally unique index of the change
	Index int64 `json:"index"`

	// the meta of the Transaction
	Meta struct {

		// the time at which the transaction was created
		Created *time.Time `json:"created,omitempty"`

		// the time at which the transcation was deleted
		Deleted *time.Time `json:"deleted,omitempty"`

		// the key of the Transaction
		Key      *string   `json:"key,omitempty"`
		Revision *Revision `json:"revision,omitempty"`

		// the time at which the transaction was last updated
		Updated *time.Time `json:"updated,omitempty"`

		// the version of the Transaction
		Version *int64 `json:"version,omitempty"`
	} `json:"meta"`
	Status   *Status   `json:"status,omitempty"`
	Strategy *Strategy `json:"strategy,omitempty"`

	// the name of the user that made the transaction
	Username *string `json:"username,omitempty"`
}

// TransactionAbortPhase defines model for TransactionAbortPhase.
type TransactionAbortPhase struct {
	State  *AbortPhaseState        `json:"state,omitempty"`
	Status *TransactionPhaseStatus `json:"status,omitempty"`
}

// TransactionApplyPhase defines model for TransactionApplyPhase.
type TransactionApplyPhase struct {
	Failure *Failure                `json:"failure,omitempty"`
	State   *ApplyPhaseState        `json:"state,omitempty"`
	Status  *TransactionPhaseStatus `json:"status,omitempty"`
}

// TransactionCommitPhase defines model for TransactionCommitPhase.
type TransactionCommitPhase struct {
	State  *CommitPhaseState       `json:"state,omitempty"`
	Status *TransactionPhaseStatus `json:"status,omitempty"`
}

// TransactionInitializePhase defines model for TransactionInitializePhase.
type TransactionInitializePhase struct {
	Failure *Failure                `json:"failure,omitempty"`
	State   *InitializePhaseState   `json:"state,omitempty"`
	Status  *TransactionPhaseStatus `json:"status,omitempty"`
}

// TransactionList defines model for TransactionList.
type TransactionList []Transaction

// TransactionPhaseStatus defines model for TransactionPhaseStatus.
type TransactionPhaseStatus struct {
	End   *End   `json:"end,omitempty"`
	Start *Start `json:"start,omitempty"`
}

// TransactionPhases defines model for TransactionPhases.
type TransactionPhases struct {
	Abort      *TransactionAbortPhase      `json:"abort,omitempty"`
	Apply      *TransactionApplyPhase      `json:"apply,omitempty"`
	Commit     *TransactionCommitPhase     `json:"commit,omitempty"`
	Initialize *TransactionInitializePhase `json:"initialize,omitempty"`
	Validate   *TransactionValidatePhase   `json:"validate,omitempty"`
}

// TransactionValidatePhase defines model for TransactionValidatePhase.
type TransactionValidatePhase struct {
	Failure *Failure                `json:"failure,omitempty"`
	State   *ValidatePhaseState     `json:"state,omitempty"`
	Status  *TransactionPhaseStatus `json:"status,omitempty"`
}

// TypeOpts defines model for TypeOpts.
type TypeOpts int32

// value represented as a byte array
type TypedValue struct {

	// the type for a value
	Type *ValueType `json:"type,omitempty"`

	// a set of type options
	TypeOpts *[]TypeOpts `json:"type_opts,omitempty"`

	// the bytes array
	Value *string `json:"value,omitempty"`
}

// ValidatePhaseState defines model for ValidatePhaseState.
type ValidatePhaseState string

// the type for a value
type ValueType string

// PatchTopLevelJSONBody defines parameters for PatchTopLevel.
type PatchTopLevelJSONBody PatchBody

// PatchTopLevelJSONRequestBody defines body for PatchTopLevel for application/json ContentType.
type PatchTopLevelJSONRequestBody PatchTopLevelJSONBody
