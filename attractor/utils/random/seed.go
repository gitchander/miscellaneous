package random

import (
	"fmt"
	"strconv"
)

func ParseSeedOrMake(s string) (int64, error) {
	var seed int64
	if s == "" {
		seed = NewRandNow().Int63()
		fmt.Println("seed", seed)
	} else {
		i, err := parseInt64(s)
		if err != nil {
			return 0, fmt.Errorf("invalid seed value %q", s)
		}
		seed = i
	}
	return seed, nil
}

func parseInt64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}
