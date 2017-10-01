### Simple config file reader for Golang applications


[![Build Status](https://travis-ci.org/is73/cfg.svg?branch=master)](https://travis-ci.org/is73/cfg) [![GoDoc](https://godoc.org/github.com/is73/cfg?status.svg)](https://godoc.org/github.com/is73/cfg)

**Users often forget quotes, separators, braces and don't want to learn
syntax of json, xml, yaml, toml, whateverML to create or edit config files.**

---

**Installation:**
```
go get github.com/is73/cfg
```

**Test:**
```
go test github.com/is73/cfg
```

**Features:**
* First word on line is a key name, can't contain white space
* Sections can be emulated using key prefix
* Key duplicity is not allowed
* Any white space after first word, space or tab (even multiple), is a key value separator
* Anything after separator except new line character is the value of the key
* Values may contain white space characters
* Lines without value are ignored
* Empty lines are ignored

**Config file example:**

See also
[config.txt](https://github.com/is73/cfg/blob/master/example/config.txt) and
[main.go](https://github.com/is73/cfg/blob/master/example/main.go)
in example folder

```bash
# this is a comment line
# file contains intentional spaces and tabs
key1 value1
key2	value2
key3		value3 value3	value3
# next line is ignored, value is missing
keynoval
user.id 1
user.name john
user.surname doe
# next empty line is ignored

smtp_server		smtp.example.com
smtp_port 		25
smtp_user		info@example.com
smtp_password		harDtoGueSs
```
