/*
Copyright 2019 The Kubernetes Authors.

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

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"

	"k8s.io/release/pkg/testgrid"
)

const usageFmt = `usage: %[1]s [release]
	e.g. %[1]s release-1.14
`

// TODO: deprecate this binary if releaselib.sh::set_build_version has been
// 		 ported to the `release` package
func main() {
	// Disable logrus logging
	logrus.SetOutput(ioutil.Discard)

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, usageFmt, filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	branch := os.Args[1]
	if branch == "master" {
		branch = "release-master"
	}

	tests, err := testgrid.New().BlockingTests(branch)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	for _, t := range tests {
		fmt.Println(t)
	}
}
