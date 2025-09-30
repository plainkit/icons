package lucide

import (
	html "github.com/plainkit/html"
)

// Wallpaper creates a Wallpaper Lucide icon.
func Wallpaper(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-wallpaper", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12 17v4")),
		html.SvgPath(html.AD("M8 21h8")),
		html.SvgPath(html.AD("m9 17 6.1-6.1a2 2 0 0 1 2.81.01L22 15")),
		html.SvgCircle(html.ACx("8"), html.ACy("9"), html.AR("2")),
		html.SvgRect(html.AWidth("20"), html.AHeight("14"), html.AX("2"), html.AY("3"), html.ARx("2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
