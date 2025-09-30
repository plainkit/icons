package lucide

import (
	html "github.com/plainkit/html"
)

// SpellCheck creates a Spell Check Lucide icon.
func SpellCheck(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-spell-check", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m6 16 6-12 6 12")),
		html.SvgPath(html.AD("M8 12h8")),
		html.SvgPath(html.AD("m16 20 2 2 4-4")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
