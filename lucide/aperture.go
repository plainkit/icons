package lucide

import x "github.com/bloxui/blox"

// Aperture creates a Aperture Lucide icon.
func Aperture(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-aperture", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("10"))),
		x.Child(x.Path(x.D("m14.31 8 5.74 9.94"))),
		x.Child(x.Path(x.D("M9.69 8h11.48"))),
		x.Child(x.Path(x.D("m7.38 12 5.74-9.94"))),
		x.Child(x.Path(x.D("M9.69 16 3.95 6.06"))),
		x.Child(x.Path(x.D("M14.31 16H2.83"))),
		x.Child(x.Path(x.D("m16.62 12-5.74 9.94"))),
	)
	return x.Svg(svgArgs...)
}
