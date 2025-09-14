package lucide

import x "github.com/bloxui/blox"

// SquareRoundCorner creates a Square Round Corner Lucide icon.
func SquareRoundCorner(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-square-round-corner", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M21 11a8 8 0 0 0-8-8"))),
		x.Child(x.Path(x.D("M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h4"))),
	)
	return x.Svg(svgArgs...)
}
