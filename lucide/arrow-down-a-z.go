package lucide

import x "github.com/plainkit/blox"

// ArrowDownAZ creates a Arrow Down A Z Lucide icon.
func ArrowDownAZ(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-arrow-down-a-z", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m3 16 4 4 4-4"))),
		x.Child(x.Path(x.D("M7 20V4"))),
		x.Child(x.Path(x.D("M20 8h-5"))),
		x.Child(x.Path(x.D("M15 10V6.5a2.5 2.5 0 0 1 5 0V10"))),
		x.Child(x.Path(x.D("M15 14h5l-5 6h5"))),
	)
	return x.Svg(svgArgs...)
}
