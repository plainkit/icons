package lucide

import x "github.com/bloxui/blox"

// Bath creates a Bath Lucide icon.
func Bath(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-bath", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10 4 8 6"))),
		x.Child(x.Path(x.D("M17 19v2"))),
		x.Child(x.Path(x.D("M2 12h20"))),
		x.Child(x.Path(x.D("M7 19v2"))),
		x.Child(x.Path(x.D("M9 5 7.621 3.621A2.121 2.121 0 0 0 4 5v12a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-5"))),
	)
	return x.Svg(svgArgs...)
}
