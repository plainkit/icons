package lucide

import x "github.com/plainkit/html"

// Heading5 creates a Heading 5 Lucide icon.
func Heading5(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-heading-5", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M4 12h8"))),
		x.Child(x.Path(x.D("M4 18V6"))),
		x.Child(x.Path(x.D("M12 18V6"))),
		x.Child(x.Path(x.D("M17 13v-3h4"))),
		x.Child(x.Path(x.D("M17 17.7c.4.2.8.3 1.3.3 1.5 0 2.7-1.1 2.7-2.5S19.8 13 18.3 13H17"))),
	)
	return x.Svg(svgArgs...)
}
