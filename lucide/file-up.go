package lucide

import (
	html "github.com/plainkit/html"
)

// FileUp creates a File Up Lucide icon.
func FileUp(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-file-up", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M15 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7Z")),
		html.SvgPath(html.AD("M14 2v4a2 2 0 0 0 2 2h4")),
		html.SvgPath(html.AD("M12 12v6")),
		html.SvgPath(html.AD("m15 15-3-3-3 3")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
