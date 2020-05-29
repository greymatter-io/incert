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

// Batch returns a command that issues multiple certificates based upon a configuration file.
func Batch() *cobra.Command {

	command := &cobra.Command{
		Use:   "batch",
		Short: "Issue multiple certificates based upon a configuration file.",
		RunE: func(command *cobra.Command, args []string) error {

			file, err := command.Flags().GetString("config")
			if err != nil {
				return errors.Wrap(err, "error fetching [--config] flag")
			}

			viper.SetConfigFile(file)

			err = viper.ReadInConfig()
			if err != nil {
				return errors.Wrapf(err, "error reading configuration file %s", file)
			}

			var configuration config.Batch

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

			rootCertificate, _ := root.EncodeCertificate()
			intermediateCertificate, _ := intermediate.EncodeCertificate()

			directory, err := command.Flags().GetString("directory")
			if err != nil {
				return errors.Wrap(err, "error fetching [--directory] flag")
			}

			ioutil.WriteFile(filepath.Join(directory, "root.crt"), []byte(rootCertificate), 0644)
			ioutil.WriteFile(filepath.Join(directory, "intermediate.crt"), []byte(intermediateCertificate), 0644)

			for _, single := range configuration.Values {

				leaf, err := intermediate.Leaf(single.CommonName, single.AlternativeNames, single.Expires)
				if err != nil {
					return errors.Wrapf(err, "error creating leaf certificate")
				}

				leafCertificate, _ := leaf.EncodeCertificate()
				leafKey, _ := leaf.EncodeKey()

				certificateFile := filepath.Join(directory, fmt.Sprintf("%s.crt", single.CommonName))
				keyFile := filepath.Join(directory, fmt.Sprintf("%s.key", single.CommonName))

				ioutil.WriteFile(certificateFile, []byte(leafCertificate), 0644)
				ioutil.WriteFile(keyFile, []byte(leafKey), 0600)
			}

			return nil
		},
	}

	command.Flags().String("config", "", "the configuration file")
	command.Flags().StringP("directory", "d", "", "the directory for the output files")

	cobra.MarkFlagRequired(command.Flags(), "config")

	return command
}
