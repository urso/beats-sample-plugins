package main

import (
	"github.com/elastic/beats/libbeat/outputs"
	"github.com/elastic/beats/libbeat/plugin"
	"github.com/elastic/beats/libbeat/processors"

	"github.com/urso/beats-sample-plugins/lookup/exec"
	"github.com/urso/beats-sample-plugins/my_drop_fields"
	"github.com/urso/beats-sample-plugins/myconsole"
)

var Bundle = plugin.Bundle(
	// bundle of processor plugins
	plugin.Bundle(
		processors.Plugin("lookup.exec", exec.New),
		processors.Plugin("my_drop_fields", my_drop_fields.New),
	),

	// bundle of output plugins
	plugin.Bundle(
		outputs.Plugin("myconsole", myconsole.New),
	),
)
