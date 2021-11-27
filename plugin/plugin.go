// Copyright 2019 the Drone Authors. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// that can be found in the LICENSE file.

package plugin

import (
	"context"
	"fmt"
	"regexp"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/drone/drone-go/drone"
	"github.com/drone/drone-go/plugin/converter"
)

var Platform string
var Re = regexp.MustCompile("\nsteps:\n")

func init() {
	Platform = fmt.Sprintf("\n# drocopla\nplatform: %s/%s\nsteps:\n", runtime.GOOS, runtime.GOARCH)
}

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

	if !strings.Contains(config, "\nplatform:\n") {
		config = Re.ReplaceAllString(config, Platform)
		fmt.Printf("Replaced to: %s\n", config)
	}

	requestLogger.WithFields(logrus.Fields{
		"req":  req,
		"os":   runtime.GOOS,
		"arch": runtime.GOARCH,
	}).Infoln("initiated")

	// returns the modified configuration file.
	return &drone.Config{
		Data: config,
	}, nil
}
