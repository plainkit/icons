package lucide

import (
	html "github.com/plainkit/html"
)

// Twitch creates a Twitch Lucide icon.
func Twitch(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-twitch", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M21 2H3v16h5v4l4-4h5l4-4V2zm-10 9V7m5 4V7"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
