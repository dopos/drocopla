// Copyright 2019 the Drone Authors. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// that can be found in the LICENSE file.

package plugin

import (
	"context"
	"fmt"
	"runtime"
	"testing"

	//	"github.com/sirupsen/logrus"

	"github.com/drone/drone-go/drone"
	"github.com/drone/drone-go/plugin/converter"
)

func TestPluginReplaceOKSingle(t *testing.T) {

	in := `---
kind: pipeline
type: docker
name: app

steps:

- name: deploy_local
`
	out := fmt.Sprintf(`---
kind: pipeline
type: docker
name: app

# drocopla
platform:
  os: %s
  arch: %s

steps:

- name: deploy_local
`, runtime.GOOS, runtime.GOARCH)
	//logrus.SetLevel(logrus.DebugLevel)
	rv, err := New().Convert(context.Background(), &converter.Request{
		Repo:   drone.Repo{Config: ".done.yml"},
		Config: drone.Config{Data: in},
	})
	if err != nil {
		t.Error(err)
	}
	if rv.Data != out {
		t.Errorf("Result not equal.\nwant:\n%s\ngot:\n%s\n", out, rv.Data)
	}
}

func TestPluginNoReplace(t *testing.T) {

	in := `---
platform:
`
	//logrus.SetLevel(logrus.DebugLevel)
	rv, err := New().Convert(context.Background(), &converter.Request{
		Repo:   drone.Repo{Config: ".done.yml"},
		Config: drone.Config{Data: in},
	})
	if err != nil {
		t.Error(err)
	}
	if rv != nil {
		t.Errorf("Result must be nil. Got: %+v", rv)
	}
}
