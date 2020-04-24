package main

import (
	"os"

	"github.com/GeneralElectric/GABeat/cmd"
	"github.com/elastic/beats/libbeat/logp"
)

func main() {
	logp.Info("Starting GA Beat...")
	//err := beat.Run("gabeat", "", beater.New)
	err := cmd.RootCmd.Execute()
	os.Exit(getExitStatus(err))
}

func getExitStatus(err error) (status int) {
	if err != nil {
		logp.Info("Stopping GA Beat with error status...")
		return 1
	}
	logp.Info("Stopping GA Beat normally")
	return 0
}
