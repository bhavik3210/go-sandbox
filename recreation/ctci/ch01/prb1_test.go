package ch01

import (
	"fmt"
	"testing"
)

func TestIsUnique1(t *testing.T) {
	got := IsUnique("potato")
	want := false

	if want != got {
		t.Errorf("isUniue1 test failed")
	} else {
		fmt.Println("isUnique1 test Passed")
	}
}

func TestIsUnique2(t *testing.T) {
	got := IsUnique("part")
	want := true

	if want != got {
		t.Errorf("isUniue2 test failed")
	}
}
