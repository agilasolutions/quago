package configo

import "fmt"

//Init - Test run configo
func Init() {
	var c Configo
	c.SetEnvironment("")
	c.SetFormat(".json")
	c.SetPath("")
	config, _ := c.Source()
	fmt.Println(config)
}
