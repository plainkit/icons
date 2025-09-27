package lucide

import (
	html "github.com/plainkit/html"
)

// Feather creates a Feather Lucide icon.
func Feather(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-feather", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M12.67 19a2 2 0 0 0 1.416-.588l6.154-6.172a6 6 0 0 0-8.49-8.49L5.586 9.914A2 2 0 0 0 5 11.328V18a1 1 0 0 0 1 1z"))),
		html.Child(html.SvgPath(html.AD("M16 8 2 22"))),
		html.Child(html.SvgPath(html.AD("M17.5 15H9"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
