package quago

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

//Init - Test run quago
func TestGet(t *testing.T) {
	//setup config options
	SetEnvironment("")
	SetFormat(".json")
	SetPath("")

	//create test dev config
	err := createTestConfig()
	if err != nil {
		t.Error("Error Creating File")
	}
	defer deleteTestConfig(t)
	//destroy test dev config
	config, _ := Source()
	if config.Get("testData") != "Hello World" {
		t.Error("Error configuration value")
	}
}

func createTestConfig() error {
	fmt.Println("Creating test config...")
	testData := []byte("{\"testData\":\"Hello World\"}")

	// the WriteFile method returns an error if unsuccessful
	err := ioutil.WriteFile("dev.json", testData, 0777)
	// handle this error
	if err != nil {
		return err
	}

	return nil
}

func deleteTestConfig(t *testing.T) {

	err := os.Remove("dev.json")
	if err != nil {
		t.Error("Error Deleting File")
	}
}
