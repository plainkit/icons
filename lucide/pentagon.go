package lucide

import x "github.com/bloxui/blox"

// Pentagon creates a Pentagon Lucide icon.
func Pentagon(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-pentagon", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10.83 2.38a2 2 0 0 1 2.34 0l8 5.74a2 2 0 0 1 .73 2.25l-3.04 9.26a2 2 0 0 1-1.9 1.37H7.04a2 2 0 0 1-1.9-1.37L2.1 10.37a2 2 0 0 1 .73-2.25z"))),
	)
	return x.Svg(svgArgs...)
}
