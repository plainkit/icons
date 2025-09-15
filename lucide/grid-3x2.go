package lucide

import x "github.com/plainkit/html"

// Grid3x2 creates a Grid 3x2 Lucide icon.
func Grid3x2(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-grid-3x2", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M15 3v18"))),
		x.Child(x.Path(x.D("M3 12h18"))),
		x.Child(x.Path(x.D("M9 3v18"))),
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
	)
	return x.Svg(svgArgs...)
}
