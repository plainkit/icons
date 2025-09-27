package lucide

import (
	html "github.com/plainkit/html"
)

// MicVocal creates a Mic Vocal Lucide icon.
func MicVocal(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-mic-vocal", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m11 7.601-5.994 8.19a1 1 0 0 0 .1 1.298l.817.818a1 1 0 0 0 1.314.087L15.09 12"))),
		html.Child(html.SvgPath(html.AD("M16.5 21.174C15.5 20.5 14.372 20 13 20c-2.058 0-3.928 2.356-6 2-2.072-.356-2.775-3.369-1.5-4.5"))),
		html.Child(html.SvgCircle(html.ACx("16"), html.ACy("7"), html.AR("5"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
