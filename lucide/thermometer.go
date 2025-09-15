package lucide

import x "github.com/plainkit/html"

// Thermometer creates a Thermometer Lucide icon.
func Thermometer(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-thermometer", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M14 4v10.54a4 4 0 1 1-4 0V4a2 2 0 0 1 4 0Z"))),
	)
	return x.Svg(svgArgs...)
}
