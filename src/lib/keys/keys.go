package keys

import (
	"fmt"
	"keito/lib/algo"
	"keito/lib/util"
)

const defaultKeyLength = 16

func Generate(algoStr string, length int) ([]byte, error) {

	var a algo.Algorithm
	var err error
	a, err = getAlgo(algoStr)
	if err != nil {
		return nil, err
	}

	if a == algo.None {
		if length <= 0 {
			return util.RandomBytesPrintable(defaultKeyLength)
		} else {
			return util.RandomBytesPrintable(length)
		}
	}
	minChars := a.MinKeyLength() / 8
	if length <= 0 {
		return util.RandomBytesPrintable(minChars)
	}
	if minChars > length {
		return nil, fmt.Errorf(
			"requested key length (%d) is too small for algorithm %s (%d)",
			length, a, minChars,
		)
	}
	return util.RandomBytesPrintable(length)
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
