package assert

import (
	"reflect"
	"testing"
)

func False(tb testing.TB, actual bool) {
	Equals(tb, false, actual)
}

func True(tb testing.TB, actual bool) {
	Equals(tb, true, actual)
}

func Equals[V any](tb testing.TB, expected, actual V) {
	if !reflect.DeepEqual(expected, actual) {
		tb.Logf("\n\texpected: %v\n\tgot: %v", expected, actual)
		tb.FailNow()
	}
}
