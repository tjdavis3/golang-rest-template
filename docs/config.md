# Configuration

Configuration is typically done with environment variables or command-line flags.  Not all settings have command-line flags or default values.  If there's no value set the feature is disabled.  

The order of preference is:

   1. command-line flag
   1. environment variable
   1. variable defined in `.env` file
   1. default

## Configuration Variables and Flags

Variable | Flag | Default | Description |
-------  | ---- | ------- | -----------
SENTRY_DSN |  |  | DSN for sentry crash detection
SENTRY_ENVIRONMENT |  |  | Environment to report to sentry
PORT | -p,<br/> --port  | 8080  | HTTP Port on which to listen
JWKS_RENEW_MINUTES | | 60 | Number of minutes between JWKS certificate renewals
JWT_ISSUER | | | The URL to the JWT issuing server
