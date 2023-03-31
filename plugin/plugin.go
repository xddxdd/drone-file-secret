// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE file.

package plugin

import (
	"context"

	"github.com/drone/drone-go/drone"
	"github.com/drone/drone-go/plugin/secret"
)

// New returns a new secret plugin that sources secrets
// from the AWS secrets manager.
func New() secret.Plugin {
	return &plugin{}
}

type plugin struct {}

func (p *plugin) Find(ctx context.Context, req *secret.Request) (*drone.Secret, error) {
	// path := req.Path
	name := req.Name
	if name == "" {
		name = "value"
	}

	value := "TODO"

	return &drone.Secret{
		Name: name,
		Data: value,
		Pull: true, // always true. use X-Drone-Events to prevent pull requests.
		Fork: true, // always true. use X-Drone-Events to prevent pull requests.
	}, nil
}
