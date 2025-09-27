package lucide

import (
	html "github.com/plainkit/html"
)

// ToyBrick creates a Toy Brick Lucide icon.
func ToyBrick(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-toy-brick", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("18"), html.AHeight("12"), html.AX("3"), html.AY("8"), html.ARx("1"))),
		html.Child(html.SvgPath(html.AD("M10 8V5c0-.6-.4-1-1-1H6a1 1 0 0 0-1 1v3"))),
		html.Child(html.SvgPath(html.AD("M19 8V5c0-.6-.4-1-1-1h-3a1 1 0 0 0-1 1v3"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
