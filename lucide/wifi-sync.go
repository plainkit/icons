package lucide

import (
	html "github.com/plainkit/html"
)

// WifiSync creates a Wifi Sync Lucide icon.
func WifiSync(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-wifi-sync", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M11.965 10.105v4L13.5 12.5a5 5 0 0 1 8 1.5"))),
		html.Child(html.SvgPath(html.AD("M11.965 14.105h4"))),
		html.Child(html.SvgPath(html.AD("M17.965 18.105h4L20.43 19.71a5 5 0 0 1-8-1.5"))),
		html.Child(html.SvgPath(html.AD("M2 8.82a15 15 0 0 1 20 0"))),
		html.Child(html.SvgPath(html.AD("M21.965 22.105v-4"))),
		html.Child(html.SvgPath(html.AD("M5 12.86a10 10 0 0 1 3-2.032"))),
		html.Child(html.SvgPath(html.AD("M8.5 16.429h.01"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
