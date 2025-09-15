package lucide

import x "github.com/plainkit/blox"

// CornerRightUp creates a Corner Right Up Lucide icon.
func CornerRightUp(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-corner-right-up", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m10 9 5-5 5 5"))),
		x.Child(x.Path(x.D("M4 20h7a4 4 0 0 0 4-4V4"))),
	)
	return x.Svg(svgArgs...)
}
