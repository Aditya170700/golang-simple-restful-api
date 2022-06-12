package test

import (
	"sudutkampus/gorestfulapi/dependency"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleService(t *testing.T) {
	simpleService, err := dependency.InitializedService(false)
	assert.False(t, simpleService.Error)
	assert.Nil(t, err)
}

func TestSimpleServiceError(t *testing.T) {
	simpleService, err := dependency.InitializedService(true)
	assert.NotNil(t, err)
	assert.Nil(t, simpleService)
}
