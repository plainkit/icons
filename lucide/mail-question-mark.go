package lucide

import x "github.com/plainkit/html"

// MailQuestionMark creates a Mail Question Mark Lucide icon.
func MailQuestionMark(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-mail-question-mark", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M22 10.5V6a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v12c0 1.1.9 2 2 2h12.5"))),
		x.Child(x.Path(x.D("m22 7-8.97 5.7a1.94 1.94 0 0 1-2.06 0L2 7"))),
		x.Child(x.Path(x.D("M18 15.28c.2-.4.5-.8.9-1a2.1 2.1 0 0 1 2.6.4c.3.4.5.8.5 1.3 0 1.3-2 2-2 2"))),
		x.Child(x.Path(x.D("M20 22v.01"))),
	)
	return x.Svg(svgArgs...)
}
