package lucide

import x "github.com/plainkit/html"

// BellDot creates a Bell Dot Lucide icon.
func BellDot(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-bell-dot", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10.268 21a2 2 0 0 0 3.464 0"))),
		x.Child(x.Path(x.D("M13.916 2.314A6 6 0 0 0 6 8c0 4.499-1.411 5.956-2.74 7.327A1 1 0 0 0 4 17h16a1 1 0 0 0 .74-1.673 9 9 0 0 1-.585-.665"))),
		x.Child(x.Circle(x.Cx("18"), x.Cy("8"), x.R("3"))),
	)
	return x.Svg(svgArgs...)
}
