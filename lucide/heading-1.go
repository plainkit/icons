package lucide

import x "github.com/plainkit/html"

// Heading1 creates a Heading 1 Lucide icon.
func Heading1(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-heading-1", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M4 12h8"))),
		x.Child(x.Path(x.D("M4 18V6"))),
		x.Child(x.Path(x.D("M12 18V6"))),
		x.Child(x.Path(x.D("m17 12 3-2v8"))),
	)
	return x.Svg(svgArgs...)
}
