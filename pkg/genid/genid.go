// Package genid provides the interface and implementation for generating shortid.
package genid

import (
	"time"

	"github.com/teris-io/shortid"
)

// Generator defines the interface for generating shortid.
// Using a hash strategy here rather than using the auto-increment id in database effectively
// decouple the storage layer and can easily replace SQLite with a KV database that does
// not support auto-increment ID.
type Generator interface {
	// Generate generates a shortid that may or may not correspond to k.
	Generate(k string) string
}

type generator struct {
	sid *shortid.Shortid
}

// NewGenerator creates a default implementation of Generator.
func NewGenerator() (Generator, error) {
	sid, err := shortid.New(1, shortid.DefaultABC, uint64(time.Now().Unix()))
	if err != nil {
		return nil, err
	}
	return &generator{sid: sid}, nil
}

// MustNewGenerator creates a default implementation of Generator that does not
// return errors, but instead a direct panic.
func MustNewGenerator() Generator {
	g, err := NewGenerator()
	if err != nil {
		panic(err)
	}
	return g
}

func (g *generator) Generate(k string) string {
	return g.sid.MustGenerate()
}
