package lucide

import (
	html "github.com/plainkit/html"
)

// MailSearch creates a Mail Search Lucide icon.
func MailSearch(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-mail-search", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M22 12.5V6a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v12c0 1.1.9 2 2 2h7.5"))),
		html.Child(html.SvgPath(html.AD("m22 7-8.97 5.7a1.94 1.94 0 0 1-2.06 0L2 7"))),
		html.Child(html.SvgPath(html.AD("M18 21a3 3 0 1 0 0-6 3 3 0 0 0 0 6Z"))),
		html.Child(html.SvgCircle(html.ACx("18"), html.ACy("18"), html.AR("3"))),
		html.Child(html.SvgPath(html.AD("m22 22-1.5-1.5"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
