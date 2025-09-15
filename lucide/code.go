package lucide

import x "github.com/plainkit/html"

// Code creates a Code Lucide icon.
func Code(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-code", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m16 18 6-6-6-6"))),
		x.Child(x.Path(x.D("m8 6-6 6 6 6"))),
	)
	return x.Svg(svgArgs...)
}
