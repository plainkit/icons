package lucide

import (
	html "github.com/plainkit/html"
)

// BrickWall creates a Brick Wall Lucide icon.
func BrickWall(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-brick-wall", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2")),
		html.SvgPath(html.AD("M12 9v6")),
		html.SvgPath(html.AD("M16 15v6")),
		html.SvgPath(html.AD("M16 3v6")),
		html.SvgPath(html.AD("M3 15h18")),
		html.SvgPath(html.AD("M3 9h18")),
		html.SvgPath(html.AD("M8 15v6")),
		html.SvgPath(html.AD("M8 3v6")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
