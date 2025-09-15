package lucide

import x "github.com/plainkit/html"

// Laptop creates a Laptop Lucide icon.
func Laptop(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-laptop", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M18 5a2 2 0 0 1 2 2v8.526a2 2 0 0 0 .212.897l1.068 2.127a1 1 0 0 1-.9 1.45H3.62a1 1 0 0 1-.9-1.45l1.068-2.127A2 2 0 0 0 4 15.526V7a2 2 0 0 1 2-2z"))),
		x.Child(x.Path(x.D("M20.054 15.987H3.946"))),
	)
	return x.Svg(svgArgs...)
}
