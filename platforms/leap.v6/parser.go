package leap

import (
	"encoding/json"
)

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

type Pointable struct {
	Bases                 [][]float64 `json:"bases"`
	BTipPosition          []float64   `json:"btipPosition"`
	CarpPosition          []float64   `json:"carpPosition"`
	DipPosition           []float64   `json:"dipPosition"`
	Direction             []float64   `json:"direction"`
	Extended              bool        `json:"extended"`
	HandID                int         `json:"handId"`
	ID                    int         `json:"id"`
	Length                float64     `json:"length"`
	MCPPosition           []float64   `json:"mcpPosition"`
	PIPPosition           []float64   `json:"pipPosition"`
	StabilizedTipPosition []float64   `json:"stabilizedTipPosition"`
	TimeVisible           float64     `json:"timeVisible"`
	TipPosition           []float64   `json:"tipPosition"`
	TipVelocity           []float64   `json:"tipVelocity"`
	Tool                  bool        `json:"tool"`
	TouchDistance         float64     `json:"touchDistance"`
	TouchZone             string      `json:"touchZone"`
	Type                  int         `json:"type"`
	Width                 float64     `json:"width"`
}

type InteractionBox struct {
	Center []int     `json:"center"`
	Size   []float64 `json:"size"`
}

type Device struct {
}

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

// ParseFrame converts json data to a Frame
func (l *LeapMotionDriver) ParseFrame(data []byte) Frame {
	var frame Frame
	json.Unmarshal(data, &frame)
	return frame
}
