package lucide

import x "github.com/bloxui/blox"

// ListOrdered creates a List Ordered Lucide icon.
func ListOrdered(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-list-ordered", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M11 5h10"))),
		x.Child(x.Path(x.D("M11 12h10"))),
		x.Child(x.Path(x.D("M11 19h10"))),
		x.Child(x.Path(x.D("M4 4h1v5"))),
		x.Child(x.Path(x.D("M4 9h2"))),
		x.Child(x.Path(x.D("M6.5 20H3.4c0-1 2.6-1.925 2.6-3.5a1.5 1.5 0 0 0-2.6-1.02"))),
	)
	return x.Svg(svgArgs...)
}
