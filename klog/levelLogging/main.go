package main

import (
	"flag"
	"fmt"

	"k8s.io/klog/v2"
)

type mystruct struct {
	Name  string
	other []*mystruct
}

// main, go run main.go -v 2
func main() {
	klog.InitFlags(nil)
	flag.Parse()
	klog.V(1).Info("yes")
	klog.V(2).Info("yes")
	klog.V(3).Info("yes")
	klog.V(4).Info("yes")

	klog.V(1).Info("===================")

	err := fmt.Errorf("new Error")
	err = fmt.Errorf("wrapper err %w", err)

	a := mystruct{Name: "1st layer"}
	b := mystruct{Name: "2nd layer", other: []*mystruct{&a}}
	c := mystruct{Name: "3rd layer", other: []*mystruct{&b}}
	d := mystruct{Name: "4th layer", other: []*mystruct{&c}}

	klog.V(1).InfoS("", d)
	fmt.Printf("%v", d)
}
