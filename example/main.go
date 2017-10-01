package main

import (
	"fmt"
	"log"

	"github.com/is73/cfg"
)

func main() {

	fmt.Println("Print particullar item (user name and surname):")
	config, err := cfg.Read("config.txt", "")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(config["user.name"], config["user.surname"])

	fmt.Println("\nOutput, prefix unused:")
	config, err = cfg.Read("config.txt", "")
	if err != nil {
		log.Fatalln(err)
	}

	for k, v := range config {
		fmt.Printf("'%s':'%s'\n", k, v)
	}

	fmt.Println("\nOutput, prefix used:")
	config, err = cfg.Read("config.txt", "user.")
	if err != nil {
		log.Fatalln(err)
	}

	for k, v := range config {
		fmt.Printf("'%s':'%s'\n", k, v)
	}
}

/*
Output, prefix unused:
'key1':'value1'
'key2':'value2'
'key3':'value3 value3		value3'
'user.id':'1'
'user.name':'john'
'user.surname':'doe'
'smtp_port':'25'
'smtp_password':'harDtoGueSs'
'smtp_user':'info@example.com'
'smtp_server':'smtp.example.com'

Output, prefix used:
'user.id':'1'
'user.name':'john'
'user.surname':'doe'
*/
