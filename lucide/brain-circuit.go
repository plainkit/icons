package lucide

import (
	html "github.com/plainkit/html"
)

// BrainCircuit creates a Brain Circuit Lucide icon.
func BrainCircuit(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-brain-circuit", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M12 5a3 3 0 1 0-5.997.125 4 4 0 0 0-2.526 5.77 4 4 0 0 0 .556 6.588A4 4 0 1 0 12 18Z"))),
		html.Child(html.SvgPath(html.AD("M9 13a4.5 4.5 0 0 0 3-4"))),
		html.Child(html.SvgPath(html.AD("M6.003 5.125A3 3 0 0 0 6.401 6.5"))),
		html.Child(html.SvgPath(html.AD("M3.477 10.896a4 4 0 0 1 .585-.396"))),
		html.Child(html.SvgPath(html.AD("M6 18a4 4 0 0 1-1.967-.516"))),
		html.Child(html.SvgPath(html.AD("M12 13h4"))),
		html.Child(html.SvgPath(html.AD("M12 18h6a2 2 0 0 1 2 2v1"))),
		html.Child(html.SvgPath(html.AD("M12 8h8"))),
		html.Child(html.SvgPath(html.AD("M16 8V5a2 2 0 0 1 2-2"))),
		html.Child(html.SvgCircle(html.ACx("16"), html.ACy("13"), html.AR(".5"))),
		html.Child(html.SvgCircle(html.ACx("18"), html.ACy("3"), html.AR(".5"))),
		html.Child(html.SvgCircle(html.ACx("20"), html.ACy("21"), html.AR(".5"))),
		html.Child(html.SvgCircle(html.ACx("20"), html.ACy("8"), html.AR(".5"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
