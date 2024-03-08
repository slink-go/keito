package keys

import (
	"fmt"
	"keito/lib/algo"
	"keito/lib/util"
)

const defaultKeyLength = 16

func Generate(algoStr string, length int) (string, error) {

	var a algo.Algorithm
	var err error
	a, err = getAlgo(algoStr)
	if err != nil {
		return "", err
	}

	if a == algo.None {
		if length <= 0 {
			return util.RandomString(defaultKeyLength)
		} else {
			return util.RandomString(length)
		}
	}
	if length <= 0 {
		return util.RandomString(a.MinKeyLength() / 8)
	}
	if a.MinKeyLength()/8 > length {
		return "", fmt.Errorf(
			"requested key length (%d) is too small for algorithm %s (%d)",
			length, a, a.MinKeyLength(),
		)
	}
	return util.RandomString(length)
}

func getAlgo(algoStr string) (algo.Algorithm, error) {
	if algoStr == "" {
		return algo.None, nil
	}
	a := algo.Parse(algoStr)
	if a == algo.None {
		return algo.None, fmt.Errorf("unsupported algorithm: %s", algoStr)
	}
	return a, nil
}
