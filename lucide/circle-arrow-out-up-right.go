package lucide

import x "github.com/plainkit/html"

// CircleArrowOutUpRight creates a Circle Arrow Out Up Right Lucide icon.
func CircleArrowOutUpRight(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-circle-arrow-out-up-right", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M22 12A10 10 0 1 1 12 2"))),
		x.Child(x.Path(x.D("M22 2 12 12"))),
		x.Child(x.Path(x.D("M16 2h6v6"))),
	)
	return x.Svg(svgArgs...)
}
