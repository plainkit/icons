package lucide

import x "github.com/plainkit/blox"

// ArrowLeftToLine creates a Arrow Left To Line Lucide icon.
func ArrowLeftToLine(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-arrow-left-to-line", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M3 19V5"))),
		x.Child(x.Path(x.D("m13 6-6 6 6 6"))),
		x.Child(x.Path(x.D("M7 12h14"))),
	)
	return x.Svg(svgArgs...)
}
