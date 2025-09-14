package lucide

import x "github.com/bloxui/blox"

// SquareTerminal creates a Square Terminal Lucide icon.
func SquareTerminal(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-square-terminal", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m7 11 2-2-2-2"))),
		x.Child(x.Path(x.D("M11 13h4"))),
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"), x.Ry("2"))),
	)
	return x.Svg(svgArgs...)
}
