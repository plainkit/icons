package lucide

import x "github.com/plainkit/blox"

// Pilcrow creates a Pilcrow Lucide icon.
func Pilcrow(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-pilcrow", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M13 4v16"))),
		x.Child(x.Path(x.D("M17 4v16"))),
		x.Child(x.Path(x.D("M19 4H9.5a4.5 4.5 0 0 0 0 9H13"))),
	)
	return x.Svg(svgArgs...)
}
