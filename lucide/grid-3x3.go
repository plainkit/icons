package lucide

import x "github.com/plainkit/blox"

// Grid3x3 creates a Grid 3x3 Lucide icon.
func Grid3x3(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-grid-3x3", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Path(x.D("M3 9h18"))),
		x.Child(x.Path(x.D("M3 15h18"))),
		x.Child(x.Path(x.D("M9 3v18"))),
		x.Child(x.Path(x.D("M15 3v18"))),
	)
	return x.Svg(svgArgs...)
}
