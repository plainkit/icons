package lucide

import (
	html "github.com/plainkit/html"
)

// MessageSquareDashed creates a Message Square Dashed Lucide icon.
func MessageSquareDashed(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-message-square-dashed", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M12 19h.01"))),
		html.Child(html.SvgPath(html.AD("M12 3h.01"))),
		html.Child(html.SvgPath(html.AD("M16 19h.01"))),
		html.Child(html.SvgPath(html.AD("M16 3h.01"))),
		html.Child(html.SvgPath(html.AD("M2 13h.01"))),
		html.Child(html.SvgPath(html.AD("M2 17v4.286a.71.71 0 0 0 1.212.502l2.202-2.202A2 2 0 0 1 6.828 19H8"))),
		html.Child(html.SvgPath(html.AD("M2 5a2 2 0 0 1 2-2"))),
		html.Child(html.SvgPath(html.AD("M2 9h.01"))),
		html.Child(html.SvgPath(html.AD("M20 3a2 2 0 0 1 2 2"))),
		html.Child(html.SvgPath(html.AD("M22 13h.01"))),
		html.Child(html.SvgPath(html.AD("M22 17a2 2 0 0 1-2 2"))),
		html.Child(html.SvgPath(html.AD("M22 9h.01"))),
		html.Child(html.SvgPath(html.AD("M8 3h.01"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
