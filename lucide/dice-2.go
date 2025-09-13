package lucide

import x "github.com/bloxui/blox"

// Dice2 creates a Dice 2 Lucide icon.
func Dice2(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-dice-2", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.X("3"), x.Y("3"), x.RectWidth("18"), x.RectHeight("18"), x.Rx("2"), x.Ry("2"))),
		x.Child(x.Path(x.D("M15 9h.01"))),
		x.Child(x.Path(x.D("M9 15h.01"))),
	)
	return x.Svg(svgArgs...)
}
