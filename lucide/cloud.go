package lucide

import x "github.com/plainkit/html"

// Cloud creates a Cloud Lucide icon.
func Cloud(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-cloud", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M17.5 19H9a7 7 0 1 1 6.71-9h1.79a4.5 4.5 0 1 1 0 9Z"))),
	)
	return x.Svg(svgArgs...)
}
