package lucide

import (
	html "github.com/plainkit/html"
)

// Ship creates a Ship Lucide icon.
func Ship(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-ship", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M12 10.189V14"))),
		html.Child(html.SvgPath(html.AD("M12 2v3"))),
		html.Child(html.SvgPath(html.AD("M19 13V7a2 2 0 0 0-2-2H7a2 2 0 0 0-2 2v6"))),
		html.Child(html.SvgPath(html.AD("M19.38 20A11.6 11.6 0 0 0 21 14l-8.188-3.639a2 2 0 0 0-1.624 0L3 14a11.6 11.6 0 0 0 2.81 7.76"))),
		html.Child(html.SvgPath(html.AD("M2 21c.6.5 1.2 1 2.5 1 2.5 0 2.5-2 5-2 1.3 0 1.9.5 2.5 1s1.2 1 2.5 1c2.5 0 2.5-2 5-2 1.3 0 1.9.5 2.5 1"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
