package commands

import "os"

func CommandExit(param ParamType) error {
	os.Exit(0)
	return nil
}
