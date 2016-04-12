package cfg

import (
	"fmt"
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
		{"key6.user", "username"},
		{"key6.db", "mydb"},
	}
	config, err := Read("config.txt", "")
	if err != nil {
		fmt.Println(err)
	}
	for _, test := range tests {
		if test.val != config[test.key] {
			t.Errorf("For: %s expected: %s, got: %s", test.key, test.val, config[test.key])
		}
	}
	for k, v := range config {
		fmt.Printf("'%s':'%s'\n", k, v)
	}
}

func Benchmark_Read(b *testing.B) {
	_, _ = Read("config.txt", "")
}