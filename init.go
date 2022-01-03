package pesakit

import (
	"errors"
	"fmt"
	"github.com/pesakit/pesakit/env"
	"github.com/pesakit/pesakit/flags"
	"github.com/pesakit/pesakit/home"
	"github.com/spf13/cobra"
	"github.com/techcraftlabs/airtel"
	"github.com/techcraftlabs/mpesa"
	"github.com/techcraftlabs/tigopesa"
	"io"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func (app *App) persistentPreRun(cmd *cobra.Command, args []string) error {
	appConfig, err := flags.LoadAppConfig(cmd)
	if err != nil {
		return err
	}
	logger, debugMode := app.getWriter(), appConfig.Debug
	app.setDebugMode(debugMode)
	appHomeDir, appConfigFile, err := initConfig(cmd, args, logger, debugMode)
	if err != nil {
		_, _ = fmt.Fprintf(logger, "error: %v\n", err)

		return err
	}

	_, _ = fmt.Fprintf(logger, "PERSISTENT PRE RUN: app home dir: %s, config file %s\n", appHomeDir, appConfigFile)

	app.setHomeDir(appHomeDir)
	err = env.LoadConfigFrom(appConfigFile)
	if err != nil {
		_, _ = fmt.Fprintf(logger, "error: %v\n", err)

		return err
	}
	if err = app.loadConfigAndSetClients(cmd, logger, debugMode); err != nil {
		_, _ = fmt.Fprintf(logger, "error: %v\n", err)

		return err
	}
	_, _ = fmt.Fprintf(logger, "app home is %s and config loaded from %s\n", appHomeDir, appConfigFile)
	//
	return nil
}

func (app *App) loadConfigAndSetClients(cmd *cobra.Command, logger io.Writer, debugMode bool) error {
	// loads all configurations there is
	configMpesa, err := loadMpesaConfig(cmd)
	if err != nil {
		return err
	}

	configAirtel, err := loadAirtelConfig(cmd)
	if err != nil {
		return err
	}

	configTigo, err := loadTigoConfig(cmd)
	if err != nil {
		return err
	}
	clientMpesa := mpesa.NewClient(configMpesa, mpesa.WithLogger(logger), mpesa.WithDebugMode(debugMode))
	clientTigo := tigopesa.NewClient(configTigo, tigopesa.WithDebugMode(debugMode), tigopesa.WithLogger(logger))
	clientAirtel := airtel.NewClient(configAirtel, nil, true)
	app.setMpesaClient(clientMpesa)
	app.setTigoClient(clientTigo)
	app.setAirtelClient(clientAirtel)

	return nil
}

const (
	defaultAppHomeDirectory = ""
	defaultConfigFilename   = "pesakit.env"
	defaultConfigFileData   = ""
	defaultConfigFilePerm   = fs.ModePerm
)

//
func (app *App) persistentPreRunE(cmd *cobra.Command, args []string) error {
	return nil
}

// initConfig returns a home directory of the application, and a config file path.
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
func initConfig(cmd *cobra.Command, args []string, logger io.Writer,
	debugMode bool) (string, string, error) {

	_, _ = fmt.Fprintf(logger, "here at the initConfig %v\n", args)

	defer func(debug bool) {
		if debug {
			_, _ = fmt.Fprintf(logger, "debug mode is on: args %v", args)
		}
	}(debugMode)

	rootCommand := getParentCommand(cmd)

	var (
		configFileFlagGiven = rootCommand.PersistentFlags().Changed(flags.ConfigName)
		homeDirFlagGiven    = rootCommand.PersistentFlags().Changed(flags.HomeDirectoryName)
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

		return onlyConfigFileSpecified(rootCommand)

	case homeDirGivenConfigFileNotGiven:

		return onlyHomeGiven(rootCommand)

	case bothConfigFileAndHomeDirGiven:

		return bothHomeAndConfig(rootCommand)

	case neitherConfigFileNorHomeDirGiven:

		return neitherHomeNorConfig()

	default:
		_, _ = fmt.Fprintf(logger, "default case\n")
		return "", "", fmt.Errorf("unexpected error")
	}

}

func onlyConfigFileSpecified(cmd *cobra.Command) (string, string, error) {
	// retrieve specified config file path
	// scenario 3
	configFile, err := cmd.Flags().GetString(flags.ConfigName)
	if err != nil {
		err1 := fmt.Errorf("could not read specified config file: %w", err)

		return "", "", err1
	}

	fileExists := home.IsFileExist(configFile)
	if !fileExists {
		err1 := errors.New(fmt.Sprintf("specified config file does not exist: %s", configFile))

		return "", "", err1
	}
	homeDir, err := home.At(defaultAppHomeDirectory)
	if err != nil {
		err1 := fmt.Errorf("could not create home directory: %w", err)

		return "", "", err1
	}
	dir, err := copyToDir(configFile, homeDir)
	if err != nil {
		return "", "", err
	}
	return homeDir, dir, err
}

func onlyHomeGiven(cmd *cobra.Command) (string, string, error) {
	homeDir, err := cmd.Flags().GetString(flags.HomeDirectoryName)
	if err != nil {
		err1 := fmt.Errorf("could not read specified home directory: %w", err)

		return "", "", err1
	}
	// check if specified home directory is a directory and check if available
	if !home.IsDirExist(homeDir) {
		err1 := fmt.Errorf("specified home directory is not a directory/does not exist: %s", homeDir)

		return "", "", err1
	}

	return createsDefaultConfigFile(homeDir)
}

func createsDefaultConfigFile(homeDir string) (string, string, error) {
	// create a file named .pesakit.env in the homeDir
	confFilePath := filepath.Join(homeDir, defaultConfigFilename)
	// check if the file named .pesakit.env exists in the homeDir if not
	// create it
	// if it exists, return it as the config file path

	log.Printf("confFilePath: %s", confFilePath)
	info, err := os.Stat(confFilePath)
	if os.IsNotExist(err) {
		log.Printf("file does not exist, we are going to create it")
		// create the file
		err1 := ioutil.WriteFile(confFilePath, []byte(defaultConfigFileData), defaultConfigFilePerm)
		if err1 != nil {
			err2 := fmt.Errorf("could not create config file: %w", err1)

			return "", "", err2
		}

		log.Printf("created config file: %s", confFilePath)
		return homeDir, confFilePath, nil
	}
	// if the file exists check if its regular file. If its regular return it.
	// if it's not regular, return an error
	// also check is it's a directory if it is creates a regular file with the
	// same name
	if !info.Mode().IsRegular() {
		err1 := fmt.Errorf("config file is not a regular file: %s", confFilePath)

		return "", "", err1
	}

	if info.IsDir() {
		// create the file
		err1 := ioutil.WriteFile(confFilePath, []byte(defaultConfigFileData), defaultConfigFilePerm)
		if err1 != nil {
			err2 := fmt.Errorf("could not create config file: %w", err1)

			return "", "", err2
		}
		return homeDir, confFilePath, nil
	}

	return homeDir, confFilePath, err
}

func bothHomeAndConfig(cmd *cobra.Command) (string, string, error) {
	// scenario 3
	configFile, err := cmd.Flags().GetString(flags.ConfigName)
	if err != nil {
		err1 := fmt.Errorf("could not read specified config file: %w", err)

		return "", "", err1
	}
	// check if the specified config file exists
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		err1 := fmt.Errorf("specified config file does not exist: %w", err)

		return "", "", err1
	}

	homeDir, err := cmd.Flags().GetString(flags.HomeDirectoryName)
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

func neitherHomeNorConfig() (string, string, error) {
	log.Printf("Neither home nor config file specified. Using default values\n")
	homeDir, err := home.At(defaultAppHomeDirectory)
	if err != nil {
		err1 := fmt.Errorf("could not create home directory: %w", err)

		return "", "", err1
	}

	return createsDefaultConfigFile(homeDir)

}
