package lucide

import x "github.com/plainkit/html"

// Annoyed creates a Annoyed Lucide icon.
func Annoyed(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-annoyed", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("10"))),
		x.Child(x.Path(x.D("M8 15h8"))),
		x.Child(x.Path(x.D("M8 9h2"))),
		x.Child(x.Path(x.D("M14 9h2"))),
	)
	return x.Svg(svgArgs...)
}
