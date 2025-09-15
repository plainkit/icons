package lucide

import x "github.com/plainkit/html"

// CloudLightning creates a Cloud Lightning Lucide icon.
func CloudLightning(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-cloud-lightning", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M6 16.326A7 7 0 1 1 15.71 8h1.79a4.5 4.5 0 0 1 .5 8.973"))),
		x.Child(x.Path(x.D("m13 12-3 5h4l-3 5"))),
	)
	return x.Svg(svgArgs...)
}
