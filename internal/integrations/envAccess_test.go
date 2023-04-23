package integrations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnvAccess(t *testing.T) {

	err := LoadEnvVar()
	assert.NoError(t, err)
}
