package eventbus_test

import (
	"testing"

	"github.com/yj12138/eventbus"
)

func TestOne(t *testing.T) {
	eventbus.ListenOne(1, func(data bool) {
		if data != true {
			t.Errorf(" test failed")
		}
	})
	eventbus.EmitOne(1, true)
}

func TestTwo(t *testing.T) {
	eventbus.ListenTwo(1, func(a bool, b float32) {
		if a != false || b != 1.32 {
			t.Errorf("test failed")
		}
	})
	eventbus.EmitTwo(1, false, float32(1.32))
}
