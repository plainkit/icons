package lucide

import x "github.com/bloxui/blox"

// ChevronsDownUp creates a Chevrons Down Up Lucide icon.
func ChevronsDownUp(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-chevrons-down-up", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m7 20 5-5 5 5"))),
		x.Child(x.Path(x.D("m7 4 5 5 5-5"))),
	)
	return x.Svg(svgArgs...)
}
