package lucide

import (
	html "github.com/plainkit/html"
)

// ArchiveX creates a Archive X Lucide icon.
func ArchiveX(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-archive-x", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("20"), html.AHeight("5"), html.AX("2"), html.AY("3"), html.ARx("1")),
		html.SvgPath(html.AD("M4 8v11a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8")),
		html.SvgPath(html.AD("m9.5 17 5-5")),
		html.SvgPath(html.AD("m9.5 12 5 5")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
