package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNewFrequency(t *testing.T) {
	testPath := filepath.Clean("../test.txt")

	file, err := os.Open(testPath)
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	freq := NewFrequency(file)

	wantFreq := Frequency{
		'X': 333,
		't': 223000,
	}

	for k, expectedVal := range wantFreq {
		if freq[k] != expectedVal {
			t.Errorf("expected %s:%v, got frequency %v", string(k), expectedVal, freq[k])
		}
	}
}
