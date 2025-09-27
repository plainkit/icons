package lucide

import (
	html "github.com/plainkit/html"
)

// TouchpadOff creates a Touchpad Off Lucide icon.
func TouchpadOff(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-touchpad-off", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M12 20v-6"))),
		html.Child(html.SvgPath(html.AD("M19.656 14H22"))),
		html.Child(html.SvgPath(html.AD("M2 14h12"))),
		html.Child(html.SvgPath(html.AD("m2 2 20 20"))),
		html.Child(html.SvgPath(html.AD("M20 20H4a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2"))),
		html.Child(html.SvgPath(html.AD("M9.656 4H20a2 2 0 0 1 2 2v10.344"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
