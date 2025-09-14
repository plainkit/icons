package lucide

import x "github.com/bloxui/blox"

// CornerLeftDown creates a Corner Left Down Lucide icon.
func CornerLeftDown(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-corner-left-down", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m14 15-5 5-5-5"))),
		x.Child(x.Path(x.D("M20 4h-7a4 4 0 0 0-4 4v12"))),
	)
	return x.Svg(svgArgs...)
}
