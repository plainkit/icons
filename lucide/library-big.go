package lucide

import (
	html "github.com/plainkit/html"
)

// LibraryBig creates a Library Big Lucide icon.
func LibraryBig(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-library-big", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("8"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("1"))),
		html.Child(html.SvgPath(html.AD("M7 3v18"))),
		html.Child(html.SvgPath(html.AD("M20.4 18.9c.2.5-.1 1.1-.6 1.3l-1.9.7c-.5.2-1.1-.1-1.3-.6L11.1 5.1c-.2-.5.1-1.1.6-1.3l1.9-.7c.5-.2 1.1.1 1.3.6Z"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
