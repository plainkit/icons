package lucide

import x "github.com/plainkit/blox"

// MoveDownLeft creates a Move Down Left Lucide icon.
func MoveDownLeft(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-move-down-left", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M11 19H5V13"))),
		x.Child(x.Path(x.D("M19 5L5 19"))),
	)
	return x.Svg(svgArgs...)
}
