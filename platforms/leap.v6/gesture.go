package leap

type Gesture struct {
	// circle only
	Center []float64 `json:"center"`
	// swipe, keyTap, screenTap only
	Direction    []float64 `json:"direction"`
	Duration     int       `json:"duration"`
	HandIDs      []int     `json:"handIds"`
	ID           int       `json:"id"`
	PointableIDs []int     `json:"pointablesIds"`
	// swipe, keyTap, screenTap only
	Position []float64 `json:"position"`
	// circle, keyTap, screenTap only
	Progress float64 `json:"progress"`
	// circle only
	Radius float64 `json:"radius"`
	// swipe only
	Speed float64 `json:"speed"`
	// swipe only
	StartPosition []float64 `json:"StartPosition"`
	// one of "start", "update", "stop"
	State string `json:"state"`
	// one of "circle", "swipe", "keyTap", "screenTap"
	Type string `json:"type"`
}
