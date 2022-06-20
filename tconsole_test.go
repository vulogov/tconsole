package tconsole

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_tconsole1(t *testing.T) {
	c, err := New()
	assert.Equal(t, err, nil)
	assert.NotEqual(t, c, nil)
}
