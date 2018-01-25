package main

import (
	"fmt"

	//"github.com/elastic/beats/libbeat/cmd/instance"
	"github.com/tinytub/examplebeat/beater"
	"github.com/tinytub/examplebeat/cmd/instance"
	// import modules
	//_ "github.com/elastic/beats/metricbeat/include"
	_ "github.com/tinytub/examplebeat/include"
)

func main() {

	/*
		if err := cmd.RootCmd.Execute(); err != nil {
			os.Exit(1)
		}
	*/

	err := instance.Run("metricbeat", "metricbeat", "", beater.DefaultCreator())
	if err != nil {
		fmt.Println(err)
	}
}
