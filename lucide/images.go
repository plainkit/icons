package lucide

import (
	html "github.com/plainkit/html"
)

// Images creates a Images Lucide icon.
func Images(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-images", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m22 11-1.296-1.296a2.4 2.4 0 0 0-3.408 0L11 16")),
		html.SvgPath(html.AD("M4 8a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2")),
		html.SvgCircle(html.ACx("13"), html.ACy("7"), html.AR("1"), html.AFill("currentColor")),
		html.SvgRect(html.AWidth("14"), html.AHeight("14"), html.AX("8"), html.AY("2"), html.ARx("2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
