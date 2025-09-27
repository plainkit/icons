package lucide

import (
	html "github.com/plainkit/html"
)

// Signature creates a Signature Lucide icon.
func Signature(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-signature", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m21 17-2.156-1.868A.5.5 0 0 0 18 15.5v.5a1 1 0 0 1-1 1h-2a1 1 0 0 1-1-1c0-2.545-3.991-3.97-8.5-4a1 1 0 0 0 0 5c4.153 0 4.745-11.295 5.708-13.5a2.5 2.5 0 1 1 3.31 3.284"))),
		html.Child(html.SvgPath(html.AD("M3 21h18"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
