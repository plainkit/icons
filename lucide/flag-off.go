package lucide

import x "github.com/bloxui/blox"

// FlagOff creates a Flag Off Lucide icon.
func FlagOff(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-flag-off", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M16 16c-3 0-5-2-8-2a6 6 0 0 0-4 1.528"))),
		x.Child(x.Path(x.D("m2 2 20 20"))),
		x.Child(x.Path(x.D("M4 22V4"))),
		x.Child(x.Path(x.D("M7.656 2H8c3 0 5 2 7.333 2q2 0 3.067-.8A1 1 0 0 1 20 4v10.347"))),
	)
	return x.Svg(svgArgs...)
}
