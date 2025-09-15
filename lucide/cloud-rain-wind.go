package lucide

import x "github.com/plainkit/html"

// CloudRainWind creates a Cloud Rain Wind Lucide icon.
func CloudRainWind(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-cloud-rain-wind", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M4 14.899A7 7 0 1 1 15.71 8h1.79a4.5 4.5 0 0 1 2.5 8.242"))),
		x.Child(x.Path(x.D("m9.2 22 3-7"))),
		x.Child(x.Path(x.D("m9 13-3 7"))),
		x.Child(x.Path(x.D("m17 13-3 7"))),
	)
	return x.Svg(svgArgs...)
}
