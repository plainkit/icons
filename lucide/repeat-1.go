package lucide

import x "github.com/bloxui/blox"

// Repeat1 creates a Repeat 1 Lucide icon.
func Repeat1(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-repeat-1", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m17 2 4 4-4 4"))),
		x.Child(x.Path(x.D("M3 11v-1a4 4 0 0 1 4-4h14"))),
		x.Child(x.Path(x.D("m7 22-4-4 4-4"))),
		x.Child(x.Path(x.D("M21 13v1a4 4 0 0 1-4 4H3"))),
		x.Child(x.Path(x.D("M11 10h1v4"))),
	)
	return x.Svg(svgArgs...)
}
