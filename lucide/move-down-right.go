package lucide

import x "github.com/plainkit/blox"

// MoveDownRight creates a Move Down Right Lucide icon.
func MoveDownRight(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-move-down-right", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M19 13V19H13"))),
		x.Child(x.Path(x.D("M5 5L19 19"))),
	)
	return x.Svg(svgArgs...)
}
