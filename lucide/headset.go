package lucide

import (
	html "github.com/plainkit/html"
)

// Headset creates a Headset Lucide icon.
func Headset(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-headset", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M3 11h3a2 2 0 0 1 2 2v3a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-5Zm0 0a9 9 0 1 1 18 0m0 0v5a2 2 0 0 1-2 2h-1a2 2 0 0 1-2-2v-3a2 2 0 0 1 2-2h3Z"))),
		html.Child(html.SvgPath(html.AD("M21 16v2a4 4 0 0 1-4 4h-5"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
