package lucide

import (
	html "github.com/plainkit/html"
)

// FileMusic creates a File Music Lucide icon.
func FileMusic(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-file-music", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M10.5 22H18a2 2 0 0 0 2-2V7l-5-5H6a2 2 0 0 0-2 2v8.4"))),
		html.Child(html.SvgPath(html.AD("M8 18v-7.7L16 9v7"))),
		html.Child(html.SvgCircle(html.ACx("14"), html.ACy("16"), html.AR("2"))),
		html.Child(html.SvgCircle(html.ACx("6"), html.ACy("18"), html.AR("2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
