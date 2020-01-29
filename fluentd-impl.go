package log

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/fluent/fluent-logger-golang/fluent"
)

// FluentClient manage all fluent action
type FluentClient struct {
	Config *Fluent
	Client *fluent.Fluent
}

// NewFluent function return a new fluent client based on singleton pattern
func NewFluent(config *Fluent) ILog {
	currentSession := &FluentClient{nil, nil}

	logger, err := fluent.New(getConfig(config))
	if err != nil {
		log.Println("Error when try to connect to Fluent server: ", err)
		panic(err)
	}

	currentSession = &FluentClient{config, logger}
	log.Println("Connected to Fluent Server")

	return currentSession
}

// Write log to Fluent server
func (f *FluentClient) Write(p []byte) (n int, err error) {
	data := make(map[string]interface{})
	err = json.Unmarshal(p, &data)
	if err != nil {
		return 0, err
	}

	if f.Client == nil {
		fmt.Print(string(p))
		return 0, err
	}

	err = f.Client.Post(f.Config.Tag, data)
	if err != nil {
		return 0, err
	}
	return len(p), err
}

// Close Fluent connection
func (f *FluentClient) Close() {
	f.Client.Close()
}

// getConfig function return config of fluent
func getConfig(f *Fluent) fluent.Config {
	return fluent.Config{
		FluentPort:         f.Port,
		FluentHost:         f.Host,
		TagPrefix:          f.Prefix,
		MarshalAsJSON:      false,
		SubSecondPrecision: true,
	}
}
