package lucide

import (
	html "github.com/plainkit/html"
)

// TowerControl creates a Tower Control Lucide icon.
func TowerControl(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-tower-control", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M18.2 12.27 20 6H4l1.8 6.27a1 1 0 0 0 .95.73h10.5a1 1 0 0 0 .96-.73Z"))),
		html.Child(html.SvgPath(html.AD("M8 13v9"))),
		html.Child(html.SvgPath(html.AD("M16 22v-9"))),
		html.Child(html.SvgPath(html.AD("m9 6 1 7"))),
		html.Child(html.SvgPath(html.AD("m15 6-1 7"))),
		html.Child(html.SvgPath(html.AD("M12 6V2"))),
		html.Child(html.SvgPath(html.AD("M13 2h-2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
