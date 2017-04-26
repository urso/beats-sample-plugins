package main

import (
	"github.com/elastic/beats/libbeat/outputs"
	"github.com/elastic/beats/libbeat/plugin"

	"github.com/urso/beats-sample-plugins/myconsole"
)

var Bundle = plugin.Bundle(
	outputs.Plugin("myconsole", myconsole.New),
)
