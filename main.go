// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE file.

package main

import (
	"net/http"

	"github.com/drone/drone-file-secret/plugin"
	"github.com/drone/drone-go/plugin/secret"

	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"

	_ "github.com/joho/godotenv/autoload"
)

type config struct {
	Address  string `envconfig:"DRONE_BIND"`
	Debug    bool   `envconfig:"DRONE_DEBUG"`
	Secret   string `envconfig:"DRONE_SECRET"`
	BasePath string `envconfig:"DRONE_BASE_PATH"`
}

func main() {
	spec := new(config)
	err := envconfig.Process("", spec)
	if err != nil {
		logrus.Fatal(err)
	}

	if spec.Debug {
		logrus.SetLevel(logrus.DebugLevel)
	}
	if spec.Secret == "" {
		logrus.Fatalln("missing secret key")
	}
	if spec.Address == "" {
		spec.Address = ":3000"
	}
	if spec.BasePath == "" {
		logrus.Fatalln("missing secret base path")
	}

	http.Handle("/", secret.Handler(
		spec.Secret,
		plugin.New(spec.BasePath),
		logrus.StandardLogger(),
	))

	var g errgroup.Group

	g.Go(func() error {
		logrus.Infof("server listening on address %s", spec.Address)
		return http.ListenAndServe(spec.Address, nil)
	})

	if err := g.Wait(); err != nil {
		logrus.Fatal(err)
	}
}
