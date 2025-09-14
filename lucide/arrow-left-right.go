package lucide

import x "github.com/bloxui/blox"

// ArrowLeftRight creates a Arrow Left Right Lucide icon.
func ArrowLeftRight(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-arrow-left-right", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M8 3 4 7l4 4"))),
		x.Child(x.Path(x.D("M4 7h16"))),
		x.Child(x.Path(x.D("m16 21 4-4-4-4"))),
		x.Child(x.Path(x.D("M20 17H4"))),
	)
	return x.Svg(svgArgs...)
}
