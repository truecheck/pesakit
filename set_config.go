package pesakit

import (
	"fmt"
	"github.com/pesakit/pesakit/home"
	"github.com/spf13/cobra"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

//
func (app *App) persistentPreRunE(cmd *cobra.Command, args []string) error {
	return nil
}

// setConfig returns a home directory of the application, and a config file path.
// if successful.
// the two flags --config and --home are used to determine the config file path and
// home directory.
// There are possible four scenarios:
//
// 		1. --config is not set, --home is set
// 		2. --config is set, --home is not set
// 		3. --config is not set, --home is not set
// 		4. --config is set, --home is set
//
// In the first scenario, the config file path is calculated by appending the string
// ".pesakit.env" to the home directory that is set by --home. if the home directory
// is not available or not a directory, the function returns an error.
//
// In the second scenario, the config file path is set. If the file is not available,
// the function returns an error. if available it will be copied to the home directory
// of the application. and the new config file path is set to $APP_HOME/{config_file_name}.
//
// In the third scenario, the config file path is set to the default path which is
// $APP_HOME/.pesakit.env
//
// In the fourth scenario, the config file path is set and home is set.
// If the file is not available, the function returns an error. If home is not available
// or not a directory, the function returns an error. If the file is available, it will
// be copied to the home directory of the application. and the new config file path is
// set to $APP_HOME/{config_file_name}.
func setConfig(cmd *cobra.Command, args []string, logger io.Writer,
	debugMode bool) (string, string, error) {

	defer func(debug bool) {
		if debug {
			_, _ = fmt.Fprintf(logger, "Debug mode is on: args %v", args)
		}
	}(debugMode)
	var (
		configFile          = ""
		homeDir             = ""
		configFileFlagGiven = cmd.Flags().Changed(flagConfigFile)
		homeDirFlagGiven    = cmd.Flags().Changed(flagHomeDirectory)
		err                 error
	)

	// possible scenarios:
	// 1. config file is given, home dir is not given
	// 2. config file is not given, home dir is given
	// 3. both config file and home dir are given
	// 4. neither config file nor home dir are given
	var (
		configGivenHomeDirNotGiven       = configFileFlagGiven && !homeDirFlagGiven
		homeDirGivenConfigFileNotGiven   = !configFileFlagGiven && homeDirFlagGiven
		bothConfigFileAndHomeDirGiven    = configFileFlagGiven && homeDirFlagGiven
		neitherConfigFileNorHomeDirGiven = !configFileFlagGiven && !homeDirFlagGiven
	)

	// switch on the possible scenarios
	switch {
	case configGivenHomeDirNotGiven:
		return "", "", err

	case homeDirGivenConfigFileNotGiven:
		return "", "", err

	case bothConfigFileAndHomeDirGiven:
		return bothHomeAndConfig(cmd)

	case neitherConfigFileNorHomeDirGiven:
		return neitherHomeNorConfig(homeDir)

	default:
		return "", "", fmt.Errorf("unexpected error")
	}

}

func bothHomeAndConfig(cmd *cobra.Command) (string, string, error) {
	// scenario 3
	configFile, err := cmd.Flags().GetString(flagConfigFile)
	if err != nil {
		err1 := fmt.Errorf("could not read specified config file: %w", err)

		return "", "", err1
	}
	// check if the specified config file exists
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		err1 := fmt.Errorf("specified config file does not exist: %w", err)

		return "", "", err1
	}

	homeDir, err := cmd.Flags().GetString(flagHomeDirectory)
	if err != nil {
		err1 := fmt.Errorf("could not read specified home directory: %w", err)

		return "", "", err1
	}
	// check if specified home directory is a directory and check if available
	if !home.IsDirExist(homeDir) {
		err1 := fmt.Errorf("specified home directory is not a directory/does not exist: %s", homeDir)

		return "", "", err1
	}
	// copy the config file to the home directory
	dstFile, err := copyToDir(configFile, homeDir)
	if err != nil {
		err2 := fmt.Errorf("could not copy config file to home directory: %w", err)

		return "", "", err2
	}

	return homeDir, dstFile, nil
}

func neitherHomeNorConfig(homeDir string) (string, string, error) {
	homeDir, err := home.At("")
	if err != nil {
		err1 := fmt.Errorf("could not create home directory: %w", err)

		return "", "", err1
	}
	confFilePath := filepath.Join(homeDir, ".pesakit.env")
	// check if the file named .pesakit.env exists in the homeDir if not
	// create it
	// if it exists, return it as the config file path
	info, err := os.Stat(confFilePath)
	if os.IsNotExist(err) {
		// create the file
		err1 := os.MkdirAll(homeDir, 0755)
		if err1 != nil {
			err2 := fmt.Errorf("could not create home directory: %w", err1)

			return "", "", err2
		}
		err1 = ioutil.WriteFile(confFilePath, []byte(""), 0644)
		if err1 != nil {
			err2 := fmt.Errorf("could not create config file: %w", err1)

			return "", "", err2
		}
		return homeDir, confFilePath, nil
	} else if info.IsDir() {
		err1 := fmt.Errorf("config file is a directory: %s", confFilePath)

		return "", "", err1
	} else if info.Mode().IsRegular() {
		return homeDir, confFilePath, nil
	}

	return "", "", err
}
