package lucide

import x "github.com/plainkit/blox"

// CheckLine creates a Check Line Lucide icon.
func CheckLine(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-check-line", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M20 4L9 15"))),
		x.Child(x.Path(x.D("M21 19L3 19"))),
		x.Child(x.Path(x.D("M9 15L4 10"))),
	)
	return x.Svg(svgArgs...)
}
