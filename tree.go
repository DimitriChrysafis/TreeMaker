package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

type Tree struct {
	Root *Node
}

func main() {
	fmt.Println("Welcome to the Tree Generator!")
	var depth int
	fmt.Println("Enter the depth of the tree:")
	fmt.Scanln(&depth)

	var tree Tree
	tree.generateRandomTree(depth)

	for {
		fmt.Println("Your tree:")
		displayTree(tree.Root, "")

		var choice string
		fmt.Println("Do you want to add a child node? (yes/no):")
		fmt.Scanln(&choice)

		if choice == "yes" {
			var side string
			fmt.Println("Which side do you want to add the child to? (left/right):")
			fmt.Scanln(&side)

			var targetNode *Node
			fmt.Println("Enter the data of the target node:")
			var targetData int
			fmt.Scanln(&targetData)

			if side == "left" {
				targetNode = findNode(tree.Root, targetData)
				if targetNode != nil {
					targetNode.Left = &Node{Data: getNextNumber()}
					fmt.Printf("Added node %d to the left of node %d\n", targetNode.Left.Data, targetNode.Data)
					saveGraphText(tree.Root)
					saveGraphDOT(tree.Root)
				}
			} else if side == "right" {
				targetNode = findNode(tree.Root, targetData)
				if targetNode != nil {
					targetNode.Right = &Node{Data: getNextNumber()}
					fmt.Printf("Added node %d to the right of node %d\n", targetNode.Right.Data, targetNode.Data)
					saveGraphText(tree.Root)
					saveGraphDOT(tree.Root)
				}
			} else {
				fmt.Println("Invalid choice.")
			}
		} else if choice == "no" {
			fmt.Println("Okay. Saving the current graph to graph.txt and graph.dot.")
			saveGraphText(tree.Root)
			saveGraphDOT(tree.Root)
			break
		} else {
			fmt.Println("Invalid choice.")
		}
	}
}

func saveGraphText(root *Node) {
	file, err := os.Create("graph.txt")
	if err != nil {
		fmt.Println("Error creating graph text file:", err)
		return
	}
	defer file.Close()

	saveGraphTextRecursive(file, root, "")
	fmt.Println("Graph text file saved as graph.txt")
}

func saveGraphTextRecursive(file *os.File, node *Node, prefix string) {
	if node == nil {
		return
	}

	if node.Left != nil {
		file.WriteString(fmt.Sprintf("%s%d -> %d\n", prefix, node.Data, node.Left.Data))
		saveGraphTextRecursive(file, node.Left, prefix+"  ")
	}

	if node.Right != nil {
		file.WriteString(fmt.Sprintf("%s%d -> %d\n", prefix, node.Data, node.Right.Data))
		saveGraphTextRecursive(file, node.Right, prefix+"  ")
	}
}

func saveGraphDOT(root *Node) {
	// Code for saving DOT file...
}

func findNode(node *Node, data int) *Node {
	if node == nil {
		return nil
	}
	if node.Data == data {
		return node
	}
	if found := findNode(node.Left, data); found != nil {
		return found
	}
	if found := findNode(node.Right, data); found != nil {
		return found
	}
	return nil
}

func (tree *Tree) generateRandomTree(depth int) {
	if depth <= 0 {
		return
	}

	rand.Seed(time.Now().UnixNano())

	tree.Root = &Node{Data: 1}
	tree.Root.Left = &Node{Data: 2}
	tree.Root.Right = &Node{Data: 3}

	tree.Root.Left.generateRandomSubtree(depth-1, 2)
	tree.Root.Right.generateRandomSubtree(depth-1, 2)
}

func (node *Node) generateRandomSubtree(depth int, count int) {
	if depth <= 0 {
		return
	}

	node.Left = &Node{Data: getNextNumber()}
	node.Right = &Node{Data: getNextNumber()}

	node.Left.generateRandomSubtree(depth-1, count)
	node.Right.generateRandomSubtree(depth-1, count)
}

func displayTree(node *Node, prefix string) {
	if node == nil {
		return
	}

	if node.Right != nil {
		displayTree(node.Right, prefix+"        ")
	}

	fmt.Print(prefix)
	fmt.Println(" |—-", node.Data)

	if node.Left != nil {
		displayTree(node.Left, prefix+"        ")
	}
}

var counter int

func getNextNumber() int {
	counter++
	return counter
}
