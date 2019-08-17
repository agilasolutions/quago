package configo

import "fmt"

//Init - Test run configo
func Init() {
	SetEnvironment("")
	SetFormat(".json")
	SetPath("")
	config, _ := Source()
	fmt.Println(config)
}
