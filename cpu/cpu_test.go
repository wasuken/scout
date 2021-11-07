package cpu

import "testing"

func TestGetInfo(t *testing.T) {
	er, _ := GetInfo()
	if er != nil {
		t.Fatalf("failed cpu info. %#v", er)
	}
}
