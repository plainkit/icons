package lucide

import x "github.com/bloxui/blox"

// ArrowRightFromLine creates a Arrow Right From Line Lucide icon.
func ArrowRightFromLine(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-arrow-right-from-line", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M3 5v14"))),
		x.Child(x.Path(x.D("M21 12H7"))),
		x.Child(x.Path(x.D("m15 18 6-6-6-6"))),
	)
	return x.Svg(svgArgs...)
}
