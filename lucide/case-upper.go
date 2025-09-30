package lucide

import (
	html "github.com/plainkit/html"
)

// CaseUpper creates a Case Upper Lucide icon.
func CaseUpper(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-case-upper", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M15 11h4.5a1 1 0 0 1 0 5h-4a.5.5 0 0 1-.5-.5v-9a.5.5 0 0 1 .5-.5h3a1 1 0 0 1 0 5")),
		html.SvgPath(html.AD("m2 16 4.039-9.69a.5.5 0 0 1 .923 0L11 16")),
		html.SvgPath(html.AD("M3.304 13h6.392")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
