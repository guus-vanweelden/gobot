package leap

type Hand struct {
	ArmBasis               [][]float64 `json:"armBasis"`
	ArmWidth               float64     `json:"armWidth"`
	Confidence             float64     `json:"confidence"`
	Direction              []float64   `json:"direction"`
	Elbow                  []float64   `json:"elbow"`
	GrabStrength           float64     `json:"grabStrength"`
	ID                     int         `json:"id"`
	PalmNormal             []float64   `json:"palmNormal"`
	PalmPosition           []float64   `json:"PalmPosition"`
	PalmVelocity           []float64   `json:"PalmVelocity"`
	PinchStrength          float64     `json:"PinchStrength"`
	R                      [][]float64 `json:"r"`
	S                      float64     `json:"s"`
	SphereCenter           []float64   `json:"sphereCenter"`
	SphereRadius           float64     `json:"sphereRadius"`
	StabilizedPalmPosition []float64   `json:"stabilizedPalmPosition"`
	T                      []float64   `json:"t"`
	TimeVisible            float64     `json:"TimeVisible"`
	Type                   string      `json:"type"`
	Wrist                  []float64   `json:"wrist"`
}

// X returns hand x value
func (h *Hand) X() float64 {
	return h.PalmPosition[0]
}

// Y returns hand y value
func (h *Hand) Y() float64 {
	return h.PalmPosition[1]
}

// Z returns hand z value
func (h *Hand) Z() float64 {
	return h.PalmPosition[2]
}
