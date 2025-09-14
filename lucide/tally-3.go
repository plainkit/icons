package lucide

import x "github.com/bloxui/blox"

// Tally3 creates a Tally 3 Lucide icon.
func Tally3(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-tally-3", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M4 4v16"))),
		x.Child(x.Path(x.D("M9 4v16"))),
		x.Child(x.Path(x.D("M14 4v16"))),
	)
	return x.Svg(svgArgs...)
}
