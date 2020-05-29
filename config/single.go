// Copyright 2019 Decipher Technology Studios
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config

import "time"

// Single represents the configuration for a single certificate.
type Single struct {

	// AlternativeNames defines the subject alternative names for the certificate.
	AlternativeNames []string `json:"alternativeNames" mapstructure:"alternativeNames" yaml:"alternativeNames"`

	// CommonName defines the subject common name for the certificate.
	CommonName string `json:"commonName" mapstructure:"commonName" yaml:"commonName"`

	// Expires defines the duration for which the certificate is valid.
	Expires time.Duration `json:"expires" mapstructure:"expires" yaml:"expires"`
}
