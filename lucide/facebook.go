package lucide

import (
	html "github.com/plainkit/html"
)

// Facebook creates a Facebook Lucide icon.
func Facebook(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-facebook", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M18 2h-3a5 5 0 0 0-5 5v3H7v4h3v8h4v-8h3l1-4h-4V7a1 1 0 0 1 1-1h3z"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
