package lucide

import (
	html "github.com/plainkit/html"
)

// BanknoteX creates a Banknote X Lucide icon.
func BanknoteX(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-banknote-x", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M13 18H4a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v5")),
		html.SvgPath(html.AD("m17 17 5 5")),
		html.SvgPath(html.AD("M18 12h.01")),
		html.SvgPath(html.AD("m22 17-5 5")),
		html.SvgPath(html.AD("M6 12h.01")),
		html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
