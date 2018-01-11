package main

import (
	"fmt"
	"flag"
	"encoding/xml"
	"os"
	"io/ioutil"
)

type Param struct {
	Short string
	Long string
	Description string
}

type Command struct  {
	Name string
	Description string
	Params []Param `xml:"param"`
}

type Pipeline struct {
	Commands []Command `xml:"command"`
}

var c Pipeline

func init() {
	xmlFile, err := os.Open("conf.xml")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer xmlFile.Close()
	b, _ := ioutil.ReadAll(xmlFile)

	xml.Unmarshal(b,&c)
	fmt.Printf("%+v\n",&c)
}

func main() {
	var command string
	flag.StringVar(&command, "c", "comandos", "specify the command you want to execute")
	flag.Parse()
	fmt.Printf("command = %s ", command)
	fmt.Printf("command in name in conf %s", &c.Commands[0])

} 