package lucide

import x "github.com/plainkit/html"

// Focus creates a Focus Lucide icon.
func Focus(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-focus", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("3"))),
		x.Child(x.Path(x.D("M3 7V5a2 2 0 0 1 2-2h2"))),
		x.Child(x.Path(x.D("M17 3h2a2 2 0 0 1 2 2v2"))),
		x.Child(x.Path(x.D("M21 17v2a2 2 0 0 1-2 2h-2"))),
		x.Child(x.Path(x.D("M7 21H5a2 2 0 0 1-2-2v-2"))),
	)
	return x.Svg(svgArgs...)
}
