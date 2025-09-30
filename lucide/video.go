package lucide

import (
	html "github.com/plainkit/html"
)

// Video creates a Video Lucide icon.
func Video(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-video", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m16 13 5.223 3.482a.5.5 0 0 0 .777-.416V7.87a.5.5 0 0 0-.752-.432L16 10.5")),
		html.SvgRect(html.AWidth("14"), html.AHeight("12"), html.AX("2"), html.AY("6"), html.ARx("2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
