package lucide

import x "github.com/plainkit/html"

// MoveUp creates a Move Up Lucide icon.
func MoveUp(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-move-up", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M8 6L12 2L16 6"))),
		x.Child(x.Path(x.D("M12 2V22"))),
	)
	return x.Svg(svgArgs...)
}
