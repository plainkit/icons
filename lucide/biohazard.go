package lucide

import (
	html "github.com/plainkit/html"
)

// Biohazard creates a Biohazard Lucide icon.
func Biohazard(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-biohazard", args)
	children := []html.SvgArg{
		html.SvgCircle(html.ACx("12"), html.ACy("11.9"), html.AR("2")),
		html.SvgPath(html.AD("M6.7 3.4c-.9 2.5 0 5.2 2.2 6.7C6.5 9 3.7 9.6 2 11.6")),
		html.SvgPath(html.AD("m8.9 10.1 1.4.8")),
		html.SvgPath(html.AD("M17.3 3.4c.9 2.5 0 5.2-2.2 6.7 2.4-1.2 5.2-.6 6.9 1.5")),
		html.SvgPath(html.AD("m15.1 10.1-1.4.8")),
		html.SvgPath(html.AD("M16.7 20.8c-2.6-.4-4.6-2.6-4.7-5.3-.2 2.6-2.1 4.8-4.7 5.2")),
		html.SvgPath(html.AD("M12 13.9v1.6")),
		html.SvgPath(html.AD("M13.5 5.4c-1-.2-2-.2-3 0")),
		html.SvgPath(html.AD("M17 16.4c.7-.7 1.2-1.6 1.5-2.5")),
		html.SvgPath(html.AD("M5.5 13.9c.3.9.8 1.8 1.5 2.5")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
