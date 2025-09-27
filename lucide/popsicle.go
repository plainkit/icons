package lucide

import (
	html "github.com/plainkit/html"
)

// Popsicle creates a Popsicle Lucide icon.
func Popsicle(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-popsicle", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M18.6 14.4c.8-.8.8-2 0-2.8l-8.1-8.1a4.95 4.95 0 1 0-7.1 7.1l8.1 8.1c.9.7 2.1.7 2.9-.1Z"))),
		html.Child(html.SvgPath(html.AD("m22 22-5.5-5.5"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
