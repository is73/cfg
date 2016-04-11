###Simple config file reader for Golang applications


[![Build Status](https://travis-ci.org/is73/cfg.svg?branch=master)](https://travis-ci.org/is73/cfg) [![GoDoc](https://godoc.org/github.com/is73/cfg?status.svg)](https://godoc.org/github.com/is73/cfg)

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

See [config.txt](https://github.com/is73/cfg/blob/master/example/config.txt) in examples


**Reading config file:**

See [main.go](https://github.com/is73/cfg/blob/master/example/main.go) in examples

