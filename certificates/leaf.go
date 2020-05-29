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

package certificates

import (
	"crypto/rsa"
)

// Leaf represents a leaf node in a certificate chain.
type Leaf struct {
	certificate []byte
	key         *rsa.PrivateKey
}

// EncodeCertificate returns the PEM encoded certificate for the leaf.
func (l *Leaf) EncodeCertificate() (string, error) {
	return EncodeCertificate(l.certificate)
}

// EncodeKey returns the PEM encoded key for the leaf.
func (l *Leaf) EncodeKey() (string, error) {
	return EncodeKey(l.key)
}
