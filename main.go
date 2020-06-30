package main

import (
	"os/exec"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// make array-based binary tree for sums of subsets of input
	// input: array of integers from which subsets will be summed
	// binTree: breadth-first representation of binary tree of sums
	// nodeIndex: index of binTree
	// inputIndex: index of input 
func fillSumTree (input, binTree []int, nodeIndex, inputIndex int) {
	// return when at end of input
	if( inputIndex == len(input)){ return }

	binTree[2*nodeIndex+1] = binTree[nodeIndex]+input[inputIndex] // left node = current node + current input
	binTree[2*nodeIndex+2] = binTree[nodeIndex] // right node = current node

	// inputIndex+1 for each call to maintain uniform depth on descend
	fillSumTree(input,binTree,2*nodeIndex+1,inputIndex+1) // descend left 
	fillSumTree(input,binTree,2*nodeIndex+2,inputIndex+1) // descend right
}

// handle form input
func handler(w http.ResponseWriter, r *http.Request) {
	// get form input "body", which will be a string of space-separated numbers
	inputArr := strings.Split(r.FormValue("body"), " ")

	nums := []int{} // array of input to send to fillSumTree
	// convert strings in inputArr to integers and append to nums
	for _, v := range inputArr {
		n, _ := strconv.Atoi(v)
		nums = append(nums, n)
	}

	// binary tree array to hold sums of all subets.
	// Length will accomodate all nodes for breadth-first representation
		// |bin. sum tree| = 2**(|input|+1) - 1
			// e.g. |input| = 2
			// => |len of bin. tree| = (2**3)-1 = 7
	binTree := make([]int, 1<<(len(nums)+1)-1)

	fillSumTree(nums,binTree,0,0)

	binTreeString := []string{}
	for _,v := range binTree {
		binTreeString = append(binTreeString,strconv.Itoa(v))
	}

	t, _ := template.ParseFiles("index.html")
	exec.Command("./graph-plotter.py", binTreeString...).Run()

	t.Execute(w, struct { Input, Output []int }{nums, binTree})
}

func main() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	http.HandleFunc("/favicon.ico", func (w http.ResponseWriter, r *http.Request){})
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
