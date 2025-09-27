package lucide

import (
	html "github.com/plainkit/html"
)

// CircleFadingArrowUp creates a Circle Fading Arrow Up Lucide icon.
func CircleFadingArrowUp(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-circle-fading-arrow-up", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M12 2a10 10 0 0 1 7.38 16.75"))),
		html.Child(html.SvgPath(html.AD("m16 12-4-4-4 4"))),
		html.Child(html.SvgPath(html.AD("M12 16V8"))),
		html.Child(html.SvgPath(html.AD("M2.5 8.875a10 10 0 0 0-.5 3"))),
		html.Child(html.SvgPath(html.AD("M2.83 16a10 10 0 0 0 2.43 3.4"))),
		html.Child(html.SvgPath(html.AD("M4.636 5.235a10 10 0 0 1 .891-.857"))),
		html.Child(html.SvgPath(html.AD("M8.644 21.42a10 10 0 0 0 7.631-.38"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
