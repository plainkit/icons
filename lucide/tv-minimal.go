package lucide

import (
	html "github.com/plainkit/html"
)

// TvMinimal creates a Tv Minimal Lucide icon.
func TvMinimal(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-tv-minimal", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M7 21h10"))),
		html.Child(html.SvgRect(html.AWidth("20"), html.AHeight("14"), html.AX("2"), html.AY("3"), html.ARx("2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
