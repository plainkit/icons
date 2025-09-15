package lucide

import x "github.com/plainkit/blox"

// Rows3 creates a Rows 3 Lucide icon.
func Rows3(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-rows-3", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Path(x.D("M21 9H3"))),
		x.Child(x.Path(x.D("M21 15H3"))),
	)
	return x.Svg(svgArgs...)
}
