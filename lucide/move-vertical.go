package lucide

import x "github.com/bloxui/blox"

// MoveVertical creates a Move Vertical Lucide icon.
func MoveVertical(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-move-vertical", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 2v20"))),
		x.Child(x.Path(x.D("m8 18 4 4 4-4"))),
		x.Child(x.Path(x.D("m8 6 4-4 4 4"))),
	)
	return x.Svg(svgArgs...)
}
