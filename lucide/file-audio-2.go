package lucide

import (
	html "github.com/plainkit/html"
)

// FileAudio2 creates a File Audio 2 Lucide icon.
func FileAudio2(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-file-audio-2", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M4 22h14a2 2 0 0 0 2-2V7l-5-5H6a2 2 0 0 0-2 2v2")),
		html.SvgPath(html.AD("M14 2v4a2 2 0 0 0 2 2h4")),
		html.SvgCircle(html.ACx("3"), html.ACy("17"), html.AR("1")),
		html.SvgPath(html.AD("M2 17v-3a4 4 0 0 1 8 0v3")),
		html.SvgCircle(html.ACx("9"), html.ACy("17"), html.AR("1")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
