package lucide

import x "github.com/plainkit/html"

// SunSnow creates a Sun Snow Lucide icon.
func SunSnow(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-sun-snow", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10 21v-1"))),
		x.Child(x.Path(x.D("M10 4V3"))),
		x.Child(x.Path(x.D("M10 9a3 3 0 0 0 0 6"))),
		x.Child(x.Path(x.D("m14 20 1.25-2.5L18 18"))),
		x.Child(x.Path(x.D("m14 4 1.25 2.5L18 6"))),
		x.Child(x.Path(x.D("m17 21-3-6 1.5-3H22"))),
		x.Child(x.Path(x.D("m17 3-3 6 1.5 3"))),
		x.Child(x.Path(x.D("M2 12h1"))),
		x.Child(x.Path(x.D("m20 10-1.5 2 1.5 2"))),
		x.Child(x.Path(x.D("m3.64 18.36.7-.7"))),
		x.Child(x.Path(x.D("m4.34 6.34-.7-.7"))),
	)
	return x.Svg(svgArgs...)
}
