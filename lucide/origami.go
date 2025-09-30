package lucide

import (
	html "github.com/plainkit/html"
)

// Origami creates a Origami Lucide icon.
func Origami(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-origami", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12 12V4a1 1 0 0 1 1-1h6.297a1 1 0 0 1 .651 1.759l-4.696 4.025")),
		html.SvgPath(html.AD("m12 21-7.414-7.414A2 2 0 0 1 4 12.172V6.415a1.002 1.002 0 0 1 1.707-.707L20 20.009")),
		html.SvgPath(html.AD("m12.214 3.381 8.414 14.966a1 1 0 0 1-.167 1.199l-1.168 1.163a1 1 0 0 1-.706.291H6.351a1 1 0 0 1-.625-.219L3.25 18.8a1 1 0 0 1 .631-1.781l4.165.027")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
