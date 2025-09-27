package lucide

import (
	html "github.com/plainkit/html"
)

// MessageSquareLock creates a Message Square Lock Lucide icon.
func MessageSquareLock(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-message-square-lock", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M22 8.5V5a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v16.286a.71.71 0 0 0 1.212.502l2.202-2.202A2 2 0 0 1 6.828 19H10"))),
		html.Child(html.SvgPath(html.AD("M20 15v-2a2 2 0 0 0-4 0v2"))),
		html.Child(html.SvgRect(html.AWidth("8"), html.AHeight("5"), html.AX("14"), html.AY("15"), html.ARx("1"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
