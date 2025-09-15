package lucide

import x "github.com/plainkit/html"

// Weight creates a Weight Lucide icon.
func Weight(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-weight", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("12"), x.Cy("5"), x.R("3"))),
		x.Child(x.Path(x.D("M6.5 8a2 2 0 0 0-1.905 1.46L2.1 18.5A2 2 0 0 0 4 21h16a2 2 0 0 0 1.925-2.54L19.4 9.5A2 2 0 0 0 17.48 8Z"))),
	)
	return x.Svg(svgArgs...)
}
