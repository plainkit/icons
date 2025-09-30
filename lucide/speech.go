package lucide

import (
	html "github.com/plainkit/html"
)

// Speech creates a Speech Lucide icon.
func Speech(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-speech", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M8.8 20v-4.1l1.9.2a2.3 2.3 0 0 0 2.164-2.1V8.3A5.37 5.37 0 0 0 2 8.25c0 2.8.656 3.054 1 4.55a5.77 5.77 0 0 1 .029 2.758L2 20")),
		html.SvgPath(html.AD("M19.8 17.8a7.5 7.5 0 0 0 .003-10.603")),
		html.SvgPath(html.AD("M17 15a3.5 3.5 0 0 0-.025-4.975")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
