package cmd

import (
	"fmt"
	"log"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// FIXME: replace 'APPNAME' with your application name
const appname = "APPNAME"
const description = "FIXME: description"
const longDescription = "FIXME: long description"

const defaultUser = "foobar"

// defaultCfgFile is the default config file name without extention
const defaultCfgFile = "." + appname
const defaultCfgFileType = "yaml"
const envPrefix = appname // FIXME: if appname contains dashes (`-`), replace with underscores (`_`)

// nolint: gochecknoglobals
var (
	version = "dev"
	commit  = ""
	date    = ""
	builtBy = ""
)

// cfgFile contains the config file path if set by a CLI flag
var cfgFile string

// printVersion is true when version flag is set
var printVersion bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   appname,
	Short: description,
	Long:  longDescription,
	Run: func(cmd *cobra.Command, args []string) {
		if printVersion {
			fmt.Println(description)
			fmt.Println(buildVersion(version, commit, date))
			return
		}

		log.SetFlags(log.Ltime)

		var cfg config
		if err := viper.Unmarshal(&cfg); err != nil {
			log.Fatalf("Error: Unable to read config: %s\n", err)
		}

		// FIXME: check config / params
		/* if cfg.Username == "" || cfg.Password == "" {
			log.Fatalln("Please set your username and password in the config file (" + defaultCfgFile + "." + defaultCfgFileType + ").")
		} */

		fmt.Println("This CLI tool is just an example.")
		fmt.Printf("User: %s\n", cfg.Username)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Define flags
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is "+defaultCfgFile+"."+defaultCfgFileType+" in program dir, CWD or $HOME)")
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "username", "u", "", "username (default is "+defaultUser+")")
	rootCmd.Flags().BoolVarP(&printVersion, "version", "v", false, "show version and exit")

	// Set default config values
	viper.SetDefault("username", defaultUser)

	// Bind flags to config values
	viper.BindPFlag("username", rootCmd.PersistentFlags().Lookup("username"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if printVersion {
		// skip reading config when printVersionis set
		return
	}

	if cfgFile != "" {
		// Use config file from the flag
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Default config name
		viper.SetConfigName(defaultCfgFile)
		// Default config type
		viper.SetConfigType(defaultCfgFileType)

		if ex, err := os.Executable(); err == nil {
			// Search for config in directory of executable
			viper.AddConfigPath(ex)
		}

		// Search for config in CWD
		viper.AddConfigPath(".")
		// Search for config in home dir
		viper.AddConfigPath(home)
	}

	// Read in environment variables that match
	viper.AutomaticEnv()
	viper.SetEnvPrefix(envPrefix)

	// If a config file is found, read it in
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func buildVersion(version, commit, date string) string {
	var result = fmt.Sprintf("version: %s", version)
	if commit != "" {
		result = fmt.Sprintf("%s\ncommit: %s", result, commit)
	}
	if date != "" {
		result = fmt.Sprintf("%s\nbuilt at: %s", result, date)
	}
	if builtBy != "" {
		result = fmt.Sprintf("%s\nbuilt by: %s", result, builtBy)
	}
	return result
}
