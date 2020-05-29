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

package issue

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"time"

	"github.com/greymatter-io/incert/certificates"
	"github.com/greymatter-io/incert/config"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Single returns a command that issues a certificate based upon command line arguments.
func Single() *cobra.Command {

	command := &cobra.Command{
		Use:   "single",
		Short: "Issue a certificate based upon command line arguments.",
		RunE: func(command *cobra.Command, args []string) error {

			file, err := command.Flags().GetString("config")
			if err != nil {
				return errors.Wrap(err, "error fetching [--config] flag")
			}

			if file != "" {
				viper.SetConfigFile(file)

				err = viper.ReadInConfig()
				if err != nil {
					return errors.Wrapf(err, "error reading configuration file %s", file)
				}
			}

			var configuration config.Single

			err = viper.Unmarshal(&configuration)
			if err != nil {
				return errors.Wrapf(err, "error unmarshalling configuration")
			}

			timestamp := time.Now().Unix()

			root, err := certificates.Root(fmt.Sprintf("Incert Root (%d)", timestamp))
			if err != nil {
				return errors.Wrapf(err, "error creating root certificate authority")
			}

			intermediate, err := root.Branch(fmt.Sprintf("Incert Intermediate (%d)", timestamp))
			if err != nil {
				return errors.Wrapf(err, "error creating intermediate certificate authority")
			}

			leaf, err := intermediate.Leaf(configuration.CommonName, configuration.AlternativeNames, configuration.Expires)
			if err != nil {
				return errors.Wrapf(err, "error creating leaf certificate")
			}

			rootCertificate, _ := root.EncodeCertificate()
			intermediateCertificate, _ := intermediate.EncodeCertificate()
			leafCertificate, _ := leaf.EncodeCertificate()
			leafKey, _ := leaf.EncodeKey()

			directory, err := command.Flags().GetString("directory")
			if err != nil {
				return errors.Wrap(err, "error fetching [--directory] flag")
			}

			rootFile := filepath.Join(directory, "root.crt")
			intermediateFile := filepath.Join(directory, "intermediate.crt")
			certificateFile := filepath.Join(directory, fmt.Sprintf("%s.crt", configuration.CommonName))
			keyFile := filepath.Join(directory, fmt.Sprintf("%s.key", configuration.CommonName))

			ioutil.WriteFile(rootFile, []byte(rootCertificate), 0644)
			ioutil.WriteFile(intermediateFile, []byte(intermediateCertificate), 0644)
			ioutil.WriteFile(certificateFile, []byte(leafCertificate), 0644)
			ioutil.WriteFile(keyFile, []byte(leafKey), 0600)

			return nil
		},
	}

	command.Flags().String("config", "", "configuration file")

	command.Flags().DurationP("expires", "e", (time.Hour * 24 * 365), "duration for which the certificate is valid.")
	command.Flags().StringP("common", "c", "localhost", "common name of the certificate")
	command.Flags().StringSliceP("alternative", "a", []string{"localhost"}, "subject alternative name of the certificate")
	command.Flags().StringP("directory", "d", "", "the directory for the output files")

	viper.BindPFlag("expires", command.Flags().Lookup("expires"))
	viper.BindPFlag("commonName", command.Flags().Lookup("common"))
	viper.BindPFlag("alternativeNames", command.Flags().Lookup("alternative"))

	return command
}
