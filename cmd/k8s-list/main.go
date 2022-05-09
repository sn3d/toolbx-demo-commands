package main

import "fmt"

type Cluster struct {
	name string
	url  string
}

var clusters []Cluster = []Cluster{
	Cluster{name: "cluster-1", url: "https://cluster-1:6443"},
	Cluster{name: "cluster-2", url: "https://cluster-2:6443"},
}

func main() {
	fmt.Printf("organization.com k8s clusters\n")
	for _, c := range clusters {
		fmt.Printf("%s\t%s\n", c.name, c.url)
	}
}
