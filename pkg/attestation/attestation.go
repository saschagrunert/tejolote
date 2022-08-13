/*
Copyright 2022 Adolfo García Veytia

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

package attestation

import (
	intoto "github.com/in-toto/in-toto-golang/in_toto"
	slsa "github.com/in-toto/in-toto-golang/in_toto/slsa_provenance/v0.2"
)

type Attestation intoto.Statement

func New() *Attestation {
	attestation := &Attestation{
		StatementHeader: intoto.StatementHeader{
			Type:          intoto.StatementInTotoV01,
			PredicateType: slsa.PredicateSLSAProvenance,
			Subject:       []intoto.Subject{},
		},
	}
	return attestation
}

func (att *Attestation) SLSA() *Attestation {
	att.Predicate = NewSLSAPredicate()
	return att
}

func NewSLSAPredicate() *slsa.ProvenancePredicate {
	// invocation, err := r.InvocationData()
	predicate := slsa.ProvenancePredicate{
		Builder: slsa.ProvenanceBuilder{
			ID: "", // TODO: Read builder from trusted environment
		},
		BuildType: "",
		Invocation: slsa.ProvenanceInvocation{
			ConfigSource: slsa.ConfigSource{
				URI:        "",
				Digest:     map[string]string{},
				EntryPoint: "",
			},
			Parameters:  nil,
			Environment: nil,
		},
		BuildConfig: nil,
		Metadata: &slsa.ProvenanceMetadata{
			BuildInvocationID: "",
			BuildStartedOn:    nil,
			BuildFinishedOn:   nil,
			Completeness: slsa.ProvenanceComplete{
				Parameters:  true,
				Environment: false,
				Materials:   false,
			},
			Reproducible: false,
		},
		Materials: []slsa.ProvenanceMaterial{},
	}

	return &predicate
}