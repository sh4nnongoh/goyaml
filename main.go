package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

func main() {
	data, err := ioutil.ReadFile("./main.yaml")
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(string(data))

	// Unmarshal yaml file into struct
	t := Yaml{}
	{
		err = yaml.Unmarshal(data, &t)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		fmt.Printf("--- t:\n%v\n\n", t)

		d, err := yaml.Marshal(&t)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		fmt.Printf("--- t dump:\n%s\n\n", string(d))
	}
	/*
		for _, e := range t.Specification.Endpoints {
			fmt.Println(e.Name)
		}
	*/
	/*
		m := make(map[interface{}]interface{})

		err = yaml.Unmarshal([]byte(data), &m)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		fmt.Printf("--- m:\n%v\n\n", m)

		d, err = yaml.Marshal(&m)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		fmt.Printf("--- m dump:\n%s\n\n", string(d))
	*/
}
