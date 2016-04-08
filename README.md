**Simplest config file reader for Golang applications**


**Installation:**
```
go get github.com/is73/cfg
```


**Config file example:**  
* First word is key name, can't contain whitespace  
* Any whitespace after first word, space or tab (even multiple) is key value separator  
* Anything after separator is value of key  
* Lines without value are ignored
* Empty lines are ignored
```bash
# sample config file comment
key1	value1
key2 value2
# line without value is ignored
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

	fmt.Println("Printing all:")
	for k, v := range config {
		fmt.Printf("'%s':'%s'\n", k, v)
	}

	fmt.Println("\nPrinting particular key:")
	fmt.Printf("Value of key3 is: %s\n", config["key3"])
}
```
