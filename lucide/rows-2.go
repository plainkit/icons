package lucide

import x "github.com/plainkit/blox"

// Rows2 creates a Rows 2 Lucide icon.
func Rows2(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-rows-2", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Path(x.D("M3 12h18"))),
	)
	return x.Svg(svgArgs...)
}
