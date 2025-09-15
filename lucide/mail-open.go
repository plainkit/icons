package lucide

import x "github.com/plainkit/blox"

// MailOpen creates a Mail Open Lucide icon.
func MailOpen(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-mail-open", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M21.2 8.4c.5.38.8.97.8 1.6v10a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V10a2 2 0 0 1 .8-1.6l8-6a2 2 0 0 1 2.4 0l8 6Z"))),
		x.Child(x.Path(x.D("m22 10-8.97 5.7a1.94 1.94 0 0 1-2.06 0L2 10"))),
	)
	return x.Svg(svgArgs...)
}
