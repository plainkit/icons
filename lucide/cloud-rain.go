package lucide

import (
	html "github.com/plainkit/html"
)

// CloudRain creates a Cloud Rain Lucide icon.
func CloudRain(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-cloud-rain", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M4 14.899A7 7 0 1 1 15.71 8h1.79a4.5 4.5 0 0 1 2.5 8.242")),
		html.SvgPath(html.AD("M16 14v6")),
		html.SvgPath(html.AD("M8 14v6")),
		html.SvgPath(html.AD("M12 16v6")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
