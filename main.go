package main

import "fmt"

func findParent(data int) int {
	var res int
	if data%3 == 0 {
		res = data
	} else if (data+1)%3 == 0 {
		res = data + 1
	} else if (data-1)%3 == 0 {
		res = data - 1
	}

	return res / 3
}

func main() {

	node := 3135

	parent := findParent(node)

	fmt.Println(fmt.Sprintf(`Parent dari node %d adalah %d`, node, parent))

}
