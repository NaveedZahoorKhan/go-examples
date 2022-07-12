// <html>
//   <h1 id="heading">Hello, World!</h1>
//   <p>This is a simple HTML document.</p>
// </html>

package main

import "fmt"

type Node struct {
	tag      string
	text     string
	id       string
	class    string
	children []*Node
}

func main() {
	p := Node{
		tag:  "p",
		text: "This is a simple HTML document.",
	}
	h1 := Node{
		tag:  "h1",
		text: "Hello, World!",
		id:   "heading",
	}

	html := Node{
		tag: "html",
		children: []*Node{
			&p, &h1,
		},
	}

	n := getIdByBFS(&html, "heading")
	fmt.Println(n)
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
			for _, child := range nextUp.children {
				queue = append(queue, child)
			}
		}
	}
	return nil
}
