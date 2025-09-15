package lucide

import x "github.com/plainkit/blox"

// BatteryWarning creates a Battery Warning Lucide icon.
func BatteryWarning(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-battery-warning", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10 17h.01"))),
		x.Child(x.Path(x.D("M10 7v6"))),
		x.Child(x.Path(x.D("M14 6h2a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2h-2"))),
		x.Child(x.Path(x.D("M22 14v-4"))),
		x.Child(x.Path(x.D("M6 18H4a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h2"))),
	)
	return x.Svg(svgArgs...)
}
