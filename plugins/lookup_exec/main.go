package main

import (
	"github.com/elastic/beats/libbeat/plugin"
	"github.com/elastic/beats/libbeat/processors"

	"github.com/urso/beats-sample-plugins/lookup/exec"
)

var Bundle = plugin.Bundle(
	processors.Plugin("lookup.exec", exec.New),
)
