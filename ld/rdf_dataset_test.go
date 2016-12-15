package ld_test

import (
	"github.com/stretchr/testify/assert"
	. "github.com/verisart/json-gold-light/ld"
	"testing"
)

func TestGetCanonicalDouble(t *testing.T) {
	assert.Equal(t, "5.3E0", GetCanonicalDouble(5.3))
}
