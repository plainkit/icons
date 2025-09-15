package lucide

import x "github.com/plainkit/blox"

// ListPlus creates a List Plus Lucide icon.
func ListPlus(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-list-plus", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M16 5H3"))),
		x.Child(x.Path(x.D("M11 12H3"))),
		x.Child(x.Path(x.D("M16 19H3"))),
		x.Child(x.Path(x.D("M18 9v6"))),
		x.Child(x.Path(x.D("M21 12h-6"))),
	)
	return x.Svg(svgArgs...)
}
