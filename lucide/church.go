package lucide

import (
	html "github.com/plainkit/html"
)

// Church creates a Church Lucide icon.
func Church(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-church", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M10 9h4"))),
		html.Child(html.SvgPath(html.AD("M12 7v5"))),
		html.Child(html.SvgPath(html.AD("M14 21v-3a2 2 0 0 0-4 0v3"))),
		html.Child(html.SvgPath(html.AD("m18 9 3.52 2.147a1 1 0 0 1 .48.854V19a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2v-6.999a1 1 0 0 1 .48-.854L6 9"))),
		html.Child(html.SvgPath(html.AD("M6 21V7a1 1 0 0 1 .376-.782l5-3.999a1 1 0 0 1 1.249.001l5 4A1 1 0 0 1 18 7v14"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
