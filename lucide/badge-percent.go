package lucide

import x "github.com/bloxui/blox"

// BadgePercent creates a Badge Percent Lucide icon.
func BadgePercent(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-badge-percent", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M3.85 8.62a4 4 0 0 1 4.78-4.77 4 4 0 0 1 6.74 0 4 4 0 0 1 4.78 4.78 4 4 0 0 1 0 6.74 4 4 0 0 1-4.77 4.78 4 4 0 0 1-6.75 0 4 4 0 0 1-4.78-4.77 4 4 0 0 1 0-6.76Z"))),
		x.Child(x.Path(x.D("m15 9-6 6"))),
		x.Child(x.Path(x.D("M9 9h.01"))),
		x.Child(x.Path(x.D("M15 15h.01"))),
	)
	return x.Svg(svgArgs...)
}
