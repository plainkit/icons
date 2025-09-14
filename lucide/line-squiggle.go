package lucide

import x "github.com/bloxui/blox"

// LineSquiggle creates a Line Squiggle Lucide icon.
func LineSquiggle(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-line-squiggle", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M7 3.5c5-2 7 2.5 3 4C1.5 10 2 15 5 16c5 2 9-10 14-7s.5 13.5-4 12c-5-2.5.5-11 6-2"))),
	)
	return x.Svg(svgArgs...)
}
