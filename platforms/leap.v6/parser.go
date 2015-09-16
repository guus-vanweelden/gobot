package leap

import (
	"encoding/json"
)

// ParseFrame converts json data to a Frame
func (l *LeapMotionDriver) ParseFrame(data []byte) Frame {
	var frame Frame
	json.Unmarshal(data, &frame)
	return frame
}
