package lucide

import (
	html "github.com/plainkit/html"
)

// SquareDashedMousePointer creates a Square Dashed Mouse Pointer Lucide icon.
func SquareDashedMousePointer(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-square-dashed-mouse-pointer", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12.034 12.681a.498.498 0 0 1 .647-.647l9 3.5a.5.5 0 0 1-.033.943l-3.444 1.068a1 1 0 0 0-.66.66l-1.067 3.443a.5.5 0 0 1-.943.033z")),
		html.SvgPath(html.AD("M5 3a2 2 0 0 0-2 2")),
		html.SvgPath(html.AD("M19 3a2 2 0 0 1 2 2")),
		html.SvgPath(html.AD("M5 21a2 2 0 0 1-2-2")),
		html.SvgPath(html.AD("M9 3h1")),
		html.SvgPath(html.AD("M9 21h2")),
		html.SvgPath(html.AD("M14 3h1")),
		html.SvgPath(html.AD("M3 9v1")),
		html.SvgPath(html.AD("M21 9v2")),
		html.SvgPath(html.AD("M3 14v1")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
