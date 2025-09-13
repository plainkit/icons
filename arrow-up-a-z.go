package lucide

import x "github.com/bloxui/blox"

// ArrowUpAZ creates a Arrow Up A Z Lucide icon.
func ArrowUpAZ(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-arrow-up-a-z", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m3 8 4-4 4 4"))),
		x.Child(x.Path(x.D("M7 4v16"))),
		x.Child(x.Path(x.D("M20 8h-5"))),
		x.Child(x.Path(x.D("M15 10V6.5a2.5 2.5 0 0 1 5 0V10"))),
		x.Child(x.Path(x.D("M15 14h5l-5 6h5"))),
	)
	return x.Svg(svgArgs...)
}
