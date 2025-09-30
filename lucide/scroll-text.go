package lucide

import (
	html "github.com/plainkit/html"
)

// ScrollText creates a Scroll Text Lucide icon.
func ScrollText(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-scroll-text", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M15 12h-5")),
		html.SvgPath(html.AD("M15 8h-5")),
		html.SvgPath(html.AD("M19 17V5a2 2 0 0 0-2-2H4")),
		html.SvgPath(html.AD("M8 21h12a2 2 0 0 0 2-2v-1a1 1 0 0 0-1-1H11a1 1 0 0 0-1 1v1a2 2 0 1 1-4 0V5a2 2 0 1 0-4 0v2a1 1 0 0 0 1 1h3")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
