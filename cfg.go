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
		pair := re.FindStringSubmatch(line)
		counter++

		if pair != nil && strings.HasPrefix(pair[0], "#") == false {
			key := strings.TrimSpace(pair[1])
			val := strings.TrimSpace(pair[2])
			if _, ok := config[key]; ok {
				// duplicity, return error
				err := fmt.Errorf("Duplicity for key: %s on line %d", key, counter)
				return nil, err
			}
			config[key] = val
		}
	}

	return config, nil
}
