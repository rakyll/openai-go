// an interactive CLI for the OpenAI API
package main

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	apiKey   string
	logLevel string

	rootCmd = &cobra.Command{
		Use:   "openai",
		Short: "an interactive CLI for the OpenAI API",
		Long:  `an interactive CLI for the OpenAI API`,
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
)

func setLogLevel(level string) {
	if l, err := zerolog.ParseLevel(level); err == nil {
		zerolog.SetGlobalLevel(l)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
}

func init() {
	viper.AutomaticEnv()
	flags := rootCmd.Flags()

	flags.StringVarP(&apiKey, "openai_api_key", "k", "", "API Key provided by OpenAI")
	_ = viper.BindPFlag("openai_api_key", flags.Lookup("openai_api_key"))
	_ = rootCmd.MarkFlagRequired("openai_api_key")

	flags.StringVarP(&logLevel, "log_level", "l", "info", "log level")
	_ = viper.BindPFlag("log_level", flags.Lookup("log_level"))
	setLogLevel(logLevel)

	// change completion command name to avoid conflict with completion subcommand
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	// a custom yet identical shell completion command
	var completionCmd = &cobra.Command{
		Use:                   "shell_completion [bash|zsh|fish|powershell]",
		Short:                 "Generate completion script",
		DisableFlagsInUseLine: true,
		ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
		Args:                  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		Run: func(cmd *cobra.Command, args []string) {
			switch args[0] {
			case "bash":
				_ = cmd.Root().GenBashCompletion(os.Stdout)
			case "zsh":
				_ = cmd.Root().GenZshCompletion(os.Stdout)
			case "fish":
				_ = cmd.Root().GenFishCompletion(os.Stdout, true)
			case "powershell":
				_ = cmd.Root().GenPowerShellCompletionWithDesc(os.Stdout)
			}
		},
	}
	rootCmd.AddCommand(completionCmd)

}

func main() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
