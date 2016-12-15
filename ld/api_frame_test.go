package ld_test

import (
	"github.com/stretchr/testify/assert"
	. "github.com/verisart/json-gold-light/ld"
	"testing"
)

func TestGetFrameFlag(t *testing.T) {
	assert.Equal(t, true, GetFrameFlag(
		map[string]interface{}{"test": []interface{}{true, false}},
		"test",
		false,
	),
	)

	assert.Equal(t, true, GetFrameFlag(
		map[string]interface{}{
			"test": map[string]interface{}{
				"@value": true,
			},
		},
		"test",
		false,
	),
	)

	assert.Equal(t, true, GetFrameFlag(
		map[string]interface{}{"test": true},
		"test",
		false,
	),
	)

	assert.Equal(t, false, GetFrameFlag(
		map[string]interface{}{"test": "not_boolean"},
		"test",
		false,
	),
	)
}
