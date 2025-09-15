package lucide

import x "github.com/plainkit/blox"

// Lightbulb creates a Lightbulb Lucide icon.
func Lightbulb(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-lightbulb", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M15 14c.2-1 .7-1.7 1.5-2.5 1-.9 1.5-2.2 1.5-3.5A6 6 0 0 0 6 8c0 1 .2 2.2 1.5 3.5.7.7 1.3 1.5 1.5 2.5"))),
		x.Child(x.Path(x.D("M9 18h6"))),
		x.Child(x.Path(x.D("M10 22h4"))),
	)
	return x.Svg(svgArgs...)
}
