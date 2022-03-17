package main

import (
	"io/ioutil"
	"log"
	"os"
	"text/template"

	"gopkg.in/yaml.v3"
)

type bgp_conf struct {
	Globals struct {
		LocalAsn int    `yaml:"local_asn"`
		LocalRid string `yaml:"local_rid"`
	} `yaml:"globals"`
	BgpNeighbors map[interface{}]struct {
		Prefix    string `yaml:"prefix"`
		Remoteasn int    `yaml:"remoteasn"`
		Keepalive int    `yaml:"keepalive"`
		Holdtime  int    `yaml:"holdtime"`
	} `yaml:"bgp_neighbors"`
}

func main() {
	yangfile, err := ioutil.ReadFile("bgp_conf.yaml")

	if err != nil {
		log.Fatal(err)
	}

	var data bgp_conf

	err2 := yaml.Unmarshal(yangfile, &data)

	if err2 != nil {
		log.Fatal(err2)
	}

	// fmt.Println(data.Globals)

	// for _, v := range data.BgpNeighbors {
	// 	fmt.Printf("%v\n", v)
	// }

	temp, err := template.ParseFiles("bgp_template.txt")

	if err != nil {
		log.Fatalln(err)
	}

	err = temp.Execute(os.Stdout, data)
}
