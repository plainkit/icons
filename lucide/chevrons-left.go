package lucide

import x "github.com/plainkit/blox"

// ChevronsLeft creates a Chevrons Left Lucide icon.
func ChevronsLeft(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-chevrons-left", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m11 17-5-5 5-5"))),
		x.Child(x.Path(x.D("m18 17-5-5 5-5"))),
	)
	return x.Svg(svgArgs...)
}
