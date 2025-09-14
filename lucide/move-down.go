package lucide

import x "github.com/bloxui/blox"

// MoveDown creates a Move Down Lucide icon.
func MoveDown(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-move-down", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M8 18L12 22L16 18"))),
		x.Child(x.Path(x.D("M12 2V22"))),
	)
	return x.Svg(svgArgs...)
}
