package lucide

import (
	html "github.com/plainkit/html"
)

// TypeOutline creates a Type Outline Lucide icon.
func TypeOutline(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-type-outline", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M14 16.5a.5.5 0 0 0 .5.5h.5a2 2 0 0 1 0 4H9a2 2 0 0 1 0-4h.5a.5.5 0 0 0 .5-.5v-9a.5.5 0 0 0-.5-.5h-3a.5.5 0 0 0-.5.5V8a2 2 0 0 1-4 0V5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v3a2 2 0 0 1-4 0v-.5a.5.5 0 0 0-.5-.5h-3a.5.5 0 0 0-.5.5Z"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
