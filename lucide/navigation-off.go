package lucide

import x "github.com/plainkit/html"

// NavigationOff creates a Navigation Off Lucide icon.
func NavigationOff(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-navigation-off", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M8.43 8.43 3 11l8 2 2 8 2.57-5.43"))),
		x.Child(x.Path(x.D("M17.39 11.73 22 2l-9.73 4.61"))),
		x.Child(x.Line(x.X1("2"), x.X2("22"), x.Y1("2"), x.Y2("22"))),
	)
	return x.Svg(svgArgs...)
}
