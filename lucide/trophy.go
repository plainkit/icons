package lucide

import x "github.com/plainkit/blox"

// Trophy creates a Trophy Lucide icon.
func Trophy(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-trophy", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10 14.66v1.626a2 2 0 0 1-.976 1.696A5 5 0 0 0 7 21.978"))),
		x.Child(x.Path(x.D("M14 14.66v1.626a2 2 0 0 0 .976 1.696A5 5 0 0 1 17 21.978"))),
		x.Child(x.Path(x.D("M18 9h1.5a1 1 0 0 0 0-5H18"))),
		x.Child(x.Path(x.D("M4 22h16"))),
		x.Child(x.Path(x.D("M6 9a6 6 0 0 0 12 0V3a1 1 0 0 0-1-1H7a1 1 0 0 0-1 1z"))),
		x.Child(x.Path(x.D("M6 9H4.5a1 1 0 0 1 0-5H6"))),
	)
	return x.Svg(svgArgs...)
}
