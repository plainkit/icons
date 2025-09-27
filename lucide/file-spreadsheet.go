package lucide

import (
	html "github.com/plainkit/html"
)

// FileSpreadsheet creates a File Spreadsheet Lucide icon.
func FileSpreadsheet(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-file-spreadsheet", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M15 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7Z"))),
		html.Child(html.SvgPath(html.AD("M14 2v4a2 2 0 0 0 2 2h4"))),
		html.Child(html.SvgPath(html.AD("M8 13h2"))),
		html.Child(html.SvgPath(html.AD("M14 13h2"))),
		html.Child(html.SvgPath(html.AD("M8 17h2"))),
		html.Child(html.SvgPath(html.AD("M14 17h2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
