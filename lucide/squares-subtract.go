package lucide

import x "github.com/bloxui/blox"

// SquaresSubtract creates a Squares Subtract Lucide icon.
func SquaresSubtract(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-squares-subtract", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10 22a2 2 0 0 1-2-2"))),
		x.Child(x.Path(x.D("M16 22h-2"))),
		x.Child(x.Path(x.D("M16 4a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h3a1 1 0 0 0 1-1v-5a2 2 0 0 1 2-2h5a1 1 0 0 0 1-1z"))),
		x.Child(x.Path(x.D("M20 8a2 2 0 0 1 2 2"))),
		x.Child(x.Path(x.D("M22 14v2"))),
		x.Child(x.Path(x.D("M22 20a2 2 0 0 1-2 2"))),
	)
	return x.Svg(svgArgs...)
}
