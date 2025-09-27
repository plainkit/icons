package lucide

import (
	html "github.com/plainkit/html"
)

// PartyPopper creates a Party Popper Lucide icon.
func PartyPopper(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-party-popper", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M5.8 11.3 2 22l10.7-3.79"))),
		html.Child(html.SvgPath(html.AD("M4 3h.01"))),
		html.Child(html.SvgPath(html.AD("M22 8h.01"))),
		html.Child(html.SvgPath(html.AD("M15 2h.01"))),
		html.Child(html.SvgPath(html.AD("M22 20h.01"))),
		html.Child(html.SvgPath(html.AD("m22 2-2.24.75a2.9 2.9 0 0 0-1.96 3.12c.1.86-.57 1.63-1.45 1.63h-.38c-.86 0-1.6.6-1.76 1.44L14 10"))),
		html.Child(html.SvgPath(html.AD("m22 13-.82-.33c-.86-.34-1.82.2-1.98 1.11c-.11.7-.72 1.22-1.43 1.22H17"))),
		html.Child(html.SvgPath(html.AD("m11 2 .33.82c.34.86-.2 1.82-1.11 1.98C9.52 4.9 9 5.52 9 6.23V7"))),
		html.Child(html.SvgPath(html.AD("M11 13c1.93 1.93 2.83 4.17 2 5-.83.83-3.07-.07-5-2-1.93-1.93-2.83-4.17-2-5 .83-.83 3.07.07 5 2Z"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
