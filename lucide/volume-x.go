package lucide

import x "github.com/plainkit/html"

// VolumeX creates a Volume X Lucide icon.
func VolumeX(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-volume-x", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M11 4.702a.705.705 0 0 0-1.203-.498L6.413 7.587A1.4 1.4 0 0 1 5.416 8H3a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h2.416a1.4 1.4 0 0 1 .997.413l3.383 3.384A.705.705 0 0 0 11 19.298z"))),
		x.Child(x.Line(x.X1("22"), x.X2("16"), x.Y1("9"), x.Y2("15"))),
		x.Child(x.Line(x.X1("16"), x.X2("22"), x.Y1("9"), x.Y2("15"))),
	)
	return x.Svg(svgArgs...)
}
