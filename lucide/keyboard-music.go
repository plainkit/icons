package lucide

import (
	html "github.com/plainkit/html"
)

// KeyboardMusic creates a Keyboard Music Lucide icon.
func KeyboardMusic(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-keyboard-music", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("20"), html.AHeight("16"), html.AX("2"), html.AY("4"), html.ARx("2")),
		html.SvgPath(html.AD("M6 8h4")),
		html.SvgPath(html.AD("M14 8h.01")),
		html.SvgPath(html.AD("M18 8h.01")),
		html.SvgPath(html.AD("M2 12h20")),
		html.SvgPath(html.AD("M6 12v4")),
		html.SvgPath(html.AD("M10 12v4")),
		html.SvgPath(html.AD("M14 12v4")),
		html.SvgPath(html.AD("M18 12v4")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
