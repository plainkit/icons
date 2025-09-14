package lucide

import x "github.com/bloxui/blox"

// MoveDiagonal2 creates a Move Diagonal 2 Lucide icon.
func MoveDiagonal2(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-move-diagonal-2", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M19 13v6h-6"))),
		x.Child(x.Path(x.D("M5 11V5h6"))),
		x.Child(x.Path(x.D("m5 5 14 14"))),
	)
	return x.Svg(svgArgs...)
}
