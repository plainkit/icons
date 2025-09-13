package lucide

import x "github.com/bloxui/blox"

// MoveUpRight creates a Move Up Right Lucide icon.
func MoveUpRight(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-move-up-right", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M13 5H19V11"))),
		x.Child(x.Path(x.D("M19 5L5 19"))),
	)
	return x.Svg(svgArgs...)
}
