package lucide

import (
	html "github.com/plainkit/html"
)

// MailCheck creates a Mail Check Lucide icon.
func MailCheck(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-mail-check", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M22 13V6a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v12c0 1.1.9 2 2 2h8"))),
		html.Child(html.SvgPath(html.AD("m22 7-8.97 5.7a1.94 1.94 0 0 1-2.06 0L2 7"))),
		html.Child(html.SvgPath(html.AD("m16 19 2 2 4-4"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
