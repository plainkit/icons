package lucide

import (
	html "github.com/plainkit/html"
)

// ScanQrCode creates a Scan Qr Code Lucide icon.
func ScanQrCode(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-scan-qr-code", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M17 12v4a1 1 0 0 1-1 1h-4")),
		html.SvgPath(html.AD("M17 3h2a2 2 0 0 1 2 2v2")),
		html.SvgPath(html.AD("M17 8V7")),
		html.SvgPath(html.AD("M21 17v2a2 2 0 0 1-2 2h-2")),
		html.SvgPath(html.AD("M3 7V5a2 2 0 0 1 2-2h2")),
		html.SvgPath(html.AD("M7 17h.01")),
		html.SvgPath(html.AD("M7 21H5a2 2 0 0 1-2-2v-2")),
		html.SvgRect(html.AWidth("5"), html.AHeight("5"), html.AX("7"), html.AY("7"), html.ARx("1")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
