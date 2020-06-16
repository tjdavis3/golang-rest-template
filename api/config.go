package api

import (
	"log"
	"os"

	"github.com/appsflyer/go-logger/shims/zerolog"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (

	// Log ("json" || "")
	ConfigLogFormat = "log"

	// http
	ConfigHTTPAddr = "http"

	// certificate
	ConfigTLSCert = "tls.crt"
	ConfigTLSKey  = "tls.key"
)

var Log = zerolog.New(nil)

func buildFlags() *pflag.FlagSet {
	flags := pflag.NewFlagSet(os.Args[0], pflag.ExitOnError)

	return flags
}

func Configure(args []string) *viper.Viper {
	log.SetOutput(Log)
	err := sentry.Init(sentry.ClientOptions{})
	if err != nil {
		log.Error("Error Initializing sentry: ", "error", err.Error())
	}

	v := viper.New()

	// Setup command line flags
	flags := buildFlags()
	if err := flags.Parse(args); err != nil {
		panic(err)
	}

	// Configuration from flags
	if err := v.BindPFlags(flags); err != nil {
		panic(err)
	}

	v.SetEnvPrefix(os.Args[0])

	// Configuration from env
	v.AutomaticEnv()
	return v
}
