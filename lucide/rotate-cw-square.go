package lucide

import x "github.com/bloxui/blox"

// RotateCwSquare creates a Rotate Cw Square Lucide icon.
func RotateCwSquare(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-rotate-cw-square", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 5H6a2 2 0 0 0-2 2v3"))),
		x.Child(x.Path(x.D("m9 8 3-3-3-3"))),
		x.Child(x.Path(x.D("M4 14v4a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2h-2"))),
	)
	return x.Svg(svgArgs...)
}
