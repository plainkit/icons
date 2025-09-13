package lucide

import x "github.com/bloxui/blox"

// MoveUpLeft creates a Move Up Left Lucide icon.
func MoveUpLeft(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-move-up-left", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M5 11V5H11"))),
		x.Child(x.Path(x.D("M5 5L19 19"))),
	)
	return x.Svg(svgArgs...)
}
