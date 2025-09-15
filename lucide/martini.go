package lucide

import x "github.com/plainkit/html"

// Martini creates a Martini Lucide icon.
func Martini(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-martini", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M8 22h8"))),
		x.Child(x.Path(x.D("M12 11v11"))),
		x.Child(x.Path(x.D("m19 3-7 8-7-8Z"))),
	)
	return x.Svg(svgArgs...)
}
