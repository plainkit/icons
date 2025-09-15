package lucide

import x "github.com/plainkit/html"

// Dam creates a Dam Lucide icon.
func Dam(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-dam", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M11 11.31c1.17.56 1.54 1.69 3.5 1.69 2.5 0 2.5-2 5-2 1.3 0 1.9.5 2.5 1"))),
		x.Child(x.Path(x.D("M11.75 18c.35.5 1.45 1 2.75 1 2.5 0 2.5-2 5-2 1.3 0 1.9.5 2.5 1"))),
		x.Child(x.Path(x.D("M2 10h4"))),
		x.Child(x.Path(x.D("M2 14h4"))),
		x.Child(x.Path(x.D("M2 18h4"))),
		x.Child(x.Path(x.D("M2 6h4"))),
		x.Child(x.Path(x.D("M7 3a1 1 0 0 0-1 1v16a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1L10 4a1 1 0 0 0-1-1z"))),
	)
	return x.Svg(svgArgs...)
}
