// Copyright 2019 the Drone Authors. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// that can be found in the LICENSE file.

package plugin

import (
	"context"
	"runtime"
	"fmt"

	"github.com/drone/drone-go/drone"
	"github.com/drone/drone-go/plugin/converter"
)

// New returns a new conversion plugin.
func New() converter.Plugin {
	return &plugin{}
}

type plugin struct {
}

func (p *plugin) Convert(ctx context.Context, req *converter.Request) (*drone.Config, error) {

	// get the configuration file from the request.
	config := req.Config.Data

	// TODO this should be modified or removed. For
	// demonstration purposes we make a simple modification
	// to the configuration file and add a newline.
	config = config + "\n"
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
	// returns the modified configuration file.
	return &drone.Config{
		Data: config,
	}, nil
}
