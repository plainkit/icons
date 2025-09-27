package lucide

import (
	html "github.com/plainkit/html"
)

// SquareMousePointer creates a Square Mouse Pointer Lucide icon.
func SquareMousePointer(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-square-mouse-pointer", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M12.034 12.681a.498.498 0 0 1 .647-.647l9 3.5a.5.5 0 0 1-.033.943l-3.444 1.068a1 1 0 0 0-.66.66l-1.067 3.443a.5.5 0 0 1-.943.033z"))),
		html.Child(html.SvgPath(html.AD("M21 11V5a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h6"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
