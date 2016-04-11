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
func Read(fileName string) (map[string]string, error) {

	f, err := os.Open(fileName)
	if err != nil {
		log.Panic("Config file error: ", err)
	}
	defer f.Close()

	config := make(map[string]string, 20)
	re := regexp.MustCompile(`([^\s]+)[\s\t]+([^\n]+)`)
	lines := bufio.NewScanner(f)
	counter := 0

	for lines.Scan() {

		line := strings.TrimSpace(lines.Text())
		kv := re.FindStringSubmatch(line)
		counter++

		if kv != nil && strings.HasPrefix(kv[0], "#") == false {
			k := strings.TrimSpace(kv[1])
			v := strings.TrimSpace(kv[2])
			if _, ok := config[k]; ok {
				// duplicity, return error
				err := fmt.Errorf("Duplicity for key: %s on line %d", k, counter)
				return nil, err
			}
			config[k] = v
		}
	}

	return config, nil
}
