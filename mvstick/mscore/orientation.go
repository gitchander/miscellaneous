package mscore

import (
	"fmt"
)

type Orientation int

const (
	OriUnknown Orientation = iota
	OriHorizontal
	OriVertical
)

var namesOrientation = map[Orientation]string{
	OriHorizontal: "horizontal",
	OriVertical:   "vertical",
}

var valuesOrientation = map[string]Orientation{
	"horizontal": OriHorizontal,
	"vertical":   OriVertical,
}

func (v Orientation) String() string {
	if name, ok := namesOrientation[v]; ok {
		return name
	}
	return fmt.Sprintf("%s(%d)", "Orientation", int(v))
}

func ParseOrientation(s string) (Orientation, error) {
	if v, ok := valuesOrientation[s]; ok {
		return v, nil
	}
	var zv Orientation // zero value
	return zv, fmt.Errorf("invalid orientation name (%s)", s)
}
