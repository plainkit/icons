package lucide

import (
	html "github.com/plainkit/html"
)

// Store creates a Store Lucide icon.
func Store(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-store", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M15 21v-5a1 1 0 0 0-1-1h-4a1 1 0 0 0-1 1v5")),
		html.SvgPath(html.AD("M17.774 10.31a1.12 1.12 0 0 0-1.549 0 2.5 2.5 0 0 1-3.451 0 1.12 1.12 0 0 0-1.548 0 2.5 2.5 0 0 1-3.452 0 1.12 1.12 0 0 0-1.549 0 2.5 2.5 0 0 1-3.77-3.248l2.889-4.184A2 2 0 0 1 7 2h10a2 2 0 0 1 1.653.873l2.895 4.192a2.5 2.5 0 0 1-3.774 3.244")),
		html.SvgPath(html.AD("M4 10.95V19a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-8.05")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
