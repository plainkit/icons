package lucide

import x "github.com/bloxui/blox"

// Tally1 creates a Tally 1 Lucide icon.
func Tally1(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-tally-1", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M4 4v16"))),
	)
	return x.Svg(svgArgs...)
}
