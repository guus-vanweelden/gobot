package leap

// Base representation returned that holds every other objects
type Frame struct {
	CurrentFrameRate float64        `json:"currentFrameRate"`
	Devices          []Device       `json:"devices"`
	Gestures         []Gesture      `json:"gestures"`
	Hands            []Hand         `json:"hands"`
	ID               int            `json:"id"`
	InteractionBox   InteractionBox `json:"interactionBox"`
	Pointables       []Pointable    `json:"pointables"`
	R                [][]float64    `json:"r"`
	S                float64        `json:"s"`
	T                []float64      `json:"t"`
	Timestamp        int            `json:"timestamp"`
}
