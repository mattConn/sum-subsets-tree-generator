package main

import (
//	"fmt"
	"os/exec"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type tree struct {
	Input, Output       []int
	SumTimeElapsed, PlotTimeElapsed, TotalTimeElapsed time.Duration
}

// make array-based binary tree for sums of subsets of input
func fillSumTree (input, output []int, n, i int) {
	// i = input index, n = output index
	if( i == len(input)){
		return
	}

	output[2*n+1] = output[n]+input[i] // left node
	output[2*n+2] = output[n] // right node

	// i+1 for each call to match depth
	fillSumTree(input,output,2*n+1,i+1) // descend left 
	fillSumTree(input,output,2*n+2,i+1) // descend right
}

func handler(w http.ResponseWriter, r *http.Request) {
	inputArr := strings.Split(r.FormValue("body"), " ")
	nums := []int{}
	for _, v := range inputArr {
		n, _ := strconv.Atoi(v)
		nums = append(nums, n)
	}

	t, _ := template.ParseFiles("index.html")

	output := make([]int, 1<<(len(nums)+1)-1) // accomodate all nodes for breadth-first representation

	sumStart := time.Now()
	fillSumTree(nums,output,0,0)
	sumFinish := time.Since(sumStart)

	outputString := []string{}
	for _,v := range output {
		outputString = append(outputString,strconv.Itoa(v))
	}

	plotStart := time.Now()
	//  err := exec.Command("./graph-plotter.py", outputString...).Run()
	exec.Command("./graph-plotter.py", outputString...).Run()
	plotFinish := time.Since(plotStart)
	/*
	if err != nil {
		fmt.Println(err)
	}
	*/

	totalTime := sumFinish + plotFinish

	t.Execute(w, tree{Input: nums, Output: output, SumTimeElapsed: sumFinish, PlotTimeElapsed: plotFinish, TotalTimeElapsed: totalTime})
}

func main() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	http.HandleFunc("/favicon.ico", func (w http.ResponseWriter, r *http.Request){})
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
