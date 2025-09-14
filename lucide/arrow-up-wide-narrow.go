package lucide

import x "github.com/bloxui/blox"

// ArrowUpWideNarrow creates a Arrow Up Wide Narrow Lucide icon.
func ArrowUpWideNarrow(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-arrow-up-wide-narrow", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m3 8 4-4 4 4"))),
		x.Child(x.Path(x.D("M7 4v16"))),
		x.Child(x.Path(x.D("M11 12h10"))),
		x.Child(x.Path(x.D("M11 16h7"))),
		x.Child(x.Path(x.D("M11 20h4"))),
	)
	return x.Svg(svgArgs...)
}
