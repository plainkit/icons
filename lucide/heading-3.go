package lucide

import x "github.com/plainkit/html"

// Heading3 creates a Heading 3 Lucide icon.
func Heading3(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-heading-3", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M4 12h8"))),
		x.Child(x.Path(x.D("M4 18V6"))),
		x.Child(x.Path(x.D("M12 18V6"))),
		x.Child(x.Path(x.D("M17.5 10.5c1.7-1 3.5 0 3.5 1.5a2 2 0 0 1-2 2"))),
		x.Child(x.Path(x.D("M17 17.5c2 1.5 4 .3 4-1.5a2 2 0 0 0-2-2"))),
	)
	return x.Svg(svgArgs...)
}
