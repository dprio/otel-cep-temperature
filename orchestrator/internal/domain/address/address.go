package address

import (
	"errors"
	"strconv"
)

var (
	ErrInvalidZipCode    = errors.New("invalid zip code. Must have 8 digits")
	ErrNotNumericZipCode = errors.New("invalid zip code. Must be numeric")
)

type (
	Address struct {
		ZipCode ZipCode
		City    string
	}

	ZipCode string
)

func New(zipCode, city string) (*Address, error) {
	zc, err := NewZipCode(zipCode)
	if err != nil {
		return nil, err
	}

	return &Address{
		ZipCode: zc,
		City:    city,
	}, nil
}

func NewZipCode(aZipCode string) (ZipCode, error) {
	if len(aZipCode) != 8 {
		return "", ErrInvalidZipCode
	}

	if _, err := strconv.Atoi(aZipCode); err != nil {
		return "", ErrNotNumericZipCode
	}

	return ZipCode(aZipCode), nil
}

func (z ZipCode) Value() string {
	return string(z)
}
