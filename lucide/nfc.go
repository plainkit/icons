package lucide

import (
	html "github.com/plainkit/html"
)

// Nfc creates a Nfc Lucide icon.
func Nfc(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-nfc", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M6 8.32a7.43 7.43 0 0 1 0 7.36")),
		html.SvgPath(html.AD("M9.46 6.21a11.76 11.76 0 0 1 0 11.58")),
		html.SvgPath(html.AD("M12.91 4.1a15.91 15.91 0 0 1 .01 15.8")),
		html.SvgPath(html.AD("M16.37 2a20.16 20.16 0 0 1 0 20")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
