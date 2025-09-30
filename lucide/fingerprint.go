package lucide

import (
	html "github.com/plainkit/html"
)

// Fingerprint creates a Fingerprint Lucide icon.
func Fingerprint(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-fingerprint", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12 10a2 2 0 0 0-2 2c0 1.02-.1 2.51-.26 4")),
		html.SvgPath(html.AD("M14 13.12c0 2.38 0 6.38-1 8.88")),
		html.SvgPath(html.AD("M17.29 21.02c.12-.6.43-2.3.5-3.02")),
		html.SvgPath(html.AD("M2 12a10 10 0 0 1 18-6")),
		html.SvgPath(html.AD("M2 16h.01")),
		html.SvgPath(html.AD("M21.8 16c.2-2 .131-5.354 0-6")),
		html.SvgPath(html.AD("M5 19.5C5.5 18 6 15 6 12a6 6 0 0 1 .34-2")),
		html.SvgPath(html.AD("M8.65 22c.21-.66.45-1.32.57-2")),
		html.SvgPath(html.AD("M9 6.8a6 6 0 0 1 9 5.2v2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
