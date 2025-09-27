package lucide

import (
	html "github.com/plainkit/html"
)

// Share creates a Share Lucide icon.
func Share(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-share", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M12 2v13"))),
		html.Child(html.SvgPath(html.AD("m16 6-4-4-4 4"))),
		html.Child(html.SvgPath(html.AD("M4 12v8a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-8"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
