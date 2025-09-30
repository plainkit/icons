package lucide

import (
	html "github.com/plainkit/html"
)

// Hospital creates a Hospital Lucide icon.
func Hospital(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-hospital", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12 7v4")),
		html.SvgPath(html.AD("M14 21v-3a2 2 0 0 0-4 0v3")),
		html.SvgPath(html.AD("M14 9h-4")),
		html.SvgPath(html.AD("M18 11h2a2 2 0 0 1 2 2v6a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2v-9a2 2 0 0 1 2-2h2")),
		html.SvgPath(html.AD("M18 21V5a2 2 0 0 0-2-2H8a2 2 0 0 0-2 2v16")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
