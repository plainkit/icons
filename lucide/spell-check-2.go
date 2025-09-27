package lucide

import (
	html "github.com/plainkit/html"
)

// SpellCheck2 creates a Spell Check 2 Lucide icon.
func SpellCheck2(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-spell-check-2", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m6 16 6-12 6 12"))),
		html.Child(html.SvgPath(html.AD("M8 12h8"))),
		html.Child(html.SvgPath(html.AD("M4 21c1.1 0 1.1-1 2.3-1s1.1 1 2.3 1c1.1 0 1.1-1 2.3-1 1.1 0 1.1 1 2.3 1 1.1 0 1.1-1 2.3-1 1.1 0 1.1 1 2.3 1 1.1 0 1.1-1 2.3-1"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
