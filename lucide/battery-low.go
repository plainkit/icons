package lucide

import (
	html "github.com/plainkit/html"
)

// BatteryLow creates a Battery Low Lucide icon.
func BatteryLow(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-battery-low", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M22 14v-4")),
		html.SvgPath(html.AD("M6 14v-4")),
		html.SvgRect(html.AWidth("16"), html.AHeight("12"), html.AX("2"), html.AY("6"), html.ARx("2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
