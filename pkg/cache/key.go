package cache

import (
	"errors"
	"strings"
)

func BuildCacheKey(keyPrefix string, key string) (cacheKey string, err error) {
	if key == "" {
		return "", errors.New("[cache] key should not be empty")
	}

	cacheKey, err = strings.Join([]string{keyPrefix, key}, ":"), nil
	return
}