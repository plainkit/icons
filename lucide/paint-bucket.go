package lucide

import (
	html "github.com/plainkit/html"
)

// PaintBucket creates a Paint Bucket Lucide icon.
func PaintBucket(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-paint-bucket", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m19 11-8-8-8.6 8.6a2 2 0 0 0 0 2.8l5.2 5.2c.8.8 2 .8 2.8 0L19 11Z"))),
		html.Child(html.SvgPath(html.AD("m5 2 5 5"))),
		html.Child(html.SvgPath(html.AD("M2 13h15"))),
		html.Child(html.SvgPath(html.AD("M22 20a2 2 0 1 1-4 0c0-1.6 1.7-2.4 2-4 .3 1.6 2 2.4 2 4Z"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
