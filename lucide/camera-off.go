package lucide

import (
	html "github.com/plainkit/html"
)

// CameraOff creates a Camera Off Lucide icon.
func CameraOff(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-camera-off", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M14.564 14.558a3 3 0 1 1-4.122-4.121"))),
		html.Child(html.SvgPath(html.AD("m2 2 20 20"))),
		html.Child(html.SvgPath(html.AD("M20 20H4a2 2 0 0 1-2-2V9a2 2 0 0 1 2-2h1.997a2 2 0 0 0 .819-.175"))),
		html.Child(html.SvgPath(html.AD("M9.695 4.024A2 2 0 0 1 10.004 4h3.993a2 2 0 0 1 1.76 1.05l.486.9A2 2 0 0 0 18.003 7H20a2 2 0 0 1 2 2v7.344"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
