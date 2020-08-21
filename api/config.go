package api

import (
	log "github.com/magna5/go-logger"
	"github.com/magna5/go-logger/shims/zerolog"

	sentry "github.com/getsentry/sentry-go"
	flags "github.com/jessevdk/go-flags"
	"github.com/joeshaw/envdecode"
	"github.com/magna5/godotenv"
)

// Cfg configuration structure
type Cfg struct {
	Port                 int    `env:"PORT,default=8080" short:"p" long:"port" description:"HTTP Port"`
	JwksCertRenewMinutes int    `env:"JWKS_RENEW_MINUTES,default=60" description:"Number of minutes to wait before renewing JWKS certificates"`
	JWTIssuer            string `env:"JWT_ISSUER" description:"The URL to the JWT issuing server"`
}

// Config is the current application configuration
var Config = &Cfg{}

func Configure(args []string) *Cfg {
	log.RootLogger = zerolog.New(nil)

	err := sentry.Init(sentry.ClientOptions{AttachStacktrace: true})
	if err != nil {
		log.Fatal("Error Initializing sentry: ", "error", err.Error())
	}

	err = godotenv.LoadOpt()
	if err != nil {
		panic(err)
	}
	err = envdecode.Decode(Config)
	if err != nil {
		panic(err)
	}
	_, err = flags.Parse(Config)
	if err != nil {
		panic(err)
	}
	return Config
}
