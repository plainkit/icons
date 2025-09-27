package lucide

import (
	html "github.com/plainkit/html"
)

// BatteryPlus creates a Battery Plus Lucide icon.
func BatteryPlus(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-battery-plus", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M10 9v6"))),
		html.Child(html.SvgPath(html.AD("M12.543 6H16a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2h-3.605"))),
		html.Child(html.SvgPath(html.AD("M22 14v-4"))),
		html.Child(html.SvgPath(html.AD("M7 12h6"))),
		html.Child(html.SvgPath(html.AD("M7.606 18H4a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h3.606"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
