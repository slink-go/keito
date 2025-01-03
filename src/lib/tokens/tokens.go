package tokens

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/xhit/go-str2duration/v2"
	"keito/lib/algo"
	"keito/lib/keys"
	"keito/lib/util"
	"math/big"
	"strconv"
	"strings"
	"time"
)

func Generate(algoStr, issuer, subject, durationStr, claimsStr string, secret []byte, singleUse bool) (string, error) {

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

	if secret == nil {
		secret = util.ReadKeyConfig()
	}

	if secret == nil {
		secret, err = keys.Generate(algoStr, -1)
		if err != nil {
			return "", err
		}
	}

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func Parse(token string, key []byte) (map[string]interface{}, bool, error) {
	tk, _, err := new(jwt.Parser).ParseUnverified(token, jwt.MapClaims{})
	if err != nil {
		fmt.Println(err)
		return nil, false, err
	}
	claims, ok := tk.Claims.(jwt.MapClaims)
	if ok {
		for k, v := range claims {
			if k == "iat" || k == "exp" {
				claims[k] = convertTs(fmt.Sprintf("%v", v))
			}
		}
	} else {
		return nil, false, fmt.Errorf("could not parse claims")
	}
	claims["algo"] = tk.Method.Alg()

	verified := false
	if key == nil {
		key = util.ReadKeyConfig()
	}
	if key != nil {
		_, err = jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			return key, nil
		})
		if err == nil {
			verified = true
		}
	}

	return claims, verified, nil
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
	default:
		return nil, fmt.Errorf("unsupported algorithm: %s", a)
	}
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

func convertTs(input string) time.Time {
	v, err := scientificNotationToInt(input)
	if err != nil {
		panic(err)
	}
	return time.Unix(v, 0)
}
func scientificNotationToInt(scientificNotation string) (int64, error) {
	flt, _, err := big.ParseFloat(scientificNotation, 10, 0, big.ToNearestEven)
	if err != nil {
		return 0, err
	}
	fltVal := fmt.Sprintf("%.0f", flt)
	intVal, err := strconv.ParseInt(fltVal, 10, 64)
	if err != nil {
		return 0, err
	}
	return intVal, nil
}
