package common

import (
	"testing"
)

func TestDoSome(t *testing.T) {
	Warning.Println("adffdf")
	Warning.Printf("%v %v", "abc", "bbb")
}
