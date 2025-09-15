package lucide

import x "github.com/plainkit/html"

// BeanOff creates a Bean Off Lucide icon.
func BeanOff(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-bean-off", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M9 9c-.64.64-1.521.954-2.402 1.165A6 6 0 0 0 8 22a13.96 13.96 0 0 0 9.9-4.1"))),
		x.Child(x.Path(x.D("M10.75 5.093A6 6 0 0 1 22 8c0 2.411-.61 4.68-1.683 6.66"))),
		x.Child(x.Path(x.D("M5.341 10.62a4 4 0 0 0 6.487 1.208M10.62 5.341a4.015 4.015 0 0 1 2.039 2.04"))),
		x.Child(x.Line(x.X1("2"), x.X2("22"), x.Y1("2"), x.Y2("22"))),
	)
	return x.Svg(svgArgs...)
}
