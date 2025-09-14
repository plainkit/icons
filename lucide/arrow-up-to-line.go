package lucide

import x "github.com/bloxui/blox"

// ArrowUpToLine creates a Arrow Up To Line Lucide icon.
func ArrowUpToLine(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-arrow-up-to-line", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M5 3h14"))),
		x.Child(x.Path(x.D("m18 13-6-6-6 6"))),
		x.Child(x.Path(x.D("M12 7v14"))),
	)
	return x.Svg(svgArgs...)
}
