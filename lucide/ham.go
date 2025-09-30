package lucide

import (
	html "github.com/plainkit/html"
)

// Ham creates a Ham Lucide icon.
func Ham(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-ham", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M13.144 21.144A7.274 10.445 45 1 0 2.856 10.856")),
		html.SvgPath(html.AD("M13.144 21.144A7.274 4.365 45 0 0 2.856 10.856a7.274 4.365 45 0 0 10.288 10.288")),
		html.SvgPath(html.AD("M16.565 10.435 18.6 8.4a2.501 2.501 0 1 0 1.65-4.65 2.5 2.5 0 1 0-4.66 1.66l-2.024 2.025")),
		html.SvgPath(html.AD("m8.5 16.5-1-1")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
