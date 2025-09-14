package lucide

import x "github.com/bloxui/blox"

// TrendingUp creates a Trending Up Lucide icon.
func TrendingUp(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-trending-up", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M16 7h6v6"))),
		x.Child(x.Path(x.D("m22 7-8.5 8.5-5-5L2 17"))),
	)
	return x.Svg(svgArgs...)
}
