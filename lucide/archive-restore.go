package lucide

import (
	html "github.com/plainkit/html"
)

// ArchiveRestore creates a Archive Restore Lucide icon.
func ArchiveRestore(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-archive-restore", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("20"), html.AHeight("5"), html.AX("2"), html.AY("3"), html.ARx("1")),
		html.SvgPath(html.AD("M4 8v11a2 2 0 0 0 2 2h2")),
		html.SvgPath(html.AD("M20 8v11a2 2 0 0 1-2 2h-2")),
		html.SvgPath(html.AD("m9 15 3-3 3 3")),
		html.SvgPath(html.AD("M12 12v9")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
