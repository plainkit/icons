package lucide

import (
	html "github.com/plainkit/html"
)

// FileLock creates a File Lock Lucide icon.
func FileLock(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-file-lock", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M15 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7Z")),
		html.SvgRect(html.AWidth("8"), html.AHeight("6"), html.AX("8"), html.AY("12"), html.ARx("1")),
		html.SvgPath(html.AD("M10 12v-2a2 2 0 1 1 4 0v2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
