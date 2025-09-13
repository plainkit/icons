package lucide

import x "github.com/bloxui/blox"

// MailWarning creates a Mail Warning Lucide icon.
func MailWarning(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-mail-warning", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M22 10.5V6a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v12c0 1.1.9 2 2 2h12.5"))),
		x.Child(x.Path(x.D("m22 7-8.97 5.7a1.94 1.94 0 0 1-2.06 0L2 7"))),
		x.Child(x.Path(x.D("M20 14v4"))),
		x.Child(x.Path(x.D("M20 22v.01"))),
	)
	return x.Svg(svgArgs...)
}
