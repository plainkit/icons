package lucide

import x "github.com/bloxui/blox"

// Cloudy creates a Cloudy Lucide icon.
func Cloudy(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-cloudy", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M17.5 21H9a7 7 0 1 1 6.71-9h1.79a4.5 4.5 0 1 1 0 9Z"))),
		x.Child(x.Path(x.D("M22 10a3 3 0 0 0-3-3h-2.207a5.502 5.502 0 0 0-10.702.5"))),
	)
	return x.Svg(svgArgs...)
}
