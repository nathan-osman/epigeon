package main

import (
	"github.com/ChimeraCoder/anaconda"
	"github.com/hectane/go-smtpsrv"

	"errors"
	"io"
	"io/ioutil"
	"log"
	"mime"
	"mime/multipart"
	"net/mail"
	"strings"
)

// Adapter receives new emails and tweets their contents.
type Adapter struct {
	config *Config
	server *smtpsrv.Server
	api    *anaconda.TwitterApi
}

// getBody recursively searches for text content, decodes, and returns it.
func (a *Adapter) getBody(header mail.Header, r io.Reader) (string, error) {
	contentType := header.Get("Content-Type")
	if contentType == "" {
		return "", errors.New("invalid content type")
	}
	mediaType, params, err := mime.ParseMediaType(contentType)
	if err != nil {
		return "", err
	}
	if strings.HasPrefix(mediaType, "text/") {
		b, err := ioutil.ReadAll(r)
		if err != nil {
			return "", err
		}
		return string(b), nil
	}
	if strings.HasPrefix(mediaType, "multipart/") {
		var (
			boundary = params["boundary"]
			mr       = multipart.NewReader(r, boundary)
		)
		for {
			p, err := mr.NextPart()
			if err == io.EOF {
				break
			}
			b, err := a.getBody(mail.Header(p.Header), p)
			if err != nil {
				continue
			}
			return b, nil
		}
	}
	return "", errors.New("no valid text stream found")
}

// tweetBody sends a tweet with the specified body.
func (a *Adapter) tweetBody(body string) error {
	if len(body) > 140 {
		body = body[:139] + "…"
	}
	_, err := a.api.PostTweet(body, nil)
	return err
}

// run receives emails and tweets them. This is far more difficult than it
// sounds since the email must be decoded (headers removed, etc.).
func (a *Adapter) run() {
	for m := range a.server.NewMessage {
		log.Printf("message received from %s", m.From)
		r := strings.NewReader(m.Body)
		m, err := mail.ReadMessage(r)
		if err != nil {
			log.Print(err)
			continue
		}
		b, err := a.getBody(m.Header, m.Body)
		if err != nil {
			log.Print(err)
			continue
		}
		if err := a.tweetBody(b); err != nil {
			log.Print(err)
		}
		log.Print("tweet sent")
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
		api:    anaconda.NewTwitterApi(config.TwitterAccessToken, config.TwitterAccessSecret),
	}
	go a.run()
	return a, nil
}

// Close shuts down the adapter.
func (a *Adapter) Close() {
	a.server.Close(true)
}
