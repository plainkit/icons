package lucide

import (
	html "github.com/plainkit/html"
)

// ReplyAll creates a Reply All Lucide icon.
func ReplyAll(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-reply-all", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m12 17-5-5 5-5"))),
		html.Child(html.SvgPath(html.AD("M22 18v-2a4 4 0 0 0-4-4H7"))),
		html.Child(html.SvgPath(html.AD("m7 17-5-5 5-5"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
