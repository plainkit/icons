package lucide

import x "github.com/plainkit/html"

// SquareArrowOutUpLeft creates a Square Arrow Out Up Left Lucide icon.
func SquareArrowOutUpLeft(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-square-arrow-out-up-left", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M13 3h6a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-6"))),
		x.Child(x.Path(x.D("m3 3 9 9"))),
		x.Child(x.Path(x.D("M3 9V3h6"))),
	)
	return x.Svg(svgArgs...)
}
