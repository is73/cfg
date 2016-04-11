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

    fmt.Printf("\nValue of key3 is: %s\n",
    	config["key3"])
}