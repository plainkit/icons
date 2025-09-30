package lucide

import (
	html "github.com/plainkit/html"
)

// NotepadText creates a Notepad Text Lucide icon.
func NotepadText(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-notepad-text", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M8 2v4")),
		html.SvgPath(html.AD("M12 2v4")),
		html.SvgPath(html.AD("M16 2v4")),
		html.SvgRect(html.AWidth("16"), html.AHeight("18"), html.AX("4"), html.AY("4"), html.ARx("2")),
		html.SvgPath(html.AD("M8 10h6")),
		html.SvgPath(html.AD("M8 14h8")),
		html.SvgPath(html.AD("M8 18h5")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
