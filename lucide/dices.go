package lucide

import x "github.com/bloxui/blox"

// Dices creates a Dices Lucide icon.
func Dices(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-dices", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("12"), x.RectHeight("12"), x.X("2"), x.Y("10"), x.Rx("2"), x.Ry("2"))),
		x.Child(x.Path(x.D("m17.92 14 3.5-3.5a2.24 2.24 0 0 0 0-3l-5-4.92a2.24 2.24 0 0 0-3 0L10 6"))),
		x.Child(x.Path(x.D("M6 18h.01"))),
		x.Child(x.Path(x.D("M10 14h.01"))),
		x.Child(x.Path(x.D("M15 6h.01"))),
		x.Child(x.Path(x.D("M18 9h.01"))),
	)
	return x.Svg(svgArgs...)
}
