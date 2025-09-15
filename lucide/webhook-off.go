package lucide

import x "github.com/plainkit/blox"

// WebhookOff creates a Webhook Off Lucide icon.
func WebhookOff(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-webhook-off", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M17 17h-5c-1.09-.02-1.94.92-2.5 1.9A3 3 0 1 1 2.57 15"))),
		x.Child(x.Path(x.D("M9 3.4a4 4 0 0 1 6.52.66"))),
		x.Child(x.Path(x.D("m6 17 3.1-5.8a2.5 2.5 0 0 0 .057-2.05"))),
		x.Child(x.Path(x.D("M20.3 20.3a4 4 0 0 1-2.3.7"))),
		x.Child(x.Path(x.D("M18.6 13a4 4 0 0 1 3.357 3.414"))),
		x.Child(x.Path(x.D("m12 6 .6 1"))),
		x.Child(x.Path(x.D("m2 2 20 20"))),
	)
	return x.Svg(svgArgs...)
}
