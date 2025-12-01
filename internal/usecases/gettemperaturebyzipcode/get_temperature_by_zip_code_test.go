package gettemperaturebyzipcode

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/dprio/otel-cep-temperature/internal/domain/address"
	"github.com/dprio/otel-cep-temperature/internal/domain/weather"
	"github.com/dprio/otel-cep-temperature/mocks/gettemperaturebyzipcode"
)

func TestExecute(t *testing.T) {

	t.Run("should execute successfully", func(t *testing.T) {
		//given
		mockViaCEPGateway := gettemperaturebyzipcode.NewMockViaCEPGateway(t)
		mockWeatherAPIGateway := gettemperaturebyzipcode.NewMockWeatherAPIGateway(t)

		target := New(mockViaCEPGateway, mockWeatherAPIGateway)

		zipCode := address.ZipCode("04303001")
		address := address.Address{
			ZipCode: zipCode,
			City:    "São Paulo",
		}

		weather := weather.Weather{
			Temperature: weather.Temperature{
				C: 100,
				F: 100,
				K: 100,
			},
		}

		expectedOut := &Output{
			Address: address,
			Weather: weather,
		}

		mockViaCEPGateway.EXPECT().
			GetAddressByZipCode(mock.Anything, mock.Anything).
			Return(&address, nil)

		mockWeatherAPIGateway.EXPECT().
			GetWeatherByCity(mock.Anything, mock.Anything).
			Return(&weather, nil)

		//when
		resp, err := target.Execute(context.Background(), zipCode.Value())

		//then
		assert.NoError(t, err)
		assert.Equal(t, expectedOut, resp)
	})

	t.Run("should execute with error when zip-code is not numeric", func(t *testing.T) {
		//given
		mockViaCEPGateway := gettemperaturebyzipcode.NewMockViaCEPGateway(t)
		mockWeatherAPIGateway := gettemperaturebyzipcode.NewMockWeatherAPIGateway(t)

		target := New(mockViaCEPGateway, mockWeatherAPIGateway)

		zipCode := address.ZipCode("0430300f")

		//when
		resp, err := target.Execute(context.Background(), zipCode.Value())

		//then
		assert.EqualError(t, err, address.ErrNotNumericZipCode.Error())
		assert.Nil(t, resp)

		mockViaCEPGateway.AssertNotCalled(t, "GetAddressByZipCode")
		mockWeatherAPIGateway.AssertNotCalled(t, "GetWeatherByCity")
	})

	t.Run("should execute with error when zip-code has wrong length", func(t *testing.T) {
		//given
		mockViaCEPGateway := gettemperaturebyzipcode.NewMockViaCEPGateway(t)
		mockWeatherAPIGateway := gettemperaturebyzipcode.NewMockWeatherAPIGateway(t)

		target := New(mockViaCEPGateway, mockWeatherAPIGateway)

		zipCode := address.ZipCode("0430300")

		//when
		resp, err := target.Execute(context.Background(), zipCode.Value())

		//then
		assert.EqualError(t, err, address.ErrInvalidZipCode.Error())
		assert.Nil(t, resp)

		mockViaCEPGateway.AssertNotCalled(t, "GetAddressByZipCode")
		mockWeatherAPIGateway.AssertNotCalled(t, "GetWeatherByCity")
	})

	t.Run("should execute with error when viaCepGateway returns error ", func(t *testing.T) {
		//given
		mockViaCEPGateway := gettemperaturebyzipcode.NewMockViaCEPGateway(t)
		mockWeatherAPIGateway := gettemperaturebyzipcode.NewMockWeatherAPIGateway(t)

		target := New(mockViaCEPGateway, mockWeatherAPIGateway)

		zipCode := address.ZipCode("04303001")

		expectedError := errors.New("some error")

		mockViaCEPGateway.EXPECT().
			GetAddressByZipCode(mock.Anything, mock.Anything).
			Return(nil, expectedError)

		//when
		resp, err := target.Execute(context.Background(), zipCode.Value())

		//then
		assert.EqualError(t, err, expectedError.Error())
		assert.Nil(t, resp)

		mockWeatherAPIGateway.AssertNotCalled(t, "GetWeatherByCity")
	})

	t.Run("should execute with error when weatherAPIGateway returns error ", func(t *testing.T) {
		//given
		mockViaCEPGateway := gettemperaturebyzipcode.NewMockViaCEPGateway(t)
		mockWeatherAPIGateway := gettemperaturebyzipcode.NewMockWeatherAPIGateway(t)

		target := New(mockViaCEPGateway, mockWeatherAPIGateway)

		zipCode := address.ZipCode("04303001")
		address := address.Address{
			ZipCode: zipCode,
			City:    "São Paulo",
		}

		expectedError := errors.New("some error")

		mockViaCEPGateway.EXPECT().
			GetAddressByZipCode(mock.Anything, mock.Anything).
			Return(&address, nil)

		mockWeatherAPIGateway.EXPECT().
			GetWeatherByCity(mock.Anything, mock.Anything).
			Return(nil, expectedError)

		//when
		resp, err := target.Execute(context.Background(), zipCode.Value())

		//then
		assert.EqualError(t, err, expectedError.Error())
		assert.Nil(t, resp)

		mockViaCEPGateway.AssertCalled(t, "GetAddressByZipCode", mock.Anything, mock.Anything)
	})
}
