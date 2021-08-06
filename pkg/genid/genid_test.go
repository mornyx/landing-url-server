package genid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenID(t *testing.T) {
	g := MustNewGenerator()
	assert.NotPanics(t, func() {
		for n := 0; n < 10; n++ {
			assert.NotEmpty(t, g.Generate(""))
		}
	})
}
