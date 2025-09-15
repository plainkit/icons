package lucide

import x "github.com/plainkit/html"

// IdCardLanyard creates a Id Card Lanyard Lucide icon.
func IdCardLanyard(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-id-card-lanyard", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M13.5 8h-3"))),
		x.Child(x.Path(x.D("m15 2-1 2h3a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h3"))),
		x.Child(x.Path(x.D("M16.899 22A5 5 0 0 0 7.1 22"))),
		x.Child(x.Path(x.D("m9 2 3 6"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("15"), x.R("3"))),
	)
	return x.Svg(svgArgs...)
}
