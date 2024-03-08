package algo

import "strings"

type Algorithm int64

const (
	None Algorithm = iota
	HS256
	HS384
	HS512
	RS256
	RS384
	RS512
	ES256
	ES384
	ES512
	PS256
	PS384
	PS512
)

func Parse(s string) Algorithm {
	switch strings.ToLower(s) {
	case "hs256":
		return HS256
	case "hs384":
		return HS384
	case "hs512":
		return HS512
	case "rs256":
		return RS256
	case "rs384":
		return RS384
	case "rs512":
		return RS512
	case "es256":
		return ES256
	case "es384":
		return ES384
	case "es512":
		return ES512
	case "ps256":
		return PS256
	case "ps384":
		return PS384
	case "ps512":
		return PS512
	case "none":
		fallthrough
	default:
		return None
	}
}
func (a Algorithm) MinKeyLength() int {
	switch a {
	case None:
		return -1
	case HS256:
		return 256
	case HS384:
		return 384
	case HS512:
		return 512
	case RS256:
		return 2048
	case RS384:
		return 2048
	case RS512:
		return 2048
	case ES256:
		return 256
	case ES384:
		return 384
	case ES512:
		return 512
	case PS256:
		return 2048
	case PS384:
		return 2048
	case PS512:
		return 2048
	default:
		return -1
	}
}

func (a Algorithm) String() string {
	switch a {
	case None:
		return "none"
	case HS256:
		return "hs256"
	case HS384:
		return "hs384"
	case HS512:
		return "hs512"
	case RS256:
		return "rs256"
	case RS384:
		return "rs384"
	case RS512:
		return "rs512"
	case ES256:
		return "es256"
	case ES384:
		return "es384"
	case ES512:
		return "es512"
	case PS256:
		return "ps256"
	case PS384:
		return "ps384"
	case PS512:
		return "ps512"
	default:
		return ""
	}
}
