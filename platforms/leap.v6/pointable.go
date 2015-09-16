package leap

type Pointable struct {
	Bases                 [][][]float64 `json:"bases"`
	BTipPosition          []float64     `json:"btipPosition"`
	CarpPosition          []float64     `json:"carpPosition"`
	DipPosition           []float64     `json:"dipPosition"`
	Direction             []float64     `json:"direction"`
	Extended              bool          `json:"extended"`
	HandID                int           `json:"handId"`
	ID                    int           `json:"id"`
	Length                float64       `json:"length"`
	MCPPosition           []float64     `json:"mcpPosition"`
	PIPPosition           []float64     `json:"pipPosition"`
	StabilizedTipPosition []float64     `json:"stabilizedTipPosition"`
	TimeVisible           float64       `json:"timeVisible"`
	TipPosition           []float64     `json:"tipPosition"`
	TipVelocity           []float64     `json:"tipVelocity"`
	Tool                  bool          `json:"tool"`
	TouchDistance         float64       `json:"touchDistance"`
	TouchZone             string        `json:"touchZone"`
	Type                  int           `json:"type"`
	Width                 float64       `json:"width"`
}
