// Licensed under the MIT license, see LICENCE file for details.

package quicktest_test

import (
	"errors"
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestBadCheckf(t *testing.T) {
	err := qt.BadCheckf("bad %s", "wolf")
	expectedMessage := "bad wolf"
	if err.Error() != expectedMessage {
		t.Fatalf("error:\ngot  %q\nwant %q", err, expectedMessage)
	}
}

func TestIsBadCheck(t *testing.T) {
	err := qt.BadCheckf("bad wolf")
	assertBool(t, qt.IsBadCheck(err), true)
	err = errors.New("bad wolf")
	assertBool(t, qt.IsBadCheck(err), false)
}

func TestSilentFailuref(t *testing.T) {
	err := qt.SilentFailure()
	if err.Error() != "" {
		t.Fatalf("error:\ngot  %q\nwant empty string", err)
	}
}

func TestIsSilentFailure(t *testing.T) {
	err := qt.SilentFailure()
	assertBool(t, qt.IsSilentFailure(err), true)
	err = errors.New("bad wolf")
	assertBool(t, qt.IsSilentFailure(err), false)
}
