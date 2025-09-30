package lucide

import (
	html "github.com/plainkit/html"
)

// BringToFront creates a Bring To Front Lucide icon.
func BringToFront(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-bring-to-front", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("8"), html.AHeight("8"), html.AX("8"), html.AY("8"), html.ARx("2")),
		html.SvgPath(html.AD("M4 10a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2")),
		html.SvgPath(html.AD("M14 20a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2v-4a2 2 0 0 0-2-2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
