package lucide

import x "github.com/plainkit/blox"

// ArrowDownZA creates a Arrow Down Z A Lucide icon.
func ArrowDownZA(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-arrow-down-z-a", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m3 16 4 4 4-4"))),
		x.Child(x.Path(x.D("M7 4v16"))),
		x.Child(x.Path(x.D("M15 4h5l-5 6h5"))),
		x.Child(x.Path(x.D("M15 20v-3.5a2.5 2.5 0 0 1 5 0V20"))),
		x.Child(x.Path(x.D("M20 18h-5"))),
	)
	return x.Svg(svgArgs...)
}
