package lucide

import (
	html "github.com/plainkit/html"
)

// HdmiPort creates a Hdmi Port Lucide icon.
func HdmiPort(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-hdmi-port", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M22 9a1 1 0 0 0-1-1H3a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1h1l2 2h12l2-2h1a1 1 0 0 0 1-1Z"))),
		html.Child(html.SvgPath(html.AD("M7.5 12h9"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
