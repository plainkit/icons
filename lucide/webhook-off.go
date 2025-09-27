package lucide

import (
	html "github.com/plainkit/html"
)

// WebhookOff creates a Webhook Off Lucide icon.
func WebhookOff(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-webhook-off", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M17 17h-5c-1.09-.02-1.94.92-2.5 1.9A3 3 0 1 1 2.57 15"))),
		html.Child(html.SvgPath(html.AD("M9 3.4a4 4 0 0 1 6.52.66"))),
		html.Child(html.SvgPath(html.AD("m6 17 3.1-5.8a2.5 2.5 0 0 0 .057-2.05"))),
		html.Child(html.SvgPath(html.AD("M20.3 20.3a4 4 0 0 1-2.3.7"))),
		html.Child(html.SvgPath(html.AD("M18.6 13a4 4 0 0 1 3.357 3.414"))),
		html.Child(html.SvgPath(html.AD("m12 6 .6 1"))),
		html.Child(html.SvgPath(html.AD("m2 2 20 20"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
