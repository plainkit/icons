package lucide

import x "github.com/bloxui/blox"

// ArrowDownFromLine creates a Arrow Down From Line Lucide icon.
func ArrowDownFromLine(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-arrow-down-from-line", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M19 3H5"))),
		x.Child(x.Path(x.D("M12 21V7"))),
		x.Child(x.Path(x.D("m6 15 6 6 6-6"))),
	)
	return x.Svg(svgArgs...)
}
