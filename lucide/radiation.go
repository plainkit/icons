package lucide

import (
	html "github.com/plainkit/html"
)

// Radiation creates a Radiation Lucide icon.
func Radiation(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-radiation", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M12 12h.01"))),
		html.Child(html.SvgPath(html.AD("M14 15.4641a4 4 0 0 1-4 0L7.52786 19.74597 A 1 1 0 0 0 7.99303 21.16211 10 10 0 0 0 16.00697 21.16211 1 1 0 0 0 16.47214 19.74597z"))),
		html.Child(html.SvgPath(html.AD("M16 12a4 4 0 0 0-2-3.464l2.472-4.282a1 1 0 0 1 1.46-.305 10 10 0 0 1 4.006 6.94A1 1 0 0 1 21 12z"))),
		html.Child(html.SvgPath(html.AD("M8 12a4 4 0 0 1 2-3.464L7.528 4.254a1 1 0 0 0-1.46-.305 10 10 0 0 0-4.006 6.94A1 1 0 0 0 3 12z"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
