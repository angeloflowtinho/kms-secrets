package controllers

import (
	"testing"
)

func TestShasumData(t *testing.T) {
	expected := "b6b66b55b6b03c6ee6abc0027095d38a35937eb3e6ff2dc9f2aafa846c704e3b"
	data := map[string][]byte{
		"API_KEY":  []byte("hoge"),
		"PASSWORD": []byte("fuga"),
	}
	sum := shasumData(data)
	if sum != expected {
		t.Errorf("shasum is not matched, expected: %s, returned: %s", expected, sum)
	}
}
