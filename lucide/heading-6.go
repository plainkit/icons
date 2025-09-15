package lucide

import x "github.com/plainkit/html"

// Heading6 creates a Heading 6 Lucide icon.
func Heading6(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-heading-6", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M4 12h8"))),
		x.Child(x.Path(x.D("M4 18V6"))),
		x.Child(x.Path(x.D("M12 18V6"))),
		x.Child(x.Circle(x.Cx("19"), x.Cy("16"), x.R("2"))),
		x.Child(x.Path(x.D("M20 10c-2 2-3 3.5-3 6"))),
	)
	return x.Svg(svgArgs...)
}
