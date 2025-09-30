package lucide

import (
	html "github.com/plainkit/html"
)

// PlugZap creates a Plug Zap Lucide icon.
func PlugZap(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-plug-zap", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M6.3 20.3a2.4 2.4 0 0 0 3.4 0L12 18l-6-6-2.3 2.3a2.4 2.4 0 0 0 0 3.4Z")),
		html.SvgPath(html.AD("m2 22 3-3")),
		html.SvgPath(html.AD("M7.5 13.5 10 11")),
		html.SvgPath(html.AD("M10.5 16.5 13 14")),
		html.SvgPath(html.AD("m18 3-4 4h6l-4 4")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
