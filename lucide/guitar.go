package lucide

import (
	html "github.com/plainkit/html"
)

// Guitar creates a Guitar Lucide icon.
func Guitar(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-guitar", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m11.9 12.1 4.514-4.514")),
		html.SvgPath(html.AD("M20.1 2.3a1 1 0 0 0-1.4 0l-1.114 1.114A2 2 0 0 0 17 4.828v1.344a2 2 0 0 1-.586 1.414A2 2 0 0 1 17.828 7h1.344a2 2 0 0 0 1.414-.586L21.7 5.3a1 1 0 0 0 0-1.4z")),
		html.SvgPath(html.AD("m6 16 2 2")),
		html.SvgPath(html.AD("M8.23 9.85A3 3 0 0 1 11 8a5 5 0 0 1 5 5 3 3 0 0 1-1.85 2.77l-.92.38A2 2 0 0 0 12 18a4 4 0 0 1-4 4 6 6 0 0 1-6-6 4 4 0 0 1 4-4 2 2 0 0 0 1.85-1.23z")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
