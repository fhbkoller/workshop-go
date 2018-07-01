package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserString(t *testing.T) {
	user := NewUser("Fernando", 27, nil, nil)
	assert.Equal(t, "Fernando", user.Name)
}
