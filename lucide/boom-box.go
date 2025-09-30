package lucide

import (
	html "github.com/plainkit/html"
)

// BoomBox creates a Boom Box Lucide icon.
func BoomBox(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-boom-box", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M4 9V5a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v4")),
		html.SvgPath(html.AD("M8 8v1")),
		html.SvgPath(html.AD("M12 8v1")),
		html.SvgPath(html.AD("M16 8v1")),
		html.SvgRect(html.AWidth("20"), html.AHeight("12"), html.AX("2"), html.AY("9"), html.ARx("2")),
		html.SvgCircle(html.ACx("8"), html.ACy("15"), html.AR("2")),
		html.SvgCircle(html.ACx("16"), html.ACy("15"), html.AR("2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
