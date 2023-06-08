package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/boyane126/bcpt/internal/bcptctl/cmd/show"
	"github.com/boyane126/bcpt/internal/bcptctl/util/templates"
)

var cfgFile string

func NewDefaultBCPTCommand() *cobra.Command {
	return NewBCPTCommand(os.Stdin, os.Stdout, os.Stderr)
}

func NewBCPTCommand(in io.Reader, out, err io.Writer) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "bcptctl",
		Short: "bcptctl controls the bcpt platform",
		Long: templates.LongDesc(`
		bcptctl controls the bcpt platform, is the client side tool for bcpt platform.

		Find more information at:
			https://github.com/boyane126/bcpt/README.md`),
		Run: runHelp,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return initProfiling()
		},
		PersistentPostRunE: func(*cobra.Command, []string) error {
			return flushProfiling()
		},
	}

	flags := rootCmd.PersistentFlags()
	addProfilingFlags(flags)

	_ = viper.BindPFlags(flags)

	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
	rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "author name for copyright attribution")
	rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
	viper.SetDefault("license", "apache")

	groups := templates.CommandGroups{
		{
			Message: "Basic Commands:",
			Commands: []*cobra.Command{
				show.NewCmdShow(),
				// color.NewCmdColor(f, ioStreams),
			},
		},
		//{
		//	Message: "Identity and Access Management Commands:",
		//	Commands: []*cobra.Command{
		//		user.NewCmdUser(f, ioStreams),
		//		policy.NewCmdPolicy(f, ioStreams),
		//	},
		//},
	}
	groups.Add(rootCmd)

	filters := []string{"options"}
	templates.ActsAsRootCommand(rootCmd, filters, groups...)

	return rootCmd
}

func runHelp(cmd *cobra.Command, args []string) {
	_ = cmd.Help()
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".cobra")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
