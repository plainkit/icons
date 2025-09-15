package lucide

import x "github.com/plainkit/blox"

// ArrowRightToLine creates a Arrow Right To Line Lucide icon.
func ArrowRightToLine(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-arrow-right-to-line", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M17 12H3"))),
		x.Child(x.Path(x.D("m11 18 6-6-6-6"))),
		x.Child(x.Path(x.D("M21 5v14"))),
	)
	return x.Svg(svgArgs...)
}
