###Simple config file reader for Golang applications

[![GoDoc](https://godoc.org/github.com/is73/cfg?status.svg)](https://godoc.org/github.com/is73/cfg)

*Ordinary users forget quotes, separators, braces and don't want to learn
syntax of json, xml, yaml, toml, whateverML to create or edit config files.*

---

**Installation:**
```
go get github.com/is73/cfg
```

**Test:**
```
go test github.com/is73/cfg
```

**Config file example:**
* First word on line is a key name, can't contain white space
* Key duplicity is not allowed
* Any white space after first word, space or tab (even multiple), is a key value separator
* Anything after separator is value of the key
* Values may contain white space characters
* Lines without value are ignored
* Empty lines are ignored

```bash
# a comment line
key1	value1
key2 value2
# next line is ignored, value is missing
keyX
# empty line is ignored

key3 value3
key4 value4
key5 value5 value5	value5

key6.user username
key6.db		mydb
```


**Reading config file:**
```go
package main

import (
	"fmt"
	"github.com/is73/cfg"
)

func main() {

	config, err := cfg.Read("config.txt")
	if err != nil {
		fmt.Println(err)
	}

	for k, v := range config {
		fmt.Printf("'%s':'%s'\n", k, v)
	}

	fmt.Printf("\nValue of key3 is: %s\n", config["key3"])
}

/*
Sample output:
'key1':'value1'
'key6.user':'username'
'key6.db':'mydb'
'key2':'value2'
'key3':'value3'
'key4':'value4'
'key5':'value5 value5	value5'

Value of key3 is: value3
*/
```
