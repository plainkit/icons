package lucide

import x "github.com/bloxui/blox"

// Repeat creates a Repeat Lucide icon.
func Repeat(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-repeat", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m17 2 4 4-4 4"))),
		x.Child(x.Path(x.D("M3 11v-1a4 4 0 0 1 4-4h14"))),
		x.Child(x.Path(x.D("m7 22-4-4 4-4"))),
		x.Child(x.Path(x.D("M21 13v1a4 4 0 0 1-4 4H3"))),
	)
	return x.Svg(svgArgs...)
}
