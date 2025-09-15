package lucide

import x "github.com/plainkit/html"

// IndianRupee creates a Indian Rupee Lucide icon.
func IndianRupee(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-indian-rupee", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M6 3h12"))),
		x.Child(x.Path(x.D("M6 8h12"))),
		x.Child(x.Path(x.D("m6 13 8.5 8"))),
		x.Child(x.Path(x.D("M6 13h3"))),
		x.Child(x.Path(x.D("M9 13c6.667 0 6.667-10 0-10"))),
	)
	return x.Svg(svgArgs...)
}
