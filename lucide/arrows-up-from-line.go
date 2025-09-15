package lucide

import x "github.com/plainkit/html"

// ArrowsUpFromLine creates a Arrows Up From Line Lucide icon.
func ArrowsUpFromLine(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-arrows-up-from-line", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m4 6 3-3 3 3"))),
		x.Child(x.Path(x.D("M7 17V3"))),
		x.Child(x.Path(x.D("m14 6 3-3 3 3"))),
		x.Child(x.Path(x.D("M17 17V3"))),
		x.Child(x.Path(x.D("M4 21h16"))),
	)
	return x.Svg(svgArgs...)
}
