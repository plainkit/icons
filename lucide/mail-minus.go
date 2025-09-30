package lucide

import (
	html "github.com/plainkit/html"
)

// MailMinus creates a Mail Minus Lucide icon.
func MailMinus(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-mail-minus", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M22 15V6a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v12c0 1.1.9 2 2 2h8")),
		html.SvgPath(html.AD("m22 7-8.97 5.7a1.94 1.94 0 0 1-2.06 0L2 7")),
		html.SvgPath(html.AD("M16 19h6")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
