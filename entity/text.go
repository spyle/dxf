package entity

import (
	"github.com/yofu/dxf/format"
	"github.com/yofu/dxf/table"
)

const (
	LEFT_BASE = iota
	CENTER_BASE
	RIGHT_BASE
	LEFT_BOTTOM
	CENTER_BOTTOM
	RIGHT_BOTTOM
	LEFT_CENTER
	CENTER_CENTER
	RIGHT_CENTER
	LEFT_TOP
	CENTER_TOP
	RIGHT_TOP
)

type Text struct {
	*entity
	Coord1         []float64    // 10, 20, 30
	Coord2         []float64    // 11, 21, 31
	Height         float64      // 40
	Value          string       // 1
	Style          *table.Style // 7
	genflag        int          // 71
	horizontalflag int          // 72
	verticalflag   int          // 73
}

func (t *Text) IsEntity() bool {
	return true
}

func NewText() *Text {
	t := &Text{
		entity:         NewEntity(TEXT),
		Coord1:         []float64{0.0, 0.0, 0.0},
		Coord2:         []float64{0.0, 0.0, 0.0},
		Height:         1.0,
		Value:          "",
		Style:          table.ST_STANDARD,
		genflag:        0,
		horizontalflag: 0,
		verticalflag:   0,
	}
	return t
}

func (t *Text) Format(f *format.Formatter) {
	t.entity.Format(f)
	f.WriteString(100, "AcDbText")
	for i := 0; i < 3; i++ {
		f.WriteFloat((i+1)*10, t.Coord1[i])
	}
	f.WriteFloat(40, t.Height)
	f.WriteString(1, t.Value)
	f.WriteString(7, t.Style.Name)
	if t.genflag != 0 {
		f.WriteInt(71, t.genflag)
	}
	if t.horizontalflag != 0 {
		f.WriteInt(72, t.horizontalflag)
		if t.verticalflag != 0 {
			for i := 0; i < 3; i++ {
				f.WriteFloat((i+1)*11, t.Coord1[i])
			}
		}
	}
	f.WriteString(100, "AcDbText")
	if t.verticalflag != 0 {
		f.WriteInt(73, t.verticalflag)
	}
}

func (t *Text) String() string {
	f := format.New()
	return t.FormatString(f)
}

func (t *Text) FormatString(f *format.Formatter) string {
	t.Format(f)
	return f.Output()
}

func (t *Text) togglegenflag(val int) {
	if t.genflag&val != 0 {
		t.genflag &= ^val
	} else {
		t.genflag |= val
	}
}
func (t *Text) FlipHorizontal() {
	t.togglegenflag(2)
}
func (t *Text) FlipVertical() {
	t.togglegenflag(4)
}

func (t *Text) Anchor(pos int) {
	switch pos {
	case LEFT_BASE:
		t.horizontalflag = 0
		t.verticalflag = 0
	case CENTER_BASE:
		t.horizontalflag = 1
		t.verticalflag = 0
	case RIGHT_BASE:
		t.horizontalflag = 2
		t.verticalflag = 0
	case LEFT_BOTTOM:
		t.horizontalflag = 0
		t.verticalflag = 1
	case CENTER_BOTTOM:
		t.horizontalflag = 1
		t.verticalflag = 1
	case RIGHT_BOTTOM:
		t.horizontalflag = 2
		t.verticalflag = 1
	case LEFT_CENTER:
		t.horizontalflag = 0
		t.verticalflag = 2
	case CENTER_CENTER:
		t.horizontalflag = 1
		t.verticalflag = 2
	case RIGHT_CENTER:
		t.horizontalflag = 2
		t.verticalflag = 2
	case LEFT_TOP:
		t.horizontalflag = 0
		t.verticalflag = 3
	case CENTER_TOP:
		t.horizontalflag = 1
		t.verticalflag = 3
	case RIGHT_TOP:
		t.horizontalflag = 2
		t.verticalflag = 3
	}
}
