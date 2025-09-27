package lucide

import (
	html "github.com/plainkit/html"
)

// MousePointer creates a Mouse Pointer Lucide icon.
func MousePointer(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-mouse-pointer", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M12.586 12.586 19 19"))),
		html.Child(html.SvgPath(html.AD("M3.688 3.037a.497.497 0 0 0-.651.651l6.5 15.999a.501.501 0 0 0 .947-.062l1.569-6.083a2 2 0 0 1 1.448-1.479l6.124-1.579a.5.5 0 0 0 .063-.947z"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
