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

import (
	"bytes"
	"fmt"
	"strings"
)

// IndentSize is the number of spaces per tree level.
var IndentSize = 3

// EdgeType represents edge definition for rendering.
type EdgeType string

var (
	EdgeTypeLink EdgeType = "│"
	EdgeTypeMid  EdgeType = "├──"
	EdgeTypeEnd  EdgeType = "└──"
)

func defaultRender(root *Node) string {
	return p.render(root)
}

var p *printer = &printer{}

type printer struct {
	buf        *bytes.Buffer
	levelTails []int
}

func (p *printer) render(root *Node) string {
	p.buf = new(bytes.Buffer)
	defer p.reset()

	p.buf.WriteString(fmt.Sprintf("%v\n", root.Value))

	p.renderNodes(0, root.nodes)

	return p.buf.String()
}

func (p *printer) reset() {
	p.buf.Reset()
	p.levelTails = nil
}

func (p *printer) renderNodes(level int, nodes []*Node) {
	for i, node := range nodes {
		edge := EdgeTypeMid
		if i == len(nodes)-1 {
			edge = EdgeTypeEnd
			p.levelTails = append(p.levelTails, level)
		}
		p.renderValue(level, edge, node)
		if len(node.nodes) > 0 {
			p.renderNodes(level+1, node.nodes)
		}
	}
}

func (p *printer) renderValue(level int, edge EdgeType, node *Node) {
	for i := 0; i < level; i++ {
		if isEnded(p.levelTails, i) {
			p.buf.WriteString(strings.Repeat(" ", IndentSize+1))
			continue
		}
		p.buf.WriteString(fmt.Sprintf("%s%s", EdgeTypeLink, strings.Repeat(" ", IndentSize)))
	}

	var val any
	lines := strings.Split(fmt.Sprintf("%v", node.Value), "\n")
	// If value does not contain multiple lines, return itself.
	if len(lines) < 2 {
		val = fmt.Sprint(node.Value)
	} else {
		// If value contains multiple lines,
		// generate a padding and prefix each line with it.
		pad := padding(level, node)

		for i := 1; i < len(lines); i++ {
			lines[i] = fmt.Sprintf("%s%s", pad, lines[i])
		}

		val = strings.Join(lines, "\n")
	}

	p.buf.WriteString(fmt.Sprintf("%s %v\n", edge, val))
}

func isEnded(levelsEnded []int, level int) bool {
	for _, l := range levelsEnded {
		if l == level {
			return true
		}
	}
	return false
}

// padding returns a padding for the multiline values with correctly placed link edges.
func padding(level int, node *Node) string {
	links := make([]string, level+1)

	for node.root != nil {
		if isLast(node) {
			links[level] = strings.Repeat(" ", IndentSize+1)
		} else {
			links[level] = fmt.Sprintf("%s%s", EdgeTypeLink, strings.Repeat(" ", IndentSize))
		}
		level--
		node = node.root
	}

	return strings.Join(links, "")
}

// isLast checks if the nodes is the last one in the slice of its parent children
func isLast(n *Node) bool {
	return n == n.root.getLastChild()
}

func (n *Node) getLastChild() *Node {
	if len(n.nodes) == 0 {
		return nil
	}
	return n.nodes[len(n.nodes)-1]
}
