package testfn

import (
	"reflect"
	"testing"
)

func VerifyEq(t *testing.T, got, want any) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %v, want %v", got, want)
	}
}
