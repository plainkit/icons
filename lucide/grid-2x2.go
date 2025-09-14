package lucide

import x "github.com/bloxui/blox"

// Grid2x2 creates a Grid 2x2 Lucide icon.
func Grid2x2(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-grid-2x2", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 3v18"))),
		x.Child(x.Path(x.D("M3 12h18"))),
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
	)
	return x.Svg(svgArgs...)
}
