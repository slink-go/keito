package tokens

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/xhit/go-str2duration/v2"
	"keito/lib/algo"
	"keito/lib/keys"
	"strings"
	"time"
)

func Generate(algoStr, secret, issuer, subject, durationStr, claimsStr string, singleUse bool) (string, error) {

	sign, err := getAlgo(algoStr)
	if err != nil {
		return "", err
	}

	duration, err := getDuration(durationStr)
	if err != nil {
		return "", err
	}

	cl, err := getClaims(claimsStr)
	if err != nil {
		return "", err
	}

	// https://auth0.com/docs/secure/tokens/json-web-tokens/json-web-token-claims
	token := jwt.New(sign) // see also jwt.SigningMethodEdDSA
	claims := token.Claims.(jwt.MapClaims)
	claims["iat"] = time.Now().UnixMilli() / 1000
	claims["exp"] = time.Now().Add(duration).UnixMilli() / 1000
	if subject != "" {
		claims["sub"] = subject
	}
	if issuer != "" {
		claims["iss"] = issuer
	}
	if singleUse {
		id, err := uuid.NewUUID()
		if err != nil {
			return "", err
		}
		claims["jti"] = id.String()
	}
	for k, v := range cl {
		claims[k] = v
	}

	if secret == "" {
		secret, err = keys.Generate(algoStr, -1)
		if err != nil {
			return "", err
		}
	}
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func getAlgo(input string) (jwt.SigningMethod, error) {
	a := algo.Parse(input)
	if a == algo.None {
		return nil, fmt.Errorf("unknown algorithm: %s", input)
	}
	switch a {
	case algo.HS256:
		return jwt.SigningMethodHS256, nil
	case algo.HS384:
		return jwt.SigningMethodHS384, nil
	case algo.HS512:
		return jwt.SigningMethodHS512, nil
	}
	return nil, fmt.Errorf("unsupported algorithm: %s", a)
}
func getDuration(input string) (time.Duration, error) {
	return str2duration.ParseDuration(input) // "ns", "us" (or "µs"), "ms", "s", "m", "h", "d", "w"
}
func getClaims(input string) (map[string]string, error) {
	if input == "" {
		return make(map[string]string, 0), nil
	}
	var result map[string]string = make(map[string]string, 0)
	for _, pair := range strings.Split(input, ",") {
		a, b, ok := strings.Cut(pair, "=")
		if !ok {
			continue
		}
		a = strings.TrimSpace(a)
		b = strings.TrimSpace(b)
		result[a] = b
	}
	return result, nil
}
