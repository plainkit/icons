package lucide

import (
	html "github.com/plainkit/html"
)

// Reply creates a Reply Lucide icon.
func Reply(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-reply", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M20 18v-2a4 4 0 0 0-4-4H4"))),
		html.Child(html.SvgPath(html.AD("m9 17-5-5 5-5"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
