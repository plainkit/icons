package lucide

import x "github.com/bloxui/blox"

// ChevronDown creates a Chevron Down Lucide icon.
func ChevronDown(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-chevron-down", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m6 9 6 6 6-6"))),
	)
	return x.Svg(svgArgs...)
}
