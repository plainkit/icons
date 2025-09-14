package lucide

import x "github.com/bloxui/blox"

// TrendingUpDown creates a Trending Up Down Lucide icon.
func TrendingUpDown(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-trending-up-down", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M14.828 14.828 21 21"))),
		x.Child(x.Path(x.D("M21 16v5h-5"))),
		x.Child(x.Path(x.D("m21 3-9 9-4-4-6 6"))),
		x.Child(x.Path(x.D("M21 8V3h-5"))),
	)
	return x.Svg(svgArgs...)
}
