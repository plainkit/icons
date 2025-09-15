package lucide

import x "github.com/plainkit/html"

// CircleArrowOutDownRight creates a Circle Arrow Out Down Right Lucide icon.
func CircleArrowOutDownRight(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-circle-arrow-out-down-right", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 22a10 10 0 1 1 10-10"))),
		x.Child(x.Path(x.D("M22 22 12 12"))),
		x.Child(x.Path(x.D("M22 16v6h-6"))),
	)
	return x.Svg(svgArgs...)
}
