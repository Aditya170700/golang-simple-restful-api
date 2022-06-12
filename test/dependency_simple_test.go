package test

import (
	"fmt"
	"sudutkampus/gorestfulapi/dependency"
	"sudutkampus/gorestfulapi/helper"
	"testing"
)

func TestSimpleService(t *testing.T) {
	simpleService, err := dependency.InitializedService()
	helper.PanicIfError(err)
	fmt.Println(simpleService.Error)
}
