package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

var testOk string = `1
2
3
3
3
4
5`

var testOkResult string = `1
2
3
4
5
`

func TestOk(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(testOk))
	out := new(bytes.Buffer)
	err := Unique(in, out)
	if err != nil {
		t.Errorf("test for OK failed")
	}
	result := out.String()
	if result != testOkResult {
		t.Errorf("test for OK failed - results bad rly bad %v %v", result, testOkResult)
	}
}

var testFail = `1
2
1`

func TestForError(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(testFail))
	out := new(bytes.Buffer)
	err := Unique(in, out)
	if err == nil {
		t.Errorf("test for FAIL failed")
	}
}
