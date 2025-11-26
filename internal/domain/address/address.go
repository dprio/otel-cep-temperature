package address

type Address struct {
	ZipCode string
	City    string
}

func New(zipCode, city string) *Address {
	return &Address{
		ZipCode: zipCode,
		City:    city,
	}
}
