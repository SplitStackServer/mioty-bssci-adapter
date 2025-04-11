package cmd

import (
	"bytes"
	"mioty-bssci-adapter/internal/config"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFiles *[]string // config file
var version string

var rootCmd = &cobra.Command{
	Use:   "mioty-bssci-adapter",
	Short: "abstracts the mioty bssci protocol into Protobuf or JSON over MQTT",
	Long:  `mioty BSSCI Adapter abstracts the mioty bssci protocol into Protobuf or JSON over MQTT`,
	RunE:  run,
}

// Execute the root command.
func Execute(v string) {
	version = v
	if err := rootCmd.Execute(); err != nil {
		log.Fatal().Err(err).Msg("failed to execute")
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	cfgFiles = rootCmd.PersistentFlags().StringSliceP("config", "c", []string{}, "path to configuration file (optional)")
	rootCmd.PersistentFlags().Int("log-level", 1, "debug=0, info=1 (default), warn=2, error=3, fatal=4, panic=5, disabled=7")

	viper.BindPFlag("general.log_level", rootCmd.PersistentFlags().Lookup("log-level"))

	// default values

	// logging
	viper.SetDefault("general.log_level", 1)
	viper.SetDefault("general.log_to_syslog", false)

	// bssci_v1 backend
	viper.SetDefault("backend.type", "bssci_v1")

	viper.SetDefault("backend.bssci_v1.bind", "0.0.0.0:5005")
	viper.SetDefault("backend.bssci_v1.stats_interval", time.Minute*5)
	viper.SetDefault("backend.bssci_v1.ping_interval", time.Second*30)
	viper.SetDefault("backend.bssci_v1.keep_alive_period", time.Minute*3)

	// mqtt_v3 integration
	viper.SetDefault("integration.marshaler", "protobuf")

	viper.SetDefault("integration.mqtt_v3.state_retained", true)
	viper.SetDefault("integration.mqtt_v3.keep_alive", 30*time.Second)
	viper.SetDefault("integration.mqtt_v3.max_reconnect_interval", time.Minute)
	viper.SetDefault("integration.mqtt_v3.max_token_wait", time.Minute)
	viper.SetDefault("integration.mqtt_v3.terminate_on_connect_error", false)

	viper.SetDefault("integration.mqtt_v3.auth.type", "generic")
	viper.SetDefault("integration.mqtt_v3.auth.generic.servers", []string{"tcp://127.0.0.1:1883"})
	viper.SetDefault("integration.mqtt_v3.auth.generic.clean_session", true)

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(configCmd)
}

func initConfig() {
	if cfgFiles != nil && len(*cfgFiles) != 0 {
		var filesMerged []byte
		for _, cfgFile := range *cfgFiles {
			cfgFileContent, err := os.ReadFile(cfgFile)
			if err != nil {
				log.Fatal().Err(err).Str("config", cfgFile).Msg("error loading config file")
			}
			filesMerged = bytes.Join([][]byte{
				filesMerged,
				cfgFileContent,
			}, []byte("\n"))
		}

		viper.SetConfigType("toml")
		if err := viper.ReadConfig(bytes.NewBuffer(filesMerged)); err != nil {
			log.Fatal().Err(err).Any("config", cfgFiles).Msg("error loading config file")

		}
	} else {
		viper.SetConfigName("mioty-bssci-adapter")
		viper.AddConfigPath(".")
		viper.AddConfigPath("$HOME/.config/mioty-bssci-adapter")
		viper.AddConfigPath("/etc/mioty-bssci-adapter/")
		if err := viper.ReadInConfig(); err != nil {
			switch err.(type) {
			case viper.ConfigFileNotFoundError:
			default:
				log.Fatal().Err(err).Msg("error loading config file")
			}
		}
	}

	for _, pair := range os.Environ() {
		d := strings.SplitN(pair, "=", 2)
		if strings.Contains(d[0], ".") {
			log.Warn().Msgf("Using dots in env variable is illegal and deprecated. Please use double underscore `__` for: %s", d[0])
			underscoreName := strings.ReplaceAll(d[0], ".", "__")
			// Set only when the underscore version doesn't already exist.
			if _, exists := os.LookupEnv(underscoreName); !exists {
				os.Setenv(underscoreName, d[1])
			}
		}
	}

	viperBindEnvs(config.C)

	if err := viper.Unmarshal(&config.C); err != nil {
		log.Fatal().Err(err).Msg("unmarshal config error")
	}
}

func viperBindEnvs(iface interface{}, parts ...string) {
	ifv := reflect.ValueOf(iface)
	ift := reflect.TypeOf(iface)
	for i := 0; i < ift.NumField(); i++ {
		v := ifv.Field(i)
		t := ift.Field(i)
		tv, ok := t.Tag.Lookup("mapstructure")
		if !ok {
			tv = strings.ToLower(t.Name)
		}
		if tv == "-" {
			continue
		}

		switch v.Kind() {
		case reflect.Struct:
			viperBindEnvs(v.Interface(), append(parts, tv)...)
		default:
			// Bash doesn't allow env variable names with a dot so
			// bind the double underscore version.
			keyDot := strings.Join(append(parts, tv), ".")
			keyUnderscore := strings.Join(append(parts, tv), "__")
			viper.BindEnv(keyDot, strings.ToUpper(keyUnderscore))
		}
	}
}
