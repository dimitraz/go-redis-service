package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

type Config struct {
	f interface{}
}

// As the logger can only be configured after we read the config
// I make use of the stdout for error logging
func (cfg Config) Init(fileName string) Config {
	start := time.Now()
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("%s \x1b[1;31m[%s] \x1b[0m : %s \n", "ERROR", err)
		// os.Exit(0)
	}

	err = json.Unmarshal(file, &cfg.f)
	if err != nil {
		fmt.Printf("%s \x1b[1;31m[%s] \x1b[0m : %s \n", "ERROR", err)
		// os.Exit(0)
	}

	fmt.Printf("%s \x1b[1;34m [%s] \x1b[0m  : %s \n", start.Format("2006/01/02 03:04:05"), "INFO", "Reading config settings")
	//fmt.Printf("%s \x1b[1;35m[%s] \x1b[0m : %s \n", start.Format(time.RFC3339), "DEBUG", cfg.f)
	return cfg
}

func (cfg Config) Get(parent string, name string) string {
	// manually decode hierarchy levels
	conf := cfg.f.(map[string]interface{})["object"]
	parentconf := conf.(map[string]interface{})[parent]
	retval := parentconf.(map[string]interface{})[name].(string)
	return retval
}
