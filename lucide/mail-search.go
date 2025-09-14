package lucide

import x "github.com/bloxui/blox"

// MailSearch creates a Mail Search Lucide icon.
func MailSearch(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-mail-search", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M22 12.5V6a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v12c0 1.1.9 2 2 2h7.5"))),
		x.Child(x.Path(x.D("m22 7-8.97 5.7a1.94 1.94 0 0 1-2.06 0L2 7"))),
		x.Child(x.Path(x.D("M18 21a3 3 0 1 0 0-6 3 3 0 0 0 0 6Z"))),
		x.Child(x.Circle(x.Cx("18"), x.Cy("18"), x.R("3"))),
		x.Child(x.Path(x.D("m22 22-1.5-1.5"))),
	)
	return x.Svg(svgArgs...)
}
