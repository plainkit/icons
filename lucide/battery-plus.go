package lucide

import x "github.com/plainkit/html"

// BatteryPlus creates a Battery Plus Lucide icon.
func BatteryPlus(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-battery-plus", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10 9v6"))),
		x.Child(x.Path(x.D("M12.543 6H16a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2h-3.605"))),
		x.Child(x.Path(x.D("M22 14v-4"))),
		x.Child(x.Path(x.D("M7 12h6"))),
		x.Child(x.Path(x.D("M7.606 18H4a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h3.606"))),
	)
	return x.Svg(svgArgs...)
}
