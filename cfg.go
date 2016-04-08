package cfg

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strings"
)

// Read reads config file and returns its contents
// as map a.k.a associative array of keys and values
func Read(fileName string) map[string]string {
	f, err := os.Open(fileName)
	if err != nil {
		log.Panic("Config file error: ", err)
	}
	defer f.Close()

	config := make(map[string]string, 20)
	re := regexp.MustCompile(`[\s\t]+`)
	lines := bufio.NewScanner(f)

	for lines.Scan() {
		line := strings.TrimSpace(lines.Text())
		if !strings.HasPrefix(line, "#") && len(line) >= 3 {
			pair := re.Split(line, 2)
			key := strings.TrimSpace(pair[0])
			val := strings.TrimSpace(pair[1])
			config[key] = val
		}
	}

	return config
}
