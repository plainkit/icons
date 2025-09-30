package lucide

import (
	html "github.com/plainkit/html"
)

// Croissant creates a Croissant Lucide icon.
func Croissant(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-croissant", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M10.2 18H4.774a1.5 1.5 0 0 1-1.352-.97 11 11 0 0 1 .132-6.487")),
		html.SvgPath(html.AD("M18 10.2V4.774a1.5 1.5 0 0 0-.97-1.352 11 11 0 0 0-6.486.132")),
		html.SvgPath(html.AD("M18 5a4 3 0 0 1 4 3 2 2 0 0 1-2 2 10 10 0 0 0-5.139 1.42")),
		html.SvgPath(html.AD("M5 18a3 4 0 0 0 3 4 2 2 0 0 0 2-2 10 10 0 0 1 1.42-5.14")),
		html.SvgPath(html.AD("M8.709 2.554a10 10 0 0 0-6.155 6.155 1.5 1.5 0 0 0 .676 1.626l9.807 5.42a2 2 0 0 0 2.718-2.718l-5.42-9.807a1.5 1.5 0 0 0-1.626-.676")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
