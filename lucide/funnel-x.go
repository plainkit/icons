package lucide

import x "github.com/plainkit/blox"

// FunnelX creates a Funnel X Lucide icon.
func FunnelX(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-funnel-x", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12.531 3H3a1 1 0 0 0-.742 1.67l7.225 7.989A2 2 0 0 1 10 14v6a1 1 0 0 0 .553.895l2 1A1 1 0 0 0 14 21v-7a2 2 0 0 1 .517-1.341l.427-.473"))),
		x.Child(x.Path(x.D("m16.5 3.5 5 5"))),
		x.Child(x.Path(x.D("m21.5 3.5-5 5"))),
	)
	return x.Svg(svgArgs...)
}
