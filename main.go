package main

import (
	"fmt"
	ansible "hafifa/src/ansible"
	config "hafifa/src/config"
	"os"
)

func invalidParameters() {
	colorReset := "\033[0m"

	colorRed := "\033[31m"
	colorGreen := "\033[32m"
	fmt.Println(string(colorRed), `You entered an invalid value.
Usage: hafifa [stage] [field]
stage: The stage of the internship you're in. For example, DevOps.
field: The field of the stage you're in. For example, Ansible in DevOps.`,
		"\n", string(colorReset), string(colorGreen),
		"Omer Ben David TEAM CASA 2021", string(colorReset))
}

func main() {
	if len(os.Args) != 3 {
		invalidParameters()
		os.Exit(-1)
	}
	stage := os.Args[1]
	field := os.Args[2]
	if config.IsStage(stage) {
		if config.IsFieldInStage(stage, field) {
			//fmt.Println(stage, field)
			ansible.RunPlaybook(stage, field)
		} else {
			invalidParameters()
		}
	} else if stage == "config" {
		config.HandleConfig()
	} else {
		invalidParameters()
	}
}
