package lucide

import (
	html "github.com/plainkit/html"
)

// Container creates a Container Lucide icon.
func Container(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-container", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M22 7.7c0-.6-.4-1.2-.8-1.5l-6.3-3.9a1.72 1.72 0 0 0-1.7 0l-10.3 6c-.5.2-.9.8-.9 1.4v6.6c0 .5.4 1.2.8 1.5l6.3 3.9a1.72 1.72 0 0 0 1.7 0l10.3-6c.5-.3.9-1 .9-1.5Z")),
		html.SvgPath(html.AD("M10 21.9V14L2.1 9.1")),
		html.SvgPath(html.AD("m10 14 11.9-6.9")),
		html.SvgPath(html.AD("M14 19.8v-8.1")),
		html.SvgPath(html.AD("M18 17.5V9.4")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
