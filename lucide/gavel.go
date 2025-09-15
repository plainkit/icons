package lucide

import x "github.com/plainkit/blox"

// Gavel creates a Gavel Lucide icon.
func Gavel(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-gavel", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m14 13-8.381 8.38a1 1 0 0 1-3.001-3l8.384-8.381"))),
		x.Child(x.Path(x.D("m16 16 6-6"))),
		x.Child(x.Path(x.D("m21.5 10.5-8-8"))),
		x.Child(x.Path(x.D("m8 8 6-6"))),
		x.Child(x.Path(x.D("m8.5 7.5 8 8"))),
	)
	return x.Svg(svgArgs...)
}
