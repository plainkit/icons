package lucide

import x "github.com/plainkit/html"

// CornerDownLeft creates a Corner Down Left Lucide icon.
func CornerDownLeft(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-corner-down-left", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M20 4v7a4 4 0 0 1-4 4H4"))),
		x.Child(x.Path(x.D("m9 10-5 5 5 5"))),
	)
	return x.Svg(svgArgs...)
}
