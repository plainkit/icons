package lucide

import x "github.com/bloxui/blox"

// ArrowDownNarrowWide creates a Arrow Down Narrow Wide Lucide icon.
func ArrowDownNarrowWide(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-arrow-down-narrow-wide", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m3 16 4 4 4-4"))),
		x.Child(x.Path(x.D("M7 20V4"))),
		x.Child(x.Path(x.D("M11 4h4"))),
		x.Child(x.Path(x.D("M11 8h7"))),
		x.Child(x.Path(x.D("M11 12h10"))),
	)
	return x.Svg(svgArgs...)
}
