package lucide

import x "github.com/bloxui/blox"

// Tornado creates a Tornado Lucide icon.
func Tornado(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-tornado", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M21 4H3"))),
		x.Child(x.Path(x.D("M18 8H6"))),
		x.Child(x.Path(x.D("M19 12H9"))),
		x.Child(x.Path(x.D("M16 16h-6"))),
		x.Child(x.Path(x.D("M11 20H9"))),
	)
	return x.Svg(svgArgs...)
}
