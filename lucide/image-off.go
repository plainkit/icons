package lucide

import x "github.com/plainkit/html"

// ImageOff creates a Image Off Lucide icon.
func ImageOff(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-image-off", args)
	svgArgs = append(svgArgs,
		x.Child(x.Line(x.X1("2"), x.X2("22"), x.Y1("2"), x.Y2("22"))),
		x.Child(x.Path(x.D("M10.41 10.41a2 2 0 1 1-2.83-2.83"))),
		x.Child(x.Line(x.X1("13.5"), x.X2("6"), x.Y1("13.5"), x.Y2("21"))),
		x.Child(x.Line(x.X1("18"), x.X2("21"), x.Y1("12"), x.Y2("15"))),
		x.Child(x.Path(x.D("M3.59 3.59A1.99 1.99 0 0 0 3 5v14a2 2 0 0 0 2 2h14c.55 0 1.052-.22 1.41-.59"))),
		x.Child(x.Path(x.D("M21 15V5a2 2 0 0 0-2-2H9"))),
	)
	return x.Svg(svgArgs...)
}
