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

const PlatformFmt = `
# drocopla
platform:
  os: %s
  arch: %s

steps:
`

var Platform string
var Re = regexp.MustCompile("\nsteps:\n")

func init() {
	Platform = fmt.Sprintf(PlatformFmt, runtime.GOOS, runtime.GOARCH)
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

	requestLogger := logrus.WithField("repo_name", req.Repo.Name)
	requestLogger.WithFields(logrus.Fields{
		"os":   runtime.GOOS,
		"arch": runtime.GOARCH,
	}).Infoln("initiated")

	// get the configuration file from the request.
	config := req.Config.Data

	if strings.Contains(config, "\nplatform:\n") {
		// platform is set, skip
		return nil, nil
	}

	config = Re.ReplaceAllString(config, Platform)
	requestLogger.WithField("result", config).Debug("replaced")

	// returns the modified configuration file.
	return &drone.Config{
		Data: config,
	}, nil
}
