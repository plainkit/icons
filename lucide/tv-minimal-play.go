package lucide

import (
	html "github.com/plainkit/html"
)

// TvMinimalPlay creates a Tv Minimal Play Lucide icon.
func TvMinimalPlay(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-tv-minimal-play", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M15.033 9.44a.647.647 0 0 1 0 1.12l-4.065 2.352a.645.645 0 0 1-.968-.56V7.648a.645.645 0 0 1 .967-.56z")),
		html.SvgPath(html.AD("M7 21h10")),
		html.SvgRect(html.AWidth("20"), html.AHeight("14"), html.AX("2"), html.AY("3"), html.ARx("2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
