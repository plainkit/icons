package lucide

import (
	html "github.com/plainkit/html"
)

// DropletOff creates a Droplet Off Lucide icon.
func DropletOff(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-droplet-off", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M18.715 13.186C18.29 11.858 17.384 10.607 16 9.5c-2-1.6-3.5-4-4-6.5a10.7 10.7 0 0 1-.884 2.586")),
		html.SvgPath(html.AD("m2 2 20 20")),
		html.SvgPath(html.AD("M8.795 8.797A11 11 0 0 1 8 9.5C6 11.1 5 13 5 15a7 7 0 0 0 13.222 3.208")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
