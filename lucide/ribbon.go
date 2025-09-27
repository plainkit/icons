package lucide

import (
	html "github.com/plainkit/html"
)

// Ribbon creates a Ribbon Lucide icon.
func Ribbon(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-ribbon", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M12 11.22C11 9.997 10 9 10 8a2 2 0 0 1 4 0c0 1-.998 2.002-2.01 3.22"))),
		html.Child(html.SvgPath(html.AD("m12 18 2.57-3.5"))),
		html.Child(html.SvgPath(html.AD("M6.243 9.016a7 7 0 0 1 11.507-.009"))),
		html.Child(html.SvgPath(html.AD("M9.35 14.53 12 11.22"))),
		html.Child(html.SvgPath(html.AD("M9.35 14.53C7.728 12.246 6 10.221 6 7a6 5 0 0 1 12 0c-.005 3.22-1.778 5.235-3.43 7.5l3.557 4.527a1 1 0 0 1-.203 1.43l-1.894 1.36a1 1 0 0 1-1.384-.215L12 18l-2.679 3.593a1 1 0 0 1-1.39.213l-1.865-1.353a1 1 0 0 1-.203-1.422z"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
