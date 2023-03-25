package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestName(t *testing.T) {
	a, b := 1, 2
	assert.Equal(t, a, b, "a应该等于b")
}

func TestNil(t *testing.T) {
	var a int
	assert.Nil(t, a)
	var b interface{}
	assert.Nil(t, b)
}

func TestNotNil(t *testing.T) {
	var b int
	assert.NotNil(t, b)
}

func TestEmpty(t *testing.T) {
	var a int
	assert.Empty(t, a, "输出")
}
