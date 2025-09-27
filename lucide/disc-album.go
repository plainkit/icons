package lucide

import (
	html "github.com/plainkit/html"
)

// DiscAlbum creates a Disc Album Lucide icon.
func DiscAlbum(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-disc-album", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2"))),
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("5"))),
		html.Child(html.SvgPath(html.AD("M12 12h.01"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
