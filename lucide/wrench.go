package lucide

import (
	html "github.com/plainkit/html"
)

// Wrench creates a Wrench Lucide icon.
func Wrench(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-wrench", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M14.7 6.3a1 1 0 0 0 0 1.4l1.6 1.6a1 1 0 0 0 1.4 0l3.106-3.105c.32-.322.863-.22.983.218a6 6 0 0 1-8.259 7.057l-7.91 7.91a1 1 0 0 1-2.999-3l7.91-7.91a6 6 0 0 1 7.057-8.259c.438.12.54.662.219.984z"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
