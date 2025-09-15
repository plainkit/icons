package lucide

import x "github.com/plainkit/html"

// Option creates a Option Lucide icon.
func Option(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-option", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M3 3h6l6 18h6"))),
		x.Child(x.Path(x.D("M14 3h7"))),
	)
	return x.Svg(svgArgs...)
}
