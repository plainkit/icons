package lucide

import x "github.com/plainkit/html"

// Building2 creates a Building 2 Lucide icon.
func Building2(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-building-2", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M6 22V4a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2v18Z"))),
		x.Child(x.Path(x.D("M6 12H4a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h2"))),
		x.Child(x.Path(x.D("M18 9h2a2 2 0 0 1 2 2v9a2 2 0 0 1-2 2h-2"))),
		x.Child(x.Path(x.D("M10 6h4"))),
		x.Child(x.Path(x.D("M10 10h4"))),
		x.Child(x.Path(x.D("M10 14h4"))),
		x.Child(x.Path(x.D("M10 18h4"))),
	)
	return x.Svg(svgArgs...)
}
