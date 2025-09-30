package lucide

import (
	html "github.com/plainkit/html"
)

// CloudAlert creates a Cloud Alert Lucide icon.
func CloudAlert(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-cloud-alert", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12 12v4")),
		html.SvgPath(html.AD("M12 20h.01")),
		html.SvgPath(html.AD("M17 18h.5a1 1 0 0 0 0-9h-1.79A7 7 0 1 0 7 17.708")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
