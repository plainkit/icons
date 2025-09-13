package lucide

import x "github.com/bloxui/blox"

// ChevronLast creates a Chevron Last Lucide icon.
func ChevronLast(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-chevron-last", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m7 18 6-6-6-6"))),
		x.Child(x.Path(x.D("M17 6v12"))),
	)
	return x.Svg(svgArgs...)
}
