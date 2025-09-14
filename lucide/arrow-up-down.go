package lucide

import x "github.com/bloxui/blox"

// ArrowUpDown creates a Arrow Up Down Lucide icon.
func ArrowUpDown(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-arrow-up-down", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m21 16-4 4-4-4"))),
		x.Child(x.Path(x.D("M17 20V4"))),
		x.Child(x.Path(x.D("m3 8 4-4 4 4"))),
		x.Child(x.Path(x.D("M7 4v16"))),
	)
	return x.Svg(svgArgs...)
}
