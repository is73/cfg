package cfg

import (
	"testing"
)

func Test_Read(t *testing.T) {
	var tests = []struct {
		key string
		val string
	}{
		{"key1", "value1"},
		{"key2", "value2"},
		{"key3", "value3"},
		{"key4", "value4"},
		{"key5", "value5 value5\tvalue5"},
		{"user.name", "username"},
		{"k", "v"},
	}
	config := Read("config.txt")
	for _, test := range tests {
		if test.val != config[test.key] {
			t.Errorf("For: %s expected: %s, got: %s", test.key, test.val, config[test.key])
		}
	}
}

func Benchmark_Read(b *testing.B) {
	_ = Read("config.txt")
}