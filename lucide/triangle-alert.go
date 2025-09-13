package lucide

import x "github.com/bloxui/blox"

// TriangleAlert creates a Triangle Alert Lucide icon.
func TriangleAlert(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-triangle-alert", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m21.73 18-8-14a2 2 0 0 0-3.48 0l-8 14A2 2 0 0 0 4 21h16a2 2 0 0 0 1.73-3"))),
		x.Child(x.Path(x.D("M12 9v4"))),
		x.Child(x.Path(x.D("M12 17h.01"))),
	)
	return x.Svg(svgArgs...)
}
