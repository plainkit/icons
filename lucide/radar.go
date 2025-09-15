package lucide

import x "github.com/plainkit/blox"

// Radar creates a Radar Lucide icon.
func Radar(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-radar", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M19.07 4.93A10 10 0 0 0 6.99 3.34"))),
		x.Child(x.Path(x.D("M4 6h.01"))),
		x.Child(x.Path(x.D("M2.29 9.62A10 10 0 1 0 21.31 8.35"))),
		x.Child(x.Path(x.D("M16.24 7.76A6 6 0 1 0 8.23 16.67"))),
		x.Child(x.Path(x.D("M12 18h.01"))),
		x.Child(x.Path(x.D("M17.99 11.66A6 6 0 0 1 15.77 16.67"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("2"))),
		x.Child(x.Path(x.D("m13.41 10.59 5.66-5.66"))),
	)
	return x.Svg(svgArgs...)
}
