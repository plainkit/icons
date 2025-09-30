package lucide

import (
	html "github.com/plainkit/html"
)

// Spotlight creates a Spotlight Lucide icon.
func Spotlight(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-spotlight", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M15.295 19.562 16 22")),
		html.SvgPath(html.AD("m17 16 3.758 2.098")),
		html.SvgPath(html.AD("m19 12.5 3.026-.598")),
		html.SvgPath(html.AD("M7.61 6.3a3 3 0 0 0-3.92 1.3l-1.38 2.79a3 3 0 0 0 1.3 3.91l6.89 3.597a1 1 0 0 0 1.342-.447l3.106-6.211a1 1 0 0 0-.447-1.341z")),
		html.SvgPath(html.AD("M8 9V2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
