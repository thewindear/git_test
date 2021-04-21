package test

import (
	"test.com/m/v1/help"
	"testing"
)

func TestSayGood(t *testing.T) {
	if help.SayGood("hello") == "say: hell2o" {
		t.Fatalf("no ok")
	} else {
		t.Log("ok")
	}
}
