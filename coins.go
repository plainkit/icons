package lucide

import x "github.com/bloxui/blox"

// Coins creates a Coins Lucide icon.
func Coins(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-coins", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("8"), x.Cy("8"), x.R("6"))),
		x.Child(x.Path(x.D("M18.09 10.37A6 6 0 1 1 10.34 18"))),
		x.Child(x.Path(x.D("M7 6h1v4"))),
		x.Child(x.Path(x.D("m16.71 13.88.7.71-2.82 2.82"))),
	)
	return x.Svg(svgArgs...)
}