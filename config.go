package configo

import (
	"encoding/json"
	"os"
	"path/filepath"
)

var c Configo

// Configo struct to be used as options
type Configo struct {
	environment string
	format      string
	path        string

	config map[string]interface{}
}

// SetEnvironment - set env var for deployment environment
func SetEnvironment(envar string) {
	c.SetEnvironment(envar)
}

// SetEnvironment - set env var for deployment environment
func (c *Configo) SetEnvironment(envar string) {
	env := getEnv(envar, "dev")
	c.environment = env
}

// SetFormat - set format of config file (.json)
func SetFormat(format string) {
	c.SetFormat(format)
}

// SetFormat - set format of config file (.json)
func (c *Configo) SetFormat(format string) {
	c.format = format
}

// SetPath - set path of config file
func SetPath(path string) {
	c.SetPath(path)
}

// SetPath - set path of config file
func (c *Configo) SetPath(path string) {
	c.path = path
}

// Source config file
func Source() (*Configo, error) {
	return c.Source()
}

// Source config file
func (c *Configo) Source() (*Configo, error) {

	path, _ := os.Getwd()
	file, err := os.Open(filepath.Join(path, c.path, c.environment+c.format))
	defer file.Close()
	if err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(file)
	// var config Config
	var config map[string]interface{}
	decoder.Decode(&config)
	if err != nil {
		return nil, err
	}
	c.config = config
	return c, nil
}

//Get - get value from config
func (c *Configo) Get(key string) string {
	return c.config[key].(string)
}

func getEnv(key, def string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return def
}
