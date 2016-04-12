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
		{"key3", "value3 value3\tvalue3"},
		{"user.id", "1"},
		{"user.name", "john"},
		{"user.surname", "doe"},
		{"smtp_server", "smtp.example.com"},
		{"smtp_port", "25"},
		{"smtp_user", "info@example.com"},
		{"smtp_password", "harDtoGueSs"},
	}
	config, err := Read("example/config.txt", "")
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
	_, _ = Read("example/config.txt", "")
}