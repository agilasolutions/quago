package configo

import (
	"io/ioutil"
	"testing"
)

//Init - Test run configo
func TestGet(t *testing.T) {
	//setup config options
	SetEnvironment("")
	SetFormat(".json")
	SetPath("")

	//create test dev config
	err := createTestConfig()
	if err != nil {
	}
	defer deleteTestConfig()
	//destroy test dev config
	config, _ := Source()
	if config.Get("testData") != "Hello World" {

	}
}

func createTestConfig() error {
	testData := []byte("{\"testData\":\"Hello World\"")

	// the WriteFile method returns an error if unsuccessful
	err := ioutil.WriteFile("dev.json", testData, 0777)
	// handle this error
	if err != nil {
		return err
	}

	return nil
}

func deleteTestConfig() {

}
