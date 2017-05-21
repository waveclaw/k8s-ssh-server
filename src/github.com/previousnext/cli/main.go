package main

import (
	"fmt"
	"strings"

	"github.com/gosuri/uitable"
	"k8s.io/client-go/rest"

	sshclient "github.com/previousnext/client"
)

func main() {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	sshClient, err := sshclient.NewClient(config)
	if err != nil {
		panic(err)
	}

	list, err := sshClient.ListAll()
	if err != nil {
		panic(err)
	}

	table := uitable.New()
	table.MaxColWidth = 80

	table.AddRow("NAMESPACE", "NAME", "KEYS")
	for _, key := range list.Items {
		table.AddRow(key.Metadata.Namespace, key.Metadata.Name, strings.Join(key.Spec.AuthorizedKeys, "\n"))
	}
	fmt.Println(table)
}
