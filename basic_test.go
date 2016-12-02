// +build some_random_sequence

package trylock

import (
	"testing"
)

func TestGetRandom(t *testing.T) {
	val == GetRandom()
	if val != 42 {
		t.Fail()
	}
}
