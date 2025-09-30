package lucide

import (
	html "github.com/plainkit/html"
)

// Omega creates a Omega Lucide icon.
func Omega(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-omega", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M3 20h4.5a.5.5 0 0 0 .5-.5v-.282a.52.52 0 0 0-.247-.437 8 8 0 1 1 8.494-.001.52.52 0 0 0-.247.438v.282a.5.5 0 0 0 .5.5H21")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
