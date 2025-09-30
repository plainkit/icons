package lucide

import (
	html "github.com/plainkit/html"
)

// BatteryCharging creates a Battery Charging Lucide icon.
func BatteryCharging(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-battery-charging", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m11 7-3 5h4l-3 5")),
		html.SvgPath(html.AD("M14.856 6H16a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2h-2.935")),
		html.SvgPath(html.AD("M22 14v-4")),
		html.SvgPath(html.AD("M5.14 18H4a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h2.936")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
