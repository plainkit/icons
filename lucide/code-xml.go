package lucide

import (
	html "github.com/plainkit/html"
)

// CodeXml creates a Code Xml Lucide icon.
func CodeXml(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-code-xml", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m18 16 4-4-4-4"))),
		html.Child(html.SvgPath(html.AD("m6 8-4 4 4 4"))),
		html.Child(html.SvgPath(html.AD("m14.5 4-5 16"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
