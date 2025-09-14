package lucide

import x "github.com/bloxui/blox"

// PaintRoller creates a Paint Roller Lucide icon.
func PaintRoller(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-paint-roller", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("16"), x.RectHeight("6"), x.X("2"), x.Y("2"), x.Rx("2"))),
		x.Child(x.Path(x.D("M10 16v-2a2 2 0 0 1 2-2h8a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2h-2"))),
		x.Child(x.Rect(x.RectWidth("4"), x.RectHeight("6"), x.X("8"), x.Y("16"), x.Rx("1"))),
	)
	return x.Svg(svgArgs...)
}
