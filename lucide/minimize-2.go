package lucide

import x "github.com/plainkit/html"

// Minimize2 creates a Minimize 2 Lucide icon.
func Minimize2(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-minimize-2", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m14 10 7-7"))),
		x.Child(x.Path(x.D("M20 10h-6V4"))),
		x.Child(x.Path(x.D("m3 21 7-7"))),
		x.Child(x.Path(x.D("M4 14h6v6"))),
	)
	return x.Svg(svgArgs...)
}
