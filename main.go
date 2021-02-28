// Copyright (C) 2021 The Dank Grinder authors.
//
// This source code has been released under the GNU Affero General Public
// License v3.0. A copy of this license is available at
// https://www.gnu.org/licenses/agpl-3.0.en.html

package main

import (
	"math/rand"
	"os"
	"path"
	"path/filepath"
	"sync"
	"time"

	"github.com/shiena/ansicolor"

	"github.com/dankgrinder/dankgrinder/config"
	"github.com/dankgrinder/dankgrinder/discord"
	"github.com/sirupsen/logrus"
)

var cfg config.Config
var ins []*instance
var ex string

// im pretty sure this thing loads the config, or is it just for errors?
func main() {
	logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true})
	logrus.SetOutput(ansicolor.NewAnsiColorWriter(os.Stdout))

	var err error
	ex, err = os.Executable()
	if err != nil {
		logrus.Fatalf("could not find executable path: %v", err)
	}
	ex = filepath.ToSlash(ex)
	cfg, err = config.Load(path.Dir(ex))
	if err != nil {
		logrus.Fatalf("could not load config: %v", err)
	}
	if cfg.Features.Debug {
		logrus.SetLevel(logrus.DebugLevel)
	}
	if cfg.Features.LogToFile {
		logrus.AddHook(logFileHook{dir: path.Dir(ex)})
	}
	if err = cfg.Validate(); err != nil {
		logrus.Fatalf("invalid config: %v", err)
	}
	if cfg.Compat.AwaitResponseTimeout < 3 {
		logrus.Warnf("await response timeout is less than 3, this might cause stability issues for responses")
	}
	if len(cfg.InstancesOpts) > 1 {
		logrus.Infof("more than 1 instance configured, starting in swarm mode")
	}

	rand.Seed(time.Now().UnixNano())
	//... hum .. logger? but for erors(mabye the logging is done in logger.go?)
	wg := &sync.WaitGroup{}
	for _, opts := range cfg.InstancesOpts {
		client, err := discord.NewClient(opts.Token)
		if err != nil {
			logrus.Errorf("error while creating client: %v", err)
			continue
		}
		logrus.Infof("successful authorization as %v", client.User.Username+"#"+client.User.Discriminator)
		ins = append(ins, &instance{
			client:    client,
			channelID: opts.ChannelID,
			cmds:      cmds(),
			shifts:    opts.Shifts,
			wg:        wg,
		})
	}

	for _, in := range ins {
		if err = in.start(); err != nil {
			logrus.Fatalf("error while starting instance: %v", err)
		}
	}
	wg.Wait()
	logrus.Fatalf("no running instances left")
}
