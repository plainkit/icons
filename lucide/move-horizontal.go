package lucide

import x "github.com/bloxui/blox"

// MoveHorizontal creates a Move Horizontal Lucide icon.
func MoveHorizontal(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-move-horizontal", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m18 8 4 4-4 4"))),
		x.Child(x.Path(x.D("M2 12h20"))),
		x.Child(x.Path(x.D("m6 8-4 4 4 4"))),
	)
	return x.Svg(svgArgs...)
}
