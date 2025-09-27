package lucide

import (
	html "github.com/plainkit/html"
)

// FileLock2 creates a File Lock 2 Lucide icon.
func FileLock2(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-file-lock-2", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M4 22h14a2 2 0 0 0 2-2V7l-5-5H6a2 2 0 0 0-2 2v1"))),
		html.Child(html.SvgPath(html.AD("M14 2v4a2 2 0 0 0 2 2h4"))),
		html.Child(html.SvgRect(html.AWidth("8"), html.AHeight("5"), html.AX("2"), html.AY("13"), html.ARx("1"))),
		html.Child(html.SvgPath(html.AD("M8 13v-2a2 2 0 1 0-4 0v2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
