package lucide

import x "github.com/plainkit/blox"

// ArrowDownWideNarrow creates a Arrow Down Wide Narrow Lucide icon.
func ArrowDownWideNarrow(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-arrow-down-wide-narrow", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m3 16 4 4 4-4"))),
		x.Child(x.Path(x.D("M7 20V4"))),
		x.Child(x.Path(x.D("M11 4h10"))),
		x.Child(x.Path(x.D("M11 8h7"))),
		x.Child(x.Path(x.D("M11 12h4"))),
	)
	return x.Svg(svgArgs...)
}
