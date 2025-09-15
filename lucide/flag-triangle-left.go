package lucide

import x "github.com/plainkit/blox"

// FlagTriangleLeft creates a Flag Triangle Left Lucide icon.
func FlagTriangleLeft(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-flag-triangle-left", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M18 22V2.8a.8.8 0 0 0-1.17-.71L5.45 7.78a.8.8 0 0 0 0 1.44L18 15.5"))),
	)
	return x.Svg(svgArgs...)
}
