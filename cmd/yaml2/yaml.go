package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Coonfig struct {
	Hoge string `yaml:"hoge"`
}

func main() {
	var c Coonfig
	b, _ := os.ReadFile("config.yaml")
	yaml.Unmarshal(b, &c)
	fmt.Printf("%+v\n", c)
}
