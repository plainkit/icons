package lucide

import (
	html "github.com/plainkit/html"
)

// Plug2 creates a Plug 2 Lucide icon.
func Plug2(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-plug-2", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M9 2v6")),
		html.SvgPath(html.AD("M15 2v6")),
		html.SvgPath(html.AD("M12 17v5")),
		html.SvgPath(html.AD("M5 8h14")),
		html.SvgPath(html.AD("M6 11V8h12v3a6 6 0 1 1-12 0Z")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
