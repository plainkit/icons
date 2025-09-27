package lucide

import (
	html "github.com/plainkit/html"
)

// SendHorizontal creates a Send Horizontal Lucide icon.
func SendHorizontal(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-send-horizontal", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M3.714 3.048a.498.498 0 0 0-.683.627l2.843 7.627a2 2 0 0 1 0 1.396l-2.842 7.627a.498.498 0 0 0 .682.627l18-8.5a.5.5 0 0 0 0-.904z"))),
		html.Child(html.SvgPath(html.AD("M6 12h16"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
