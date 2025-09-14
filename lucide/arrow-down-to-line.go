package lucide

import x "github.com/bloxui/blox"

// ArrowDownToLine creates a Arrow Down To Line Lucide icon.
func ArrowDownToLine(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-arrow-down-to-line", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 17V3"))),
		x.Child(x.Path(x.D("m6 11 6 6 6-6"))),
		x.Child(x.Path(x.D("M19 21H5"))),
	)
	return x.Svg(svgArgs...)
}
