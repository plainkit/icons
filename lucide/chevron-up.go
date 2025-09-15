package lucide

import x "github.com/plainkit/blox"

// ChevronUp creates a Chevron Up Lucide icon.
func ChevronUp(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-chevron-up", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m18 15-6-6-6 6"))),
	)
	return x.Svg(svgArgs...)
}
