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

package versions

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCommit(t *testing.T) {

	Convey("When versions.Commit is invoked", t, func() {

		commit = "8eb820dc5f7b138bb6e431da80de973d"

		Convey("it returns the value of versions.commit", func() {
			So(Commit(), ShouldEqual, commit)
		})
	})
}

func TestVersion(t *testing.T) {

	Convey("When versions.Version is invoked", t, func() {

		version = "259df6532a074762e1f63e9f16e1229d"

		Convey("it returns the value of versions.version", func() {
			So(Version(), ShouldEqual, version)
		})
	})
}
