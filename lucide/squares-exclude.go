package lucide

import x "github.com/bloxui/blox"

// SquaresExclude creates a Squares Exclude Lucide icon.
func SquaresExclude(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-squares-exclude", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M16 12v2a2 2 0 0 1-2 2H9a1 1 0 0 0-1 1v3a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V10a2 2 0 0 0-2-2h0"))),
		x.Child(x.Path(x.D("M4 16a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v3a1 1 0 0 1-1 1h-5a2 2 0 0 0-2 2v2"))),
	)
	return x.Svg(svgArgs...)
}
