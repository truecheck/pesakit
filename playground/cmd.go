package main

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"strings"
)

const (
	envPrefix = "PESAKIT"
	envPort   = "PORT"
	envHost   = "HOST"
	flagPort  = "port"
	flagHost  = "host"
	defPort   = "8080"
	defHost   = "localhost"
)

func init() {
	/* Todo
	1. init viper
	2. set config file
	3. add config paths
	4. set config name
	*/

	configPaths := func(paths ...string) {
		for _, path := range paths {
			viper.AddConfigPath(path)
		}
	}
	cwd, err := os.Getwd()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "could not get current working directory: %s\n", err)
	}
	configPaths("/etc/pesakit", "./playground", ".", "..", "$HOME/.pesakit", cwd)

	viper.SetConfigName("pesakit.env")
	viper.SetConfigType("env")

	//setDefaults := func(defaults map[string]interface{}) error {
	//	for key, value := range defaults {
	//		if viper.Get(key) == nil {
	//			viper.Set(key, value)
	//		}
	//	}
	//
	//	return nil
	//}
	//
	//defaultMap := map[string]interface{}{
	//	flagPort: defPort,
	//	flagHost: defHost,
	//}
	//
	//err = setDefaults(defaultMap)
	//if err != nil {
	//	fmt.Fprintln(os.Stderr, "could not set defaults")
	//
	//	return
	//}

	viper.AutomaticEnv()

	replacer := strings.NewReplacer("-", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.SetEnvPrefix(envPrefix)

	bindEnvs := func(keys ...string) error {
		for _, key := range keys {
			if err := viper.BindEnv(key); err != nil {
				return err
			}
		}

		return nil
	}

	err = bindEnvs(flagHost, flagPort)
	if err != nil {
		return
	}

	if err := viper.ReadInConfig(); err != nil {
		var _t0 viper.ConfigFileNotFoundError
		if ok := errors.Is(err, _t0); ok {
			_, _ = fmt.Fprintln(os.Stderr, "no config file found")
		} else {
			// Config file was found but another error was produced
			_, _ = fmt.Fprintln(os.Stderr, "read config error", err)
		}
		_, _ = fmt.Fprintf(os.Stderr, "using default config: %s\n", viper.ConfigFileUsed())
	}

	fmt.Fprintln(os.Stdout, viper.GetString(flagPort))
	fmt.Fprintln(os.Stdout, viper.GetString(flagHost))
}

func rootCommand() *cobra.Command {
	//configPaths := func(paths ...string) {
	//	for _, path := range paths {
	//		viper.AddConfigPath(path)
	//	}
	//}
	//cwd, err := os.Getwd()
	//if err != nil {
	//	_, _ = fmt.Fprintf(os.Stderr, "could not get current working directory: %s\n", err)
	//}
	//configPaths("/etc/pesakit", "./playground", ".", "..", "$HOME/.pesakit", cwd)
	//
	//viper.SetConfigName("pesakit.env")
	//viper.SetConfigType("env")
	rootCmd := &cobra.Command{
		Use:   "pesakit",
		Short: "Pesakit is a simple web server",
		Long:  `Pesakit is a simple web server`,
		Run: func(cmd *cobra.Command, args []string) {
			// get port and host from flags
			port := viper.GetString(flagPort)
			host := viper.GetString(flagHost)
			_, _ = fmt.Fprintf(os.Stdout, "serving on %s:%s\n", host, port)
		},
	}

	rootCmd.PersistentFlags().StringP(flagPort, "p", defPort, "port to listen on")
	rootCmd.PersistentFlags().StringP(flagHost, "H", defHost, "host to listen on")
	viper.BindPFlag(flagHost, rootCmd.PersistentFlags().Lookup(flagHost)) //nolint:errcheck
	viper.BindPFlag(flagPort, rootCmd.PersistentFlags().Lookup(flagPort)) //nolint:errcheck

	return rootCmd
}

func main() {
	app := rootCommand()
	if err := app.Execute(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
		return
	}
}
