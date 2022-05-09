package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

type Kafka struct {
	name    string
	servers []string
}

var kafkaClusters []Kafka = []Kafka{
	Kafka{name: "cluster-1", servers: []string{"server1:9092", "server2:9092"}},
	Kafka{name: "cluster-2", servers: []string{"server3:9092", "server4:9092", "server5:9092"}},
}

func main() {
	pp := &cli.App{
		Name:        "storage-kafka",
		Description: "manipulate with organization kafka clusters",
		Commands: []*cli.Command{
			&cli.Command{
				Name:   "list",
				Action: list,
			},
			&cli.Command{
				Name:      "get",
				ArgsUsage: "<cluster>",
				Action:    get,
			},
		},
	}
	pp.Run(os.Args)
}

func list(ctx *cli.Context) error {
	fmt.Printf("organization.com Kafka clusters\n")
	for _, cluster := range kafkaClusters {
		fmt.Printf("%s\t%v\n", cluster.name, cluster.servers)
	}
	return nil
}

func get(ctx *cli.Context) error {
	toFind := ctx.Args().First()
	for _, cluster := range kafkaClusters {
		if cluster.name == toFind {
			fmt.Printf("name:    %s\n", cluster.name)
			fmt.Printf("servers: %s\n", cluster.servers)
			return nil
		}
	}

	fmt.Printf("unknown cluster\n")
	return nil
}
