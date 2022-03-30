package fibonacchi

import (
	"reflect"
	"testing"
)

func TestLoop(t *testing.T) {

	in := 9
	out := 21 //21

	t.Run("Testing", func(t *testing.T) {
		result, _ := Loop(in)

		if !reflect.DeepEqual(out, result) {
			t.Error("Expected", out, " got ", result)
		}

	})
}
