package lucide

import (
	html "github.com/plainkit/html"
)

// FileKey creates a File Key Lucide icon.
func FileKey(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-file-key", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M15 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7Z")),
		html.SvgCircle(html.ACx("10"), html.ACy("16"), html.AR("2")),
		html.SvgPath(html.AD("m16 10-4.5 4.5")),
		html.SvgPath(html.AD("m15 11 1 1")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
