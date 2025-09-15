package lucide

import x "github.com/plainkit/blox"

// BatteryCharging creates a Battery Charging Lucide icon.
func BatteryCharging(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-battery-charging", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m11 7-3 5h4l-3 5"))),
		x.Child(x.Path(x.D("M14.856 6H16a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2h-2.935"))),
		x.Child(x.Path(x.D("M22 14v-4"))),
		x.Child(x.Path(x.D("M5.14 18H4a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h2.936"))),
	)
	return x.Svg(svgArgs...)
}
