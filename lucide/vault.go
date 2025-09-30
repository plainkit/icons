package lucide

import (
	html "github.com/plainkit/html"
)

// Vault creates a Vault Lucide icon.
func Vault(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-vault", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2")),
		html.SvgCircle(html.ACx("7.5"), html.ACy("7.5"), html.AR(".5"), html.AFill("currentColor")),
		html.SvgPath(html.AD("m7.9 7.9 2.7 2.7")),
		html.SvgCircle(html.ACx("16.5"), html.ACy("7.5"), html.AR(".5"), html.AFill("currentColor")),
		html.SvgPath(html.AD("m13.4 10.6 2.7-2.7")),
		html.SvgCircle(html.ACx("7.5"), html.ACy("16.5"), html.AR(".5"), html.AFill("currentColor")),
		html.SvgPath(html.AD("m7.9 16.1 2.7-2.7")),
		html.SvgCircle(html.ACx("16.5"), html.ACy("16.5"), html.AR(".5"), html.AFill("currentColor")),
		html.SvgPath(html.AD("m13.4 13.4 2.7 2.7")),
		html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
