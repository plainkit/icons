package lucide

import (
	html "github.com/plainkit/html"
)

// Magnet creates a Magnet Lucide icon.
func Magnet(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-magnet", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m12 15 4 4")),
		html.SvgPath(html.AD("M2.352 10.648a1.205 1.205 0 0 0 0 1.704l2.296 2.296a1.205 1.205 0 0 0 1.704 0l6.029-6.029a1 1 0 1 1 3 3l-6.029 6.029a1.205 1.205 0 0 0 0 1.704l2.296 2.296a1.205 1.205 0 0 0 1.704 0l6.365-6.367A1 1 0 0 0 8.716 4.282z")),
		html.SvgPath(html.AD("m5 8 4 4")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
