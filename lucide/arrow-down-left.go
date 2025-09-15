package lucide

import x "github.com/plainkit/html"

// ArrowDownLeft creates a Arrow Down Left Lucide icon.
func ArrowDownLeft(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-arrow-down-left", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M17 7 7 17"))),
		x.Child(x.Path(x.D("M17 17H7V7"))),
	)
	return x.Svg(svgArgs...)
}
