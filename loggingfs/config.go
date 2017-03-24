package loggingfs

import (
	"fmt"
	"io/ioutil"
	yaml "launchpad.net/goyaml"
	osusers "os/user"
	"regexp"
	"strconv"
)

type LoggerFsConfig struct {
	Mount_point string

	Logs       []LogOptions
	Transports []TransportOptions
}

type LogOptions struct {
	Name  string
	File  string
	Owner string
	Uid   int
	Gid   int

	properties map[string]string
}

type TransportOptions struct {
	Handler string
	Options map[string]interface{}
}

type LogConfig struct {
	options LogOptions

	pattern regexp.Regexp
}

func load_config(filename string) (config *LoggerFsConfig, err error) {
	fData, err := ioutil.ReadFile(filename)

	if err != nil {
		return
	}

	config = new(LoggerFsConfig)
	err = yaml.Unmarshal([]byte(fData), config)

	if err != nil {
		return
	}

	fmt.Printf("--- config:\n%v\n\n", config)

	// Fill in required fields and generate errors as necessary

	for i := 0; i < len(config.Logs); i++ {
		err = log_config(&config.Logs[i])
		if err != nil {
			return
		}
	}

	for i := 0; i < len(config.Transports); i++ {
		err = transport_config(&config.Transports[i])
		if err != nil {
			return
		}
	}

	return
}

func log_config(logOpts *LogOptions) (err error) {
	err = nil

	if len(logOpts.Owner) == 0 {
		user, lookup_err := osusers.LookupId(fmt.Sprint(logOpts.Uid))
		if lookup_err != nil {
			err = lookup_err
			return
		}
		logOpts.Owner = user.Username
	} else {
		user, lookup_err := osusers.Lookup(logOpts.Owner)
		if lookup_err != nil {
			err = lookup_err
			return
		}
		logOpts.Uid, err = strconv.Atoi(user.Uid)

		if err != nil {
			return
		}
	}

	return
}

func transport_config(tOpts *TransportOptions) (err error) {

	if len(tOpts.output) == 0 {
		tOpts.output = "zmq"
	}

	return
}
