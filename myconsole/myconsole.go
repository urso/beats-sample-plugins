package myconsole

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"

	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/common/op"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/libbeat/outputs"
)

type console struct {
	config config
	out    *os.File
}

type config struct {
	Pretty bool `config:"pretty"`
}

var (
	defaultConfig = config{
		Pretty: false,
	}
)

func New(_ common.BeatInfo, config *common.Config) (outputs.Outputer, error) {
	c := &console{config: defaultConfig, out: os.Stdout}
	err := config.Unpack(&c.config)
	if err != nil {
		return nil, err
	}

	// check stdout actually being available
	if runtime.GOOS != "windows" {
		if _, err = c.out.Stat(); err != nil {
			return nil, fmt.Errorf("console output initialization failed with: %v", err)
		}
	}

	return c, nil
}

func (c *console) Close() error {
	return nil
}

func (c *console) PublishEvent(
	s op.Signaler,
	opts outputs.Options,
	data outputs.Data,
) error {
	var jsonEvent []byte
	var err error

	if c.config.Pretty {
		jsonEvent, err = json.MarshalIndent(data.Event, "", "  ")
	} else {
		jsonEvent, err = json.Marshal(data.Event)
	}
	if err != nil {
		logp.Err("Fail to convert the event to JSON (%v): %#v", err, data.Event)
		op.SigCompleted(s)
		return err
	}

	if err = c.writeBuffer(jsonEvent); err != nil {
		goto fail
	}
	if err = c.writeBuffer([]byte{'\n'}); err != nil {
		goto fail
	}

	op.SigCompleted(s)
	return nil
fail:
	if opts.Guaranteed {
		logp.Critical("Unable to publish events to console: %v", err)
	}
	op.SigFailed(s, err)
	return err
}

func (c *console) writeBuffer(buf []byte) error {
	written := 0
	for written < len(buf) {
		n, err := c.out.Write(buf[written:])
		if err != nil {
			return err
		}

		written += n
	}
	return nil
}
