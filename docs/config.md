# Configuration

Configuration is typically done with environment variables or command-line flags.  Not all settings have command-line flags or default values.  If there's no value set the feature is disabled.  

The order of preference is:

   1. command-line flag
   1. environment variable
   1. variable defined in `.env` file
   1. default

## Configuration Variables and Flags

| Variable           | Flag            | Type | Default | Description                                                         |
| ------------------ | --------------- | ---- | ------- | ------------------------------------------------------------------- |
| SENTRY_DSN | | string |  | DSN for sentry crash detection |
| SENTRY_ENVIRONMENT | | string |  | Environment to report to sentry |
| PORT | -p<br/>--port | int | 8080 | HTTP Port |
| JWKS_RENEW_MINUTES |  | int | 60 | Number of minutes to wait before renewing JWKS certificates |
| JWT_ISSUER |  | string |  | The URL to the JWT issuing server |
