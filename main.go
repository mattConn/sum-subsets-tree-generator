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
	const MAX_INPUT_LEN int = 10 // max length of input sequence

	// get form input "body", which will be a string of space-separated numbers
	inputArr := strings.Split(r.FormValue("body"), " ")

	// if input > max input, exit early
	if (len(inputArr) > MAX_INPUT_LEN){
		// parse and execute template
		t, _ := template.ParseFiles("index.html")
		t.Execute(w, struct { Input, Output []int }{[]int{-1}, []int{-1}})

		return
	 }

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

	fillSumTree(nums,binTree,0,0) // make binary sum tree

	// cast binTree to array of strings to pass to python script
	binTreeString := []string{}
	for _,v := range binTree {
		binTreeString = append(binTreeString,strconv.Itoa(v))
	}

	// parse and execute template
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, struct { Input, Output []int }{nums, binTree})

	// exec python script (this is done at the end as this will take the longest)
	exec.Command("./graph-plotter.py", binTreeString...).Run()

}

func main() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	http.HandleFunc("/favicon.ico", func (w http.ResponseWriter, r *http.Request){})
	http.HandleFunc("/", handler)
	println("Listening on :8081...")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
