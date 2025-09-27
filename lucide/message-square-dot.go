package lucide

import (
	html "github.com/plainkit/html"
)

// MessageSquareDot creates a Message Square Dot Lucide icon.
func MessageSquareDot(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-message-square-dot", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M12.7 3H4a2 2 0 0 0-2 2v16.286a.71.71 0 0 0 1.212.502l2.202-2.202A2 2 0 0 1 6.828 19H20a2 2 0 0 0 2-2v-4.7"))),
		html.Child(html.SvgCircle(html.ACx("19"), html.ACy("6"), html.AR("3"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
