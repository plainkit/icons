package lucide

import (
	html "github.com/plainkit/html"
)

// Lectern creates a Lectern Lucide icon.
func Lectern(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-lectern", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M16 12h3a2 2 0 0 0 1.902-1.38l1.056-3.333A1 1 0 0 0 21 6H3a1 1 0 0 0-.958 1.287l1.056 3.334A2 2 0 0 0 5 12h3")),
		html.SvgPath(html.AD("M18 6V3a1 1 0 0 0-1-1h-3")),
		html.SvgRect(html.AWidth("8"), html.AHeight("12"), html.AX("8"), html.AY("10"), html.ARx("1")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
