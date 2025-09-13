package lucide

import x "github.com/bloxui/blox"

// Dice6 creates a Dice 6 Lucide icon.
func Dice6(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-dice-6", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"), x.Ry("2"))),
		x.Child(x.Path(x.D("M16 8h.01"))),
		x.Child(x.Path(x.D("M16 12h.01"))),
		x.Child(x.Path(x.D("M16 16h.01"))),
		x.Child(x.Path(x.D("M8 8h.01"))),
		x.Child(x.Path(x.D("M8 12h.01"))),
		x.Child(x.Path(x.D("M8 16h.01"))),
	)
	return x.Svg(svgArgs...)
}
