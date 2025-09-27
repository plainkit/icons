package lucide

import (
	html "github.com/plainkit/html"
)

// FileArchive creates a File Archive Lucide icon.
func FileArchive(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-file-archive", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M10 12v-1"))),
		html.Child(html.SvgPath(html.AD("M10 18v-2"))),
		html.Child(html.SvgPath(html.AD("M10 7V6"))),
		html.Child(html.SvgPath(html.AD("M14 2v4a2 2 0 0 0 2 2h4"))),
		html.Child(html.SvgPath(html.AD("M15.5 22H18a2 2 0 0 0 2-2V7l-5-5H6a2 2 0 0 0-2 2v16a2 2 0 0 0 .274 1.01"))),
		html.Child(html.SvgCircle(html.ACx("10"), html.ACy("20"), html.AR("2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
