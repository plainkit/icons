package lucide

import x "github.com/plainkit/html"

// CornerUpLeft creates a Corner Up Left Lucide icon.
func CornerUpLeft(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-corner-up-left", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M20 20v-7a4 4 0 0 0-4-4H4"))),
		x.Child(x.Path(x.D("M9 14 4 9l5-5"))),
	)
	return x.Svg(svgArgs...)
}
