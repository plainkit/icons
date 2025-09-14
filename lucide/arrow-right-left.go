package lucide

import x "github.com/bloxui/blox"

// ArrowRightLeft creates a Arrow Right Left Lucide icon.
func ArrowRightLeft(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-arrow-right-left", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m16 3 4 4-4 4"))),
		x.Child(x.Path(x.D("M20 7H4"))),
		x.Child(x.Path(x.D("m8 21-4-4 4-4"))),
		x.Child(x.Path(x.D("M4 17h16"))),
	)
	return x.Svg(svgArgs...)
}
