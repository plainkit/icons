package lucide

import (
	html "github.com/plainkit/html"
)

// BatteryMedium creates a Battery Medium Lucide icon.
func BatteryMedium(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-battery-medium", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M10 14v-4"))),
		html.Child(html.SvgPath(html.AD("M22 14v-4"))),
		html.Child(html.SvgPath(html.AD("M6 14v-4"))),
		html.Child(html.SvgRect(html.AWidth("16"), html.AHeight("12"), html.AX("2"), html.AY("6"), html.ARx("2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
