package lucide

import (
	html "github.com/plainkit/html"
)

// FileAudio creates a File Audio Lucide icon.
func FileAudio(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-file-audio", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M17.5 22h.5a2 2 0 0 0 2-2V7l-5-5H6a2 2 0 0 0-2 2v3"))),
		html.Child(html.SvgPath(html.AD("M14 2v4a2 2 0 0 0 2 2h4"))),
		html.Child(html.SvgPath(html.AD("M2 19a2 2 0 1 1 4 0v1a2 2 0 1 1-4 0v-4a6 6 0 0 1 12 0v4a2 2 0 1 1-4 0v-1a2 2 0 1 1 4 0"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
