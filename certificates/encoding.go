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
	"bytes"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"

	"github.com/pkg/errors"
)

// EncodeCertificate PEM encodes a DER certificate.
func EncodeCertificate(certificate []byte) (string, error) {

	buffer := bytes.NewBuffer(nil)

	err := pem.Encode(buffer, &pem.Block{Type: "CERTIFICATE", Bytes: certificate})
	if err != nil {
		return "", errors.Wrap(err, "error encoding certificate")
	}

	return buffer.String(), nil
}

// EncodeKey PEM encodes an RSA key.
func EncodeKey(key *rsa.PrivateKey) (string, error) {

	buffer := bytes.NewBuffer(nil)

	err := pem.Encode(buffer, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)})
	if err != nil {
		return "", errors.Wrap(err, "error encoding key")
	}

	return buffer.String(), nil
}
