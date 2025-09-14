package lucide

import x "github.com/bloxui/blox"

// MailMinus creates a Mail Minus Lucide icon.
func MailMinus(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-mail-minus", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M22 15V6a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v12c0 1.1.9 2 2 2h8"))),
		x.Child(x.Path(x.D("m22 7-8.97 5.7a1.94 1.94 0 0 1-2.06 0L2 7"))),
		x.Child(x.Path(x.D("M16 19h6"))),
	)
	return x.Svg(svgArgs...)
}
