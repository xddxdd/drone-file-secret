// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE file.

package plugin

import (
	"context"
	"os"
	"path"
	"strings"

	"github.com/drone/drone-go/drone"
	"github.com/drone/drone-go/plugin/secret"
	"github.com/sirupsen/logrus"
)

// New returns a new secret plugin that sources secrets
// from the AWS secrets manager.
func New(path string) secret.Plugin {
	return &plugin{Path: path}
}

type plugin struct {
	Path string
}

func (p *plugin) Find(ctx context.Context, req *secret.Request) (*drone.Secret, error) {
	safePath := path.Join("/", req.Path)
	actualPath := path.Join(p.Path, safePath)
	logrus.Debugln(actualPath)

	name := req.Name
	if name == "" {
		name = "value"
	}

	data, err := os.ReadFile(actualPath)
	if err != nil {
		return nil, err
	}
	value := strings.TrimSpace(string(data))

	return &drone.Secret{
		Name: name,
		Data: value,
		Pull: false,
		Fork: false,
	}, nil
}
