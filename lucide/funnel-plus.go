package lucide

import x "github.com/plainkit/blox"

// FunnelPlus creates a Funnel Plus Lucide icon.
func FunnelPlus(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-funnel-plus", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M13.354 3H3a1 1 0 0 0-.742 1.67l7.225 7.989A2 2 0 0 1 10 14v6a1 1 0 0 0 .553.895l2 1A1 1 0 0 0 14 21v-7a2 2 0 0 1 .517-1.341l1.218-1.348"))),
		x.Child(x.Path(x.D("M16 6h6"))),
		x.Child(x.Path(x.D("M19 3v6"))),
	)
	return x.Svg(svgArgs...)
}
