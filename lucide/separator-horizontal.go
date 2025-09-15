package lucide

import x "github.com/plainkit/blox"

// SeparatorHorizontal creates a Separator Horizontal Lucide icon.
func SeparatorHorizontal(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-separator-horizontal", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m16 16-4 4-4-4"))),
		x.Child(x.Path(x.D("M3 12h18"))),
		x.Child(x.Path(x.D("m8 8 4-4 4 4"))),
	)
	return x.Svg(svgArgs...)
}
