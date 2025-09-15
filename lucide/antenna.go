package lucide

import x "github.com/plainkit/blox"

// Antenna creates a Antenna Lucide icon.
func Antenna(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-antenna", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M2 12 7 2"))),
		x.Child(x.Path(x.D("m7 12 5-10"))),
		x.Child(x.Path(x.D("m12 12 5-10"))),
		x.Child(x.Path(x.D("m17 12 5-10"))),
		x.Child(x.Path(x.D("M4.5 7h15"))),
		x.Child(x.Path(x.D("M12 16v6"))),
	)
	return x.Svg(svgArgs...)
}
