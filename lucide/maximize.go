package lucide

import x "github.com/plainkit/html"

// Maximize creates a Maximize Lucide icon.
func Maximize(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-maximize", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M8 3H5a2 2 0 0 0-2 2v3"))),
		x.Child(x.Path(x.D("M21 8V5a2 2 0 0 0-2-2h-3"))),
		x.Child(x.Path(x.D("M3 16v3a2 2 0 0 0 2 2h3"))),
		x.Child(x.Path(x.D("M16 21h3a2 2 0 0 0 2-2v-3"))),
	)
	return x.Svg(svgArgs...)
}
