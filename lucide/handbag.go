package lucide

import x "github.com/bloxui/blox"

// Handbag creates a Handbag Lucide icon.
func Handbag(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-handbag", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M2.048 18.566A2 2 0 0 0 4 21h16a2 2 0 0 0 1.952-2.434l-2-9A2 2 0 0 0 18 8H6a2 2 0 0 0-1.952 1.566z"))),
		x.Child(x.Path(x.D("M8 11V6a4 4 0 0 1 8 0v5"))),
	)
	return x.Svg(svgArgs...)
}
