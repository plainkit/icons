package lucide

import x "github.com/plainkit/blox"

// Headphones creates a Headphones Lucide icon.
func Headphones(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-headphones", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M3 14h3a2 2 0 0 1 2 2v3a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-7a9 9 0 0 1 18 0v7a2 2 0 0 1-2 2h-1a2 2 0 0 1-2-2v-3a2 2 0 0 1 2-2h3"))),
	)
	return x.Svg(svgArgs...)
}
