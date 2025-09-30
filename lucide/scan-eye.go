package lucide

import (
	html "github.com/plainkit/html"
)

// ScanEye creates a Scan Eye Lucide icon.
func ScanEye(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-scan-eye", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M3 7V5a2 2 0 0 1 2-2h2")),
		html.SvgPath(html.AD("M17 3h2a2 2 0 0 1 2 2v2")),
		html.SvgPath(html.AD("M21 17v2a2 2 0 0 1-2 2h-2")),
		html.SvgPath(html.AD("M7 21H5a2 2 0 0 1-2-2v-2")),
		html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("1")),
		html.SvgPath(html.AD("M18.944 12.33a1 1 0 0 0 0-.66 7.5 7.5 0 0 0-13.888 0 1 1 0 0 0 0 .66 7.5 7.5 0 0 0 13.888 0")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
