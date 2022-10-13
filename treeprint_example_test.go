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

package treeprinter_test

import (
	"fmt"
	tp "github.com/tellmeac/tree-printer"
	"testing"
)

func TestTreePrintExample(t *testing.T) {
	tree := tp.New(".").Attach(
		tp.New("Successful ideas"),
		tp.New("Memes").Attach(
			tp.New("Bunny"),
			tp.New("Funny Cat"),
			tp.New("Mystical wise tree"),
			tp.New("PHP development"),
			tp.New("Others").Attach(
				tp.New("Monkey"),
				tp.New("Dad joke"),
				tp.New("Something"),
			),
		),
	)

	fmt.Print(tree.String())
}

func TestTreePrintExample_SingleNode(t *testing.T) {
	tree := tp.New("Hello")

	fmt.Print(tree.String())
}
