package lucide

import (
	html "github.com/plainkit/html"
)

// ShieldBan creates a Shield Ban Lucide icon.
func ShieldBan(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-shield-ban", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M20 13c0 5-3.5 7.5-7.66 8.95a1 1 0 0 1-.67-.01C7.5 20.5 4 18 4 13V6a1 1 0 0 1 1-1c2 0 4.5-1.2 6.24-2.72a1.17 1.17 0 0 1 1.52 0C14.51 3.81 17 5 19 5a1 1 0 0 1 1 1z")),
		html.SvgPath(html.AD("m4.243 5.21 14.39 12.472")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
