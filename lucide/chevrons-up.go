package lucide

import x "github.com/bloxui/blox"

// ChevronsUp creates a Chevrons Up Lucide icon.
func ChevronsUp(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-chevrons-up", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m17 11-5-5-5 5"))),
		x.Child(x.Path(x.D("m17 18-5-5-5 5"))),
	)
	return x.Svg(svgArgs...)
}
