package main

import (
	"fmt"
	"os"
	"os/exec"

	log "github.com/sirupsen/logrus"
)

const (
	npmLockFile = "package-lock.json"
	npmCommand = "npm"

	yarnLockFile = "yarn.lock"
	yarnCommand = "yarn"
)

func main() {
	args := os.Args[1:]
	log.SetLevel(log.ErrorLevel)

	if _, err := os.Stat(npmLockFile); !os.IsNotExist(err) { // check if npm lock exists in current directory
		log.Info("npm lock file found")

		// npm knows about "add", no need to replace with "install"
		cmd := exec.Command(npmCommand, args...)
		out, err := cmd.CombinedOutput()
		if err != nil {
			log.Error(err)
		}
		fmt.Printf("%s", out)
	} else if _, err := os.Stat(yarnLockFile); !os.IsNotExist(err) { // check if yarn lock exists in current directory
		log.Info("yarn lock file found")

		// replace install only if there's another argument/package name
		if len(args) > 1 {
			if args[0] == "install" {
				args[0] = "add"
			}
		}

		cmd := exec.Command(yarnCommand, args...)
		out, err := cmd.CombinedOutput()
		if err != nil {
			log.Error(err)
		}
		fmt.Printf("%s", out)
	} else { // throw error if there's no npm or yarn lock file in the current directory
		log.Fatal("no npm/yarn lock file found")
	}
}