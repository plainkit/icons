package lucide

import (
	html "github.com/plainkit/html"
)

// Fence creates a Fence Lucide icon.
func Fence(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-fence", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M4 3 2 5v15c0 .6.4 1 1 1h2c.6 0 1-.4 1-1V5Z"))),
		html.Child(html.SvgPath(html.AD("M6 8h4"))),
		html.Child(html.SvgPath(html.AD("M6 18h4"))),
		html.Child(html.SvgPath(html.AD("m12 3-2 2v15c0 .6.4 1 1 1h2c.6 0 1-.4 1-1V5Z"))),
		html.Child(html.SvgPath(html.AD("M14 8h4"))),
		html.Child(html.SvgPath(html.AD("M14 18h4"))),
		html.Child(html.SvgPath(html.AD("m20 3-2 2v15c0 .6.4 1 1 1h2c.6 0 1-.4 1-1V5Z"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
