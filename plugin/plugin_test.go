// Copyright 2019 the Drone Authors. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// that can be found in the LICENSE file.

package plugin

import (
	"context"
	"fmt"
	"testing"

	"github.com/drone/drone-go/drone"
	"github.com/drone/drone-go/plugin/converter"
)

func TestPlugin(t *testing.T) {
	t.Skip()

}

func TestPluginReplace(t *testing.T) {

	yaml := `---
kind: pipeline
type: docker
name: app

steps:

- name: deploy_local
`
	rv, err := New().Convert(context.Background(), &converter.Request{
		Repo:   drone.Repo{Config: ".done.yml"},
		Config: drone.Config{Data: yaml},
	})
	if err == nil {
		fmt.Printf("rv: %v", rv.Data)

	}
}
