package lucide

import (
	html "github.com/plainkit/html"
)

// Import creates a Import Lucide icon.
func Import(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-import", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12 3v12")),
		html.SvgPath(html.AD("m8 11 4 4 4-4")),
		html.SvgPath(html.AD("M8 5H4a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2h-4")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
