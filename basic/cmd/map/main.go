package main

import "fmt"

func main() {
	scene := make(map[string]int)

	// 准备map数据
	scene["route"] = 66
	scene["brazil"] = 4
	scene["china"] = 960

	delete(scene, "brazdsfil")

	fmt.Println(len(scene))

	for k, v := range scene {
		fmt.Println(k, v)
	}
}
