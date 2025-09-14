package lucide

import x "github.com/bloxui/blox"

// Minus creates a Minus Lucide icon.
func Minus(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-minus", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M5 12h14"))),
	)
	return x.Svg(svgArgs...)
}
