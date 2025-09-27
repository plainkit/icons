package lucide

import (
	html "github.com/plainkit/html"
)

// Droplet creates a Droplet Lucide icon.
func Droplet(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-droplet", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M12 22a7 7 0 0 0 7-7c0-2-1-3.9-3-5.5s-3.5-4-4-6.5c-.5 2.5-2 4.9-4 6.5C6 11.1 5 13 5 15a7 7 0 0 0 7 7z"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
