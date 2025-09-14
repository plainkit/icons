package lucide

import x "github.com/bloxui/blox"

// Parentheses creates a Parentheses Lucide icon.
func Parentheses(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-parentheses", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M8 21s-4-3-4-9 4-9 4-9"))),
		x.Child(x.Path(x.D("M16 3s4 3 4 9-4 9-4 9"))),
	)
	return x.Svg(svgArgs...)
}
