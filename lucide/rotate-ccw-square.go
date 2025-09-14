package lucide

import x "github.com/bloxui/blox"

// RotateCcwSquare creates a Rotate Ccw Square Lucide icon.
func RotateCcwSquare(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-rotate-ccw-square", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M20 9V7a2 2 0 0 0-2-2h-6"))),
		x.Child(x.Path(x.D("m15 2-3 3 3 3"))),
		x.Child(x.Path(x.D("M20 13v5a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V7a2 2 0 0 1 2-2h2"))),
	)
	return x.Svg(svgArgs...)
}
