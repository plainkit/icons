package lucide

import (
	html "github.com/plainkit/html"
)

// FileVolume creates a File Volume Lucide icon.
func FileVolume(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-file-volume", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M11 11a5 5 0 0 1 0 6"))),
		html.Child(html.SvgPath(html.AD("M14 2v4a2 2 0 0 0 2 2h4"))),
		html.Child(html.SvgPath(html.AD("M4 6.765V4a2 2 0 0 1 2-2h9l5 5v13a2 2 0 0 1-2 2H6a2 2 0 0 1-.93-.23"))),
		html.Child(html.SvgPath(html.AD("M7 10.51a.5.5 0 0 0-.826-.38l-1.893 1.628A1 1 0 0 1 3.63 12H2.5a.5.5 0 0 0-.5.5v3a.5.5 0 0 0 .5.5h1.129a1 1 0 0 1 .652.242l1.893 1.63a.5.5 0 0 0 .826-.38z"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
