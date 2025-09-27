package lucide

import (
	html "github.com/plainkit/html"
)

// PencilOff creates a Pencil Off Lucide icon.
func PencilOff(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-pencil-off", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m10 10-6.157 6.162a2 2 0 0 0-.5.833l-1.322 4.36a.5.5 0 0 0 .622.624l4.358-1.323a2 2 0 0 0 .83-.5L14 13.982"))),
		html.Child(html.SvgPath(html.AD("m12.829 7.172 4.359-4.346a1 1 0 1 1 3.986 3.986l-4.353 4.353"))),
		html.Child(html.SvgPath(html.AD("m15 5 4 4"))),
		html.Child(html.SvgPath(html.AD("m2 2 20 20"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
