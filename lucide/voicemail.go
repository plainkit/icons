package lucide

import (
	html "github.com/plainkit/html"
)

// Voicemail creates a Voicemail Lucide icon.
func Voicemail(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-voicemail", args)
	children := []html.SvgArg{
		html.Child(html.SvgCircle(html.ACx("6"), html.ACy("12"), html.AR("4"))),
		html.Child(html.SvgCircle(html.ACx("18"), html.ACy("12"), html.AR("4"))),
		html.Child(html.SvgLine(html.AX1("6"), html.AX2("18"), html.AY1("16"), html.AY2("16"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
