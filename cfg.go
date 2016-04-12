// Package cfg provides function(s) to read specific key/value configuration
// file format similar to .INI files for Go applications
package cfg

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

// Read reads config file and returns its contents
// as map a.k.a associative array of keys and values
// keyPrefix is optional, default ""
func Read(fileName, keyPrefix string) (map[string]string, error) {

	f, err := os.Open(fileName)
	if err != nil {
		log.Panic(err)
	}
	defer f.Close()

	config := make(map[string]string, 20)
	kvPatt := fmt.Sprintf(`(%s[^\s]+)[\s\t]+([^\n]+)`, keyPrefix)
	re := regexp.MustCompile(kvPatt)
	lines := bufio.NewScanner(f)
	ln := 0

	for lines.Scan() {
		line := strings.TrimSpace(lines.Text())
		kv := re.FindStringSubmatch(line)
		ln++

		if kv != nil && strings.HasPrefix(kv[0], "#") == false {
			k := strings.TrimSpace(kv[1])
			v := strings.TrimSpace(kv[2])
			if _, ok := config[k]; ok {
				err := fmt.Errorf("Dup key: %s, line %d", k, ln)
				return nil, err
			}
			config[k] = v
		}
	}
	return config, nil
}
