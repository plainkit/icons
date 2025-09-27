package lucide

import (
	html "github.com/plainkit/html"
)

// MessageCircleDashed creates a Message Circle Dashed Lucide icon.
func MessageCircleDashed(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-message-circle-dashed", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M10.1 2.182a10 10 0 0 1 3.8 0"))),
		html.Child(html.SvgPath(html.AD("M13.9 21.818a10 10 0 0 1-3.8 0"))),
		html.Child(html.SvgPath(html.AD("M17.609 3.72a10 10 0 0 1 2.69 2.7"))),
		html.Child(html.SvgPath(html.AD("M2.182 13.9a10 10 0 0 1 0-3.8"))),
		html.Child(html.SvgPath(html.AD("M20.28 17.61a10 10 0 0 1-2.7 2.69"))),
		html.Child(html.SvgPath(html.AD("M21.818 10.1a10 10 0 0 1 0 3.8"))),
		html.Child(html.SvgPath(html.AD("M3.721 6.391a10 10 0 0 1 2.7-2.69"))),
		html.Child(html.SvgPath(html.AD("m6.163 21.117-2.906.85a1 1 0 0 1-1.236-1.169l.965-2.98"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
