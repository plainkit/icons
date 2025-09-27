package lucide

import (
	html "github.com/plainkit/html"
)

// KeySquare creates a Key Square Lucide icon.
func KeySquare(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-key-square", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M12.4 2.7a2.5 2.5 0 0 1 3.4 0l5.5 5.5a2.5 2.5 0 0 1 0 3.4l-3.7 3.7a2.5 2.5 0 0 1-3.4 0L8.7 9.8a2.5 2.5 0 0 1 0-3.4z"))),
		html.Child(html.SvgPath(html.AD("m14 7 3 3"))),
		html.Child(html.SvgPath(html.AD("m9.4 10.6-6.814 6.814A2 2 0 0 0 2 18.828V21a1 1 0 0 0 1 1h3a1 1 0 0 0 1-1v-1a1 1 0 0 1 1-1h1a1 1 0 0 0 1-1v-1a1 1 0 0 1 1-1h.172a2 2 0 0 0 1.414-.586l.814-.814"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
