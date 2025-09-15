package lucide

import x "github.com/plainkit/html"

// PrinterCheck creates a Printer Check Lucide icon.
func PrinterCheck(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-printer-check", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M13.5 22H7a1 1 0 0 1-1-1v-6a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1v.5"))),
		x.Child(x.Path(x.D("m16 19 2 2 4-4"))),
		x.Child(x.Path(x.D("M6 18H4a2 2 0 0 1-2-2v-5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v2"))),
		x.Child(x.Path(x.D("M6 9V3a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1v6"))),
	)
	return x.Svg(svgArgs...)
}
