package lucide

import x "github.com/plainkit/html"

// SquareArrowOutUpRight creates a Square Arrow Out Up Right Lucide icon.
func SquareArrowOutUpRight(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-square-arrow-out-up-right", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M21 13v6a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h6"))),
		x.Child(x.Path(x.D("m21 3-9 9"))),
		x.Child(x.Path(x.D("M15 3h6v6"))),
	)
	return x.Svg(svgArgs...)
}
