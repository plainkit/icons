package lucide

import x "github.com/plainkit/html"

// CloudCheck creates a Cloud Check Lucide icon.
func CloudCheck(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-cloud-check", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m17 15-5.5 5.5L9 18"))),
		x.Child(x.Path(x.D("M5 17.743A7 7 0 1 1 15.71 10h1.79a4.5 4.5 0 0 1 1.5 8.742"))),
	)
	return x.Svg(svgArgs...)
}
