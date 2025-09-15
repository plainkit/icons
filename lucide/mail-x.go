package lucide

import x "github.com/plainkit/blox"

// MailX creates a Mail X Lucide icon.
func MailX(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-mail-x", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M22 13V6a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v12c0 1.1.9 2 2 2h9"))),
		x.Child(x.Path(x.D("m22 7-8.97 5.7a1.94 1.94 0 0 1-2.06 0L2 7"))),
		x.Child(x.Path(x.D("m17 17 4 4"))),
		x.Child(x.Path(x.D("m21 17-4 4"))),
	)
	return x.Svg(svgArgs...)
}
