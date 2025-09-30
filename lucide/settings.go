package lucide

import (
	html "github.com/plainkit/html"
)

// Settings creates a Settings Lucide icon.
func Settings(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-settings", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M9.671 4.136a2.34 2.34 0 0 1 4.659 0 2.34 2.34 0 0 0 3.319 1.915 2.34 2.34 0 0 1 2.33 4.033 2.34 2.34 0 0 0 0 3.831 2.34 2.34 0 0 1-2.33 4.033 2.34 2.34 0 0 0-3.319 1.915 2.34 2.34 0 0 1-4.659 0 2.34 2.34 0 0 0-3.32-1.915 2.34 2.34 0 0 1-2.33-4.033 2.34 2.34 0 0 0 0-3.831A2.34 2.34 0 0 1 6.35 6.051a2.34 2.34 0 0 0 3.319-1.915")),
		html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("3")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
