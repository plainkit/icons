package lucide

import (
	html "github.com/plainkit/html"
)

// Gamepad2 creates a Gamepad 2 Lucide icon.
func Gamepad2(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-gamepad-2", args)
	children := []html.SvgArg{
		html.SvgLine(html.AX1("6"), html.AX2("10"), html.AY1("11"), html.AY2("11")),
		html.SvgLine(html.AX1("8"), html.AX2("8"), html.AY1("9"), html.AY2("13")),
		html.SvgLine(html.AX1("15"), html.AX2("15.01"), html.AY1("12"), html.AY2("12")),
		html.SvgLine(html.AX1("18"), html.AX2("18.01"), html.AY1("10"), html.AY2("10")),
		html.SvgPath(html.AD("M17.32 5H6.68a4 4 0 0 0-3.978 3.59c-.006.052-.01.101-.017.152C2.604 9.416 2 14.456 2 16a3 3 0 0 0 3 3c1 0 1.5-.5 2-1l1.414-1.414A2 2 0 0 1 9.828 16h4.344a2 2 0 0 1 1.414.586L17 18c.5.5 1 1 2 1a3 3 0 0 0 3-3c0-1.545-.604-6.584-.685-7.258-.007-.05-.011-.1-.017-.151A4 4 0 0 0 17.32 5z")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
