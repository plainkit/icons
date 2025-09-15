package lucide

import x "github.com/plainkit/html"

// Bird creates a Bird Lucide icon.
func Bird(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-bird", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M16 7h.01"))),
		x.Child(x.Path(x.D("M3.4 18H12a8 8 0 0 0 8-8V7a4 4 0 0 0-7.28-2.3L2 20"))),
		x.Child(x.Path(x.D("m20 7 2 .5-2 .5"))),
		x.Child(x.Path(x.D("M10 18v3"))),
		x.Child(x.Path(x.D("M14 17.75V21"))),
		x.Child(x.Path(x.D("M7 18a6 6 0 0 0 3.84-10.61"))),
	)
	return x.Svg(svgArgs...)
}
