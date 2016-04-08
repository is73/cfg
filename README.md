**Simple config file reader for Golang applications**


**Installation:**
```
go get github.com/is73/cfg
```


**Config file example:**  
first word is key name  
any whitespace, space or tab is key value separator  
anything after separator is value of key  
```
# a comment
key1	value1
key2 value2
# next line is ignored
d
# empty line is ignored

key3 value3
key4 value4
key5 value5 value5	value5
```


**Reading config file:**
```go
package main

import (
	"fmt"
	"github.com/is73/cfg"
)
func main() {
	config := cfg.Read("config.txt")
	for k, v := range config {
		fmt.Printf("'%s':'%s'\n", k, v)
	}
}
```
