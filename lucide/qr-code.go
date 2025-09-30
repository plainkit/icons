package lucide

import (
	html "github.com/plainkit/html"
)

// QrCode creates a Qr Code Lucide icon.
func QrCode(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-qr-code", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("5"), html.AHeight("5"), html.AX("3"), html.AY("3"), html.ARx("1")),
		html.SvgRect(html.AWidth("5"), html.AHeight("5"), html.AX("16"), html.AY("3"), html.ARx("1")),
		html.SvgRect(html.AWidth("5"), html.AHeight("5"), html.AX("3"), html.AY("16"), html.ARx("1")),
		html.SvgPath(html.AD("M21 16h-3a2 2 0 0 0-2 2v3")),
		html.SvgPath(html.AD("M21 21v.01")),
		html.SvgPath(html.AD("M12 7v3a2 2 0 0 1-2 2H7")),
		html.SvgPath(html.AD("M3 12h.01")),
		html.SvgPath(html.AD("M12 3h.01")),
		html.SvgPath(html.AD("M12 16v.01")),
		html.SvgPath(html.AD("M16 12h1")),
		html.SvgPath(html.AD("M21 12v.01")),
		html.SvgPath(html.AD("M12 21v-1")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
