// Copyright 2019 the Drone Authors. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// that can be found in the LICENSE file.

package plugin

import (
	"context"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"

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
	// check type is yaml
	if strings.HasSuffix(req.Repo.Config, ".yml") == false {
		return nil, nil
	}

	requestLogger := logrus.WithFields(logrus.Fields{
		"repo_name": req.Repo.Name,
	})

	// get the configuration file from the request.
	config := req.Config.Data

	requestLogger.WithFields(logrus.Fields{
		"req":  req,
		"os":   runtime.GOOS,
		"arch": runtime.GOARCH,
	}).Infoln("initiated")
	// TODO this should be modified or removed. For
	// demonstration purposes we make a simple modification
	// to the configuration file and add a newline.
	config = config + "\n"
	// returns the modified configuration file.
	return &drone.Config{
		Data: config,
	}, nil
}
