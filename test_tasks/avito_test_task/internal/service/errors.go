package service

import "errors"

var (
	ErrUserExists    = errors.New("user already exists")
	ErrUserNotFound  = errors.New("user not found")
	ErrWrongPassword = errors.New("wrong password")
	ErrSignToken     = errors.New("can't sign token")
	ErrParseToken    = errors.New("can't parse token")

	ErrInvalidInputData = errors.New("invalid input data")
	ErrHouseExists      = errors.New("house already exists")
	ErrHouseNotFound    = errors.New("house not found")

	ErrFlatExists             = errors.New("flat already exists")
	ErrFlatNotFound           = errors.New("flat not found")
	ErrHouseFlatsNotFound     = errors.New("house flats not found")
	ErrFlatNotOnModeration    = errors.New("flat is not on moderation yet")
	ErrFlatOnModeration       = errors.New("flat is already on moderation")
	ErrFlatFinishedModeration = errors.New("flat is finished moderation")

	ErrHouseSubscriptionExists = errors.New("house subscription already exists")
)
