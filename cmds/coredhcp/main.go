// Copyright 2018-present the CoreDHCP Authors. All rights reserved
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package main

import (
	"time"

	"github.com/atpons/coredhcp"
	"github.com/atpons/coredhcp/config"
	"github.com/atpons/coredhcp/logger"
	_ "github.com/atpons/coredhcp/plugins/dns"
	_ "github.com/atpons/coredhcp/plugins/file"
	_ "github.com/atpons/coredhcp/plugins/netmask"
	_ "github.com/atpons/coredhcp/plugins/range"
	_ "github.com/atpons/coredhcp/plugins/router"
	_ "github.com/atpons/coredhcp/plugins/server_id"
)

func main() {
	logger := logger.GetLogger("main")
	config, err := config.Load()
	if err != nil {
		logger.Fatal(err)
	}
	server := coredhcp.NewServer(config)
	if err := server.Start(); err != nil {
		logger.Fatal(err)
	}
	if err := server.Wait(); err != nil {
		logger.Print(err)
	}
	time.Sleep(time.Second)
}
