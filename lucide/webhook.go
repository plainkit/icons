package lucide

import (
	html "github.com/plainkit/html"
)

// Webhook creates a Webhook Lucide icon.
func Webhook(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-webhook", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M18 16.98h-5.99c-1.1 0-1.95.94-2.48 1.9A4 4 0 0 1 2 17c.01-.7.2-1.4.57-2")),
		html.SvgPath(html.AD("m6 17 3.13-5.78c.53-.97.1-2.18-.5-3.1a4 4 0 1 1 6.89-4.06")),
		html.SvgPath(html.AD("m12 6 3.13 5.73C15.66 12.7 16.9 13 18 13a4 4 0 0 1 0 8")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
