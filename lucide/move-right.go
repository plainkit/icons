package lucide

import x "github.com/bloxui/blox"

// MoveRight creates a Move Right Lucide icon.
func MoveRight(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-move-right", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M18 8L22 12L18 16"))),
		x.Child(x.Path(x.D("M2 12H22"))),
	)
	return x.Svg(svgArgs...)
}
