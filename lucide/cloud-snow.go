package lucide

import (
	html "github.com/plainkit/html"
)

// CloudSnow creates a Cloud Snow Lucide icon.
func CloudSnow(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-cloud-snow", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M4 14.899A7 7 0 1 1 15.71 8h1.79a4.5 4.5 0 0 1 2.5 8.242")),
		html.SvgPath(html.AD("M8 15h.01")),
		html.SvgPath(html.AD("M8 19h.01")),
		html.SvgPath(html.AD("M12 17h.01")),
		html.SvgPath(html.AD("M12 21h.01")),
		html.SvgPath(html.AD("M16 15h.01")),
		html.SvgPath(html.AD("M16 19h.01")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
