package main

import (
	"flag"

	"k8s.io/klog/v2"
)

// main, go run main.go -v 2
func main() {
	klog.InitFlags(nil)
	flag.Parse()
	klog.V(1).Info("yes")
	klog.V(2).Info("yes")
	klog.V(3).Info("yes")
	klog.V(4).Info("yes")
}
