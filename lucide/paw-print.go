package lucide

import (
	html "github.com/plainkit/html"
)

// PawPrint creates a Paw Print Lucide icon.
func PawPrint(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-paw-print", args)
	children := []html.SvgArg{
		html.SvgCircle(html.ACx("11"), html.ACy("4"), html.AR("2")),
		html.SvgCircle(html.ACx("18"), html.ACy("8"), html.AR("2")),
		html.SvgCircle(html.ACx("20"), html.ACy("16"), html.AR("2")),
		html.SvgPath(html.AD("M9 10a5 5 0 0 1 5 5v3.5a3.5 3.5 0 0 1-6.84 1.045Q6.52 17.48 4.46 16.84A3.5 3.5 0 0 1 5.5 10Z")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
