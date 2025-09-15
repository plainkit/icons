package lucide

import x "github.com/plainkit/blox"

// TrendingDown creates a Trending Down Lucide icon.
func TrendingDown(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-trending-down", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M16 17h6v-6"))),
		x.Child(x.Path(x.D("m22 17-8.5-8.5-5 5L2 7"))),
	)
	return x.Svg(svgArgs...)
}
