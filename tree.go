// Copyright 2022 Lipatov Alexander

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// 	http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package treeprinter

// New creates new tree with root value.
func New(val any) *Node {
	return &Node{Value: val}
}

type Node struct {
	Value any
	root  *Node
	nodes []*Node
}

// Attach your leafs and branches with this method.
func (n *Node) Attach(nodes ...*Node) *Node {
	for _, node := range nodes {
		n.nodes = append(n.nodes, attach(node, n.root))
	}
	return n
}

func attach(node *Node, root *Node) *Node {
	node.root = root
	for i, sub := range node.nodes {
		node.nodes[i] = attach(sub, root)
	}
	return node
}

func (n *Node) String() string {
	return defaultRender(n)
}
