package lucide

import (
	html "github.com/plainkit/html"
)

// BanknoteArrowUp creates a Banknote Arrow Up Lucide icon.
func BanknoteArrowUp(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-banknote-arrow-up", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12 18H4a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v5")),
		html.SvgPath(html.AD("M18 12h.01")),
		html.SvgPath(html.AD("M19 22v-6")),
		html.SvgPath(html.AD("m22 19-3-3-3 3")),
		html.SvgPath(html.AD("M6 12h.01")),
		html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
