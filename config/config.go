//go:generate go run ../cmd/configdoc/main.go -o ../docs/content/A_config.md
package config

import (
	"fmt"
	"os"

	log "github.com/ringsq/go-logger"
	"github.com/ringsq/go-logger/shims/zerolog"

	sentry "github.com/getsentry/sentry-go"
	flags "github.com/jessevdk/go-flags"

	// Autoloads environment from .env
	_ "github.com/joho/godotenv/autoload"
)

const progDesc = `
Serves a REST API for ...
`

// Cfg configuration structure
type Cfg struct {
	Port                 int    `env:"PORT" default:"8080" short:"p" long:"port" description:"HTTP Port"`
	JwksCertRenewMinutes int    `env:"JWKS_RENEW_MINUTES" default:"60" description:"Number of minutes to wait before renewing JWKS certificates"`
	JWTIssuer            string `env:"JWT_ISSUER" description:"The URL to the JWT issuing server"`
}

// Config is the current application configuration
var Config = &Cfg{}

// Configure creates the configuration for the application
func Configure(args []string) *Cfg {
	log.RootLogger = zerolog.New(nil)

	err := sentry.Init(sentry.ClientOptions{AttachStacktrace: true})
	if err != nil {
		log.Fatal("Error Initializing sentry: ", "error", err.Error())
	}

	parser := flags.NewParser(Config, flags.Default)
	// Uncomment the following line to add a custom usage statement
	// parser.Usage = "[OPTIONS]"
	parser.ShortDescription = "A REST API generated from the OpenAPI spec"
	parser.LongDescription = progDesc
	_, err = parser.Parse()
	if err != nil {
		if flgErr, ok := err.(*flags.Error); ok {
			if flgErr.Type == flags.ErrHelp {
				os.Exit(0)
			}
		}
		fmt.Println(err.Error())
		os.Exit(2)
	}
	return Config
}
