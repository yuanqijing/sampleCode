package plugin_1

import "fmt"

type Plugin struct {
	Name string
}

func (p *Plugin) Hello() {
	fmt.Printf("Hello from plugin_1,  %s", p.Name)
}
