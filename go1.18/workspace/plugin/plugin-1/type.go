package plugin_1

import "fmt"

type Plugin struct {
	Name string
}

func (p *Plugin) Hello() {
	fmt.Printf("Hello from out side plugin_1, %s", p.Name)
}
