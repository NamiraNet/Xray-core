package errors_test

import (
	"io"
	"strings"
	"testing"

	. "github.com/NamiraNet/xray-core/common/errors"
	"github.com/NamiraNet/xray-core/common/log"
	"github.com/google/go-cmp/cmp"
)

func TestError(t *testing.T) {
	err := New("TestError")
	if v := GetSeverity(err); v != log.Severity_Info {
		t.Error("severity: ", v)
	}

	err = New("TestError2").Base(io.EOF)
	if v := GetSeverity(err); v != log.Severity_Info {
		t.Error("severity: ", v)
	}

	err = New("TestError3").Base(io.EOF).AtWarning()
	if v := GetSeverity(err); v != log.Severity_Warning {
		t.Error("severity: ", v)
	}

	err = New("TestError4").Base(io.EOF).AtWarning()
	err = New("TestError5").Base(err)
	if v := GetSeverity(err); v != log.Severity_Warning {
		t.Error("severity: ", v)
	}
	if v := err.Error(); !strings.Contains(v, "EOF") {
		t.Error("error: ", v)
	}
}

func TestErrorMessage(t *testing.T) {
	data := []struct {
		err error
		msg string
	}{
		{
			err: New("a").Base(New("b")),
			msg: "common/errors_test: a > common/errors_test: b",
		},
	}

	for _, d := range data {
		if diff := cmp.Diff(d.msg, d.err.Error()); diff != "" {
			t.Error(diff)
		}
	}
}
