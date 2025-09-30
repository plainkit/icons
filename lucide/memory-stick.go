package lucide

import (
	html "github.com/plainkit/html"
)

// MemoryStick creates a Memory Stick Lucide icon.
func MemoryStick(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-memory-stick", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M6 19v-3")),
		html.SvgPath(html.AD("M10 19v-3")),
		html.SvgPath(html.AD("M14 19v-3")),
		html.SvgPath(html.AD("M18 19v-3")),
		html.SvgPath(html.AD("M8 11V9")),
		html.SvgPath(html.AD("M16 11V9")),
		html.SvgPath(html.AD("M12 11V9")),
		html.SvgPath(html.AD("M2 15h20")),
		html.SvgPath(html.AD("M2 7a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v1.1a2 2 0 0 0 0 3.837V17a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2v-5.1a2 2 0 0 0 0-3.837Z")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
