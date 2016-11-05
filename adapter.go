package main

import (
	"github.com/hectane/go-smtpsrv"
)

// Adapter receives new emails and tweets their contents.
type Adapter struct {
	config *Config
	server *smtpsrv.Server
}

// run receives emails and tweets them. This is far more difficult than it
// sounds since the email must be decoded (headers removed, etc.).
func (a *Adapter) run() {
	for m := range a.server.NewMessage {
		//...
	}
}

// NewAdapter creates a new adapter.
func NewAdapter(config *Config) (*Adapter, error) {
	s, err := smtpsrv.NewServer(&smtpsrv.Config{
		Addr:   config.SMTPAddress,
		Banner: "epigeon",
	})
	if err != nil {
		return nil, err
	}
	a := &Adapter{
		config: config,
		server: s,
	}
	go a.run()
	return a, nil
}

// Close shuts down the adapter.
func (a *Adapter) Close() {
	a.server.Close(true)
}
