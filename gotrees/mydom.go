// <html>
//   <h1 id="heading">Hello, World!</h1>
//   <p>This is a simple HTML document.</p>
// </html>

package main

import (
	"fmt"
	"strings"
)

type Node struct {
	tag      string
	text     string
	id       string
	classes  string
	src      string
	alt      string
	children []*Node
}

func main() {
	image := Node{
		tag: "img",
		src: "http://example.com/logo.svg",
		alt: "Example's Logo",
	}

	p := Node{
		tag:      "p",
		text:     "And this is some text in a paragraph. And next to it there's an image.",
		children: []*Node{&image},
	}

	span := Node{
		tag:  "span",
		id:   "copyright",
		text: "2019 &copy; Ilija Eftimov",
	}

	div := Node{
		tag:      "div",
		classes:  "footer",
		text:     "This is the footer of the page.",
		children: []*Node{&span},
	}

	h1 := Node{
		tag:  "h1",
		text: "This is a H1",
	}

	body := Node{
		tag:      "body",
		children: []*Node{&h1, &p, &div},
	}

	html := Node{
		tag:      "html",
		children: []*Node{&body},
	}
	getIdByDFS(&html, "copyright")
	fmt.Println(dfsNode)

	n := getIdByBFS(&html, "copyright")
	fmt.Println(n)

	c1 := findAllByClassName(&html, "footer")
	for _, c := range c1 {
		fmt.Println(c)
	}
}
func findAllByClassName(root *Node, className string) []*Node {
	result := make([]*Node, 0)
	queue := make([]*Node, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		nextUp := queue[0]
		queue = queue[1:]
		if nextUp.hasClass(className) {
			result = append(result, nextUp)
		}
		if len(nextUp.children) > 0 {
			queue = append(queue, nextUp.children...)
		}
	}
	return result
}

func (n *Node) hasClass(className string) bool {
	classes := strings.Fields(n.classes)
	for _, class := range classes {
		if class == className {
			return true
		}
	}
	return false
}

func getIdByBFS(root *Node, id string) *Node {
	queue := make([]*Node, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		nextUp := queue[0]
		queue = queue[1:]

		if nextUp.id == id {
			return nextUp
		}

		if len(nextUp.children) > 0 {
			queue = append(queue, nextUp.children...)
		}
	}
	return nil
}

var dfsNode Node

func getIdByDFS(node *Node, id string) *Node {
	if node.id == id {
		dfsNode = *node
		return node
	}

	if len(node.children) > 0 {
		for _, child := range node.children {
			getIdByDFS(child, id)
		}
	}
	return nil
}
