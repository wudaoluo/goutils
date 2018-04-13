package serial

import (
	"fmt"
)

// Nodes is a custom flag Var representing a list of etcd nodes.
type Nodes []string

// String returns the string representation of a node var.
func (n *Nodes) String() string {
	return fmt.Sprintf("%s", *n)
}

// Set appends the node to the etcd node list.
func (n *Nodes) Set(node string) error {
	*n = append(*n, node)
	return nil
}


//使用方法
/*
go run node_var.go -node=127.0.0.1:6789 -node=127.0.0.1:1234
[127.0.0.1:6789 127.0.0.1:1234]


func main() {
	s
	flag.Var(&nodes, "node", "list of backend nodes")
	flag.Parse()

	fmt.Println(nodes.String())
}

*/

