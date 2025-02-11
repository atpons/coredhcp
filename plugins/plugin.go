// Copyright 2018-present the CoreDHCP Authors. All rights reserved
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package plugins

import (
	"github.com/atpons/coredhcp/handler"
	"github.com/atpons/coredhcp/logger"
)

var log = logger.GetLogger("plugins")

// Plugin represents a plugin object.
// Setup6 and Setup4 are the setup functions for DHCPv6 and DHCPv4 handlers
// respectively. Both setup functions can be nil.
type Plugin struct {
	Name   string
	Setup6 SetupFunc6
	Setup4 SetupFunc4
}

// RegisteredPlugins maps a plugin name to a Plugin instance.
var RegisteredPlugins = make(map[string]*Plugin)

// SetupFunc6 defines a plugin setup function for DHCPv6
type SetupFunc6 func(args ...string) (handler.Handler6, error)

// SetupFunc4 defines a plugin setup function for DHCPv6
type SetupFunc4 func(args ...string) (handler.Handler4, error)

// RegisterPlugin registers a plugin by its name and setup functions.
func RegisterPlugin(name string, setup6 SetupFunc6, setup4 SetupFunc4) {
	log.Printf("Registering plugin \"%s\"", name)
	if _, ok := RegisteredPlugins[name]; ok {
		// TODO this highlights that asking the plugins to register themselves
		// is not the right approach. Need to register them in the main program.
		log.Panicf("Plugin '%s' is already registered", name)
	}
	plugin := Plugin{
		Name:   name,
		Setup6: setup6,
		Setup4: setup4,
	}
	RegisteredPlugins[name] = &plugin
}
