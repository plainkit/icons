package lucide

import (
	html "github.com/plainkit/html"
)

// NotepadTextDashed creates a Notepad Text Dashed Lucide icon.
func NotepadTextDashed(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-notepad-text-dashed", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M8 2v4")),
		html.SvgPath(html.AD("M12 2v4")),
		html.SvgPath(html.AD("M16 2v4")),
		html.SvgPath(html.AD("M16 4h2a2 2 0 0 1 2 2v2")),
		html.SvgPath(html.AD("M20 12v2")),
		html.SvgPath(html.AD("M20 18v2a2 2 0 0 1-2 2h-1")),
		html.SvgPath(html.AD("M13 22h-2")),
		html.SvgPath(html.AD("M7 22H6a2 2 0 0 1-2-2v-2")),
		html.SvgPath(html.AD("M4 14v-2")),
		html.SvgPath(html.AD("M4 8V6a2 2 0 0 1 2-2h2")),
		html.SvgPath(html.AD("M8 10h6")),
		html.SvgPath(html.AD("M8 14h8")),
		html.SvgPath(html.AD("M8 18h5")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
