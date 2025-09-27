package lucide

import (
	html "github.com/plainkit/html"
)

// MailOpen creates a Mail Open Lucide icon.
func MailOpen(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-mail-open", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M21.2 8.4c.5.38.8.97.8 1.6v10a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V10a2 2 0 0 1 .8-1.6l8-6a2 2 0 0 1 2.4 0l8 6Z"))),
		html.Child(html.SvgPath(html.AD("m22 10-8.97 5.7a1.94 1.94 0 0 1-2.06 0L2 10"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
