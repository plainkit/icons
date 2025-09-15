package lucide

import x "github.com/plainkit/blox"

// Funnel creates a Funnel Lucide icon.
func Funnel(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-funnel", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10 20a1 1 0 0 0 .553.895l2 1A1 1 0 0 0 14 21v-7a2 2 0 0 1 .517-1.341L21.74 4.67A1 1 0 0 0 21 3H3a1 1 0 0 0-.742 1.67l7.225 7.989A2 2 0 0 1 10 14z"))),
	)
	return x.Svg(svgArgs...)
}
