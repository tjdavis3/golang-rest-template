package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/rs/zerolog/log"
)

// Jwks struct
type Jwks struct {
	Keys []JSONWebKeys `json:"keys"`
}

// JSONWebKeys struct
type JSONWebKeys struct {
	Kty string   `json:"kty"`
	Kid string   `json:"kid"`
	Use string   `json:"use"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	X5c []string `json:"x5c"`
}

var jwks = Jwks{}

// RefreshJWTKS refreshes jwt key set from the server
func RefreshJWTKS(cfg *Cfg) {
	refreshInterval := cfg.JwksCertRenewMinutes
	if refreshInterval != 0 {
		duration := time.Duration(refreshInterval) * time.Minute

		shutdown := make(chan os.Signal, 1)
		signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

		go func(d time.Duration) {
			ticker := time.NewTicker(d)
			defer ticker.Stop()
		refreshLoop:
			for {
				select {
				case <-ticker.C:
					FetchJWTKeySet(cfg)
				case <-shutdown:
					break refreshLoop
				}
			}
		}(duration)
	}
}

// FetchJWTKeySet stores keycloak jwt key set
func FetchJWTKeySet(cfg *Cfg) error {
	log.Info().Msg("Updating JWT Key set from the server...")
	resp, err := http.Get(cfg.JWTIssuer + "/protocol/openid-connect/certs")

	if err != nil {
		log.Error().Msg(err.Error())
		// ErrorCounter.Inc()
		return err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&jwks)

	if err != nil {
		log.Error().Msg(err.Error())
		// ErrorCounter.Inc()
		return err
	}

	log.Info().Msg("JWT Key set loaded successfully.")
	return nil
}

// JWTAuthentication middleware
func (s *server) JWTAuthentication(next http.Handler) http.Handler {
	fn := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		issuer := s.config.JWTIssuer

		if issuer != "" {
			jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
				ErrorHandler: JWTErrorHandler,
				Debug:        true,
				ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
					// Verify 'iss' claim
					checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(issuer, false)
					if !checkIss {
						return token, errors.New("invalid issuer")
					}

					publicKey, err := getPemCert(token)
					if err != nil {
						return nil, err
					}

					result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(publicKey))

					return result, nil
				},
				SigningMethod: jwt.SigningMethodRS256,
			})

			err := jwtMiddleware.CheckJWT(w, r)

			// If there was an error, do not continue.
			if err != nil {

				return
			}
		}

		next.ServeHTTP(w, r)
	})

	return fn
}

func getPemCert(token *jwt.Token) (string, error) {
	cert := ""
	for k := range jwks.Keys {
		if token.Header["kid"] == jwks.Keys[k].Kid {
			cert = "-----BEGIN CERTIFICATE-----\n" + jwks.Keys[k].X5c[0] + "\n-----END CERTIFICATE-----"
		}
	}

	if cert == "" {
		err := errors.New("unable to find appropriate key")
		return cert, err
	}

	return cert, nil
}
