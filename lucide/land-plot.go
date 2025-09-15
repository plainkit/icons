package lucide

import x "github.com/plainkit/html"

// LandPlot creates a Land Plot Lucide icon.
func LandPlot(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-land-plot", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m12 8 6-3-6-3v10"))),
		x.Child(x.Path(x.D("m8 11.99-5.5 3.14a1 1 0 0 0 0 1.74l8.5 4.86a2 2 0 0 0 2 0l8.5-4.86a1 1 0 0 0 0-1.74L16 12"))),
		x.Child(x.Path(x.D("m6.49 12.85 11.02 6.3"))),
		x.Child(x.Path(x.D("M17.51 12.85 6.5 19.15"))),
	)
	return x.Svg(svgArgs...)
}
