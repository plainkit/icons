package lucide

import x "github.com/plainkit/blox"

// PackageX creates a Package X Lucide icon.
func PackageX(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-package-x", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M21 10V8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73l7 4a2 2 0 0 0 2 0l2-1.14"))),
		x.Child(x.Path(x.D("m7.5 4.27 9 5.15"))),
		x.Child(x.Polyline(x.Points("3.29 7 12 12 20.71 7"))),
		x.Child(x.Line(x.X1("12"), x.X2("12"), x.Y1("22"), x.Y2("12"))),
		x.Child(x.Path(x.D("m17 13 5 5m-5 0 5-5"))),
	)
	return x.Svg(svgArgs...)
}
