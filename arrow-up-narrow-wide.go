package lucide

import x "github.com/bloxui/blox"

// ArrowUpNarrowWide creates an Arrow Up Narrow Wide Lucide icon.
func ArrowUpNarrowWide(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-arrow-up-narrow-wide", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m3 8 4-4 4 4"))),
		x.Child(x.Path(x.D("M7 4v16"))),
		x.Child(x.Path(x.D("M11 12h4"))),
		x.Child(x.Path(x.D("M11 16h7"))),
		x.Child(x.Path(x.D("M11 20h10"))),
	)
	return x.Svg(svgArgs...)
}
