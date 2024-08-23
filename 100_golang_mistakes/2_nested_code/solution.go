package main

import (
	"errors"
)

func join2(s1, s2 string, max int) (string, error) {
	if s1 == "" {
		return "", errors.New("s1 is empty")
	}

	if s2 == "" {
		return "", errors.New("s2 is empty")
	}

	concat, err := concatenate2(s1, s2)
	if err != nil {
		return "", err
	}

	if len(concat) > max {
		return concat[:max], nil
	}

	return concat, nil
}

func concatenate2(s1 string, s2 string) (string, error) {
	return s1 + s2, nil
}
