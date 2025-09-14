package lucide

import x "github.com/bloxui/blox"

// MoveLeft creates a Move Left Lucide icon.
func MoveLeft(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-move-left", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M6 8L2 12L6 16"))),
		x.Child(x.Path(x.D("M2 12H22"))),
	)
	return x.Svg(svgArgs...)
}
