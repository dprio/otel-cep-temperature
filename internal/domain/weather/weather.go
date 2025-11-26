package weather

type (
	Weather struct {
		Temperature Temperature
	}

	Temperature struct {
		C float64
		F float64
		K float64
	}
)

func NewTemparature(c float64) *Temperature {
	return &Temperature{
		C: c,
		F: c*1.8 + 32,
		K: c + 273,
	}
}
