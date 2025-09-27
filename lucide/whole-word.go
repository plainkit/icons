package lucide

import (
	html "github.com/plainkit/html"
)

// WholeWord creates a Whole Word Lucide icon.
func WholeWord(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-whole-word", args)
	children := []html.SvgArg{
		html.Child(html.SvgCircle(html.ACx("7"), html.ACy("12"), html.AR("3"))),
		html.Child(html.SvgPath(html.AD("M10 9v6"))),
		html.Child(html.SvgCircle(html.ACx("17"), html.ACy("12"), html.AR("3"))),
		html.Child(html.SvgPath(html.AD("M14 7v8"))),
		html.Child(html.SvgPath(html.AD("M22 17v1c0 .5-.5 1-1 1H3c-.5 0-1-.5-1-1v-1"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
