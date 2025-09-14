package lucide

import x "github.com/bloxui/blox"

// ListEnd creates a List End Lucide icon.
func ListEnd(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-list-end", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M16 5H3"))),
		x.Child(x.Path(x.D("M16 12H3"))),
		x.Child(x.Path(x.D("M9 19H3"))),
		x.Child(x.Path(x.D("m16 16-3 3 3 3"))),
		x.Child(x.Path(x.D("M21 5v12a2 2 0 0 1-2 2h-6"))),
	)
	return x.Svg(svgArgs...)
}
