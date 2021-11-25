package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsFalse(t *testing.T) {
	assert.False(t, IsTrue(false))
}

func TestIsTrue(t *testing.T) {
	assert.True(t, IsTrue(true))
}
