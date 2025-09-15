package lucide

import x "github.com/plainkit/blox"

// Package creates a Package Lucide icon.
func Package(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-package", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M11 21.73a2 2 0 0 0 2 0l7-4A2 2 0 0 0 21 16V8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73z"))),
		x.Child(x.Path(x.D("M12 22V12"))),
		x.Child(x.Polyline(x.Points("3.29 7 12 12 20.71 7"))),
		x.Child(x.Path(x.D("m7.5 4.27 9 5.15"))),
	)
	return x.Svg(svgArgs...)
}
