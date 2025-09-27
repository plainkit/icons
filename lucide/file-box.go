package lucide

import (
	html "github.com/plainkit/html"
)

// FileBox creates a File Box Lucide icon.
func FileBox(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-file-box", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M14.5 22H18a2 2 0 0 0 2-2V7l-5-5H6a2 2 0 0 0-2 2v4"))),
		html.Child(html.SvgPath(html.AD("M14 2v4a2 2 0 0 0 2 2h4"))),
		html.Child(html.SvgPath(html.AD("M3 13.1a2 2 0 0 0-1 1.76v3.24a2 2 0 0 0 .97 1.78L6 21.7a2 2 0 0 0 2.03.01L11 19.9a2 2 0 0 0 1-1.76V14.9a2 2 0 0 0-.97-1.78L8 11.3a2 2 0 0 0-2.03-.01Z"))),
		html.Child(html.SvgPath(html.AD("M7 17v5"))),
		html.Child(html.SvgPath(html.AD("M11.7 14.2 7 17l-4.7-2.8"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
