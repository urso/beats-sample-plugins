package main

import (
	"github.com/elastic/beats/libbeat/plugin"
	"github.com/elastic/beats/libbeat/processors"

	"github.com/urso/beats-sample-plugins/my_drop_fields"
)

var Bundle = plugin.Bundle(
	processors.Plugin("my_drop_fields", my_drop_fields.New),
)
