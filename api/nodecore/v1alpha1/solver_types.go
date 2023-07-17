/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	resource "k8s.io/apimachinery/pkg/api/resource"
)

type Phase string

// PhaseStatus represents the status of a phase of the solver. I.e. the status of the REAR phases.
type PhaseStatus struct {
	Phase          Phase  `json:"phase"`
	Message        string `json:"message,omitempty"`
	StartTime      string `json:"startTime,omitempty"`
	LastChangeTime string `json:"lastChangeTime,omitempty"`
	EndTime        string `json:"endTime,omitempty"`
}

type FlavourSelector struct {
	Cpu              resource.Quantity `json:"cpu"`
	Memory           resource.Quantity `json:"memory"`
	Storage          resource.Quantity `json:"storage,omitempty"`
	EphemeralStorage resource.Quantity `json:"ephemeralStorage,omitempty"`
	Gpu              resource.Quantity `json:"gpu,omitempty"`
}

// SolverSpec defines the desired state of Solver
type SolverSpec struct {

	// Selector contains the flavour requirements for the solver.
	Selector FlavourSelector `json:"selector"`

	// IntentID is the ID of the intent that the Node Orchestrator is trying to solve.
	// It is used to link the solver with the intent.
	IntentID string `json:"intentID"`

	// FindCandidate is a flag that indicates if the solver should find a candidate to solve the intent.
	FindCandidate bool `json:"findCandidate,omitempty"`

	// StipulateContract is a flag that indicates if the solver should stipulate a contract with the candidate.
	// StipulateContract bool `json:"stipulateContract,omitempty"`

	// EnstablishPeering is a flag that indicates if the solver should enstablish a peering with the candidate.
	EnstablishPeering bool `json:"enstablishPeering,omitempty"`
}

// SolverStatus defines the observed state of Solver
type SolverStatus struct {

	// DiscoveryPhase describes the status of the discovery where the Discovery Manager
	// is looking for matching flavours outside the FLUIDOS Node
	DiscoveryPhase Phase `json:"discoveryPhase,omitempty"`

	// CandidatesPhase describes the status of the PeeringCandidates phase where the
	// Rear Manager is looking for the best candidate Flavour to solve the Node Orchestrator request.
	CandidatesPhase Phase `json:"candidatesPhase,omitempty"`

	// ReservationPhase describes the status of the Reservation phase where the Contract Manager
	// is reserving the resources on the candidate node.
	ReservationPhase Phase `json:"reservationPhase,omitempty"`

	// PurchasingPhase describes the status of the Purchasing phase where the Contract Manager
	// is purchasing the reserved resources on the candidate node.
	PurchasingPhase Phase `json:"purchasingPhase,omitempty"`

	// ConsumePhase describes the status of the Consume phase where the VFM (Liqo) is enstablishing
	// a peering with the candidate node.
	ConsumePhase Phase `json:"consumePhase,omitempty"`

	// SolverPhase describes the status of the Solver generated by the Node Orchestrator.
	// It is usefull to understand if the solver is still running or if it has finished or failed.
	SolverPhase PhaseStatus `json:"solverPhase,omitempty"`

	// PeeringCandidate contains the candidate that the solver has eventually found to solve the intent.
	PeeringCandidate GenericRef `json:"peeringCandidate,omitempty"`

	// Allocation contains the allocation that the solver has eventually created for the intent.
	// It can correspond to a virtual node
	// The Node Orchestrator will use this allocation to fullfill the intent.
	Allocation GenericRef `json:"allocation,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Solver is the Schema for the solvers API
type Solver struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SolverSpec   `json:"spec,omitempty"`
	Status SolverStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// SolverList contains a list of Solver
type SolverList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Solver `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Solver{}, &SolverList{})
}
