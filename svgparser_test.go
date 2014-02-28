package svg2dxf

import (
    "testing"
)

func TestParseCircle(t *testing.T) {
    svg, err := ParseSVGFile("circle.svg")
    if err != nil {
        t.Error("Error parsing circle.svg:", err)
    }
    t.Log(svg)
}
