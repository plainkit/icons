package lucide

import x "github.com/bloxui/blox"

// CornerRightDown creates a Corner Right Down Lucide icon.
func CornerRightDown(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-corner-right-down", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m10 15 5 5 5-5"))),
		x.Child(x.Path(x.D("M4 4h7a4 4 0 0 1 4 4v12"))),
	)
	return x.Svg(svgArgs...)
}