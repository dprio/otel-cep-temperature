package weather

type (
	Weather struct {
		City string
		C    float64
		F    float64
		K    float64
	}
)

func New(city string, c, f, k float64) *Weather {
	return &Weather{
		C: c,
		F: f,
		K: k,
	}
}
