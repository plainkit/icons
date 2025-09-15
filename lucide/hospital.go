package lucide

import x "github.com/plainkit/html"

// Hospital creates a Hospital Lucide icon.
func Hospital(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-hospital", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 7v4"))),
		x.Child(x.Path(x.D("M14 21v-3a2 2 0 0 0-4 0v3"))),
		x.Child(x.Path(x.D("M14 9h-4"))),
		x.Child(x.Path(x.D("M18 11h2a2 2 0 0 1 2 2v6a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2v-9a2 2 0 0 1 2-2h2"))),
		x.Child(x.Path(x.D("M18 21V5a2 2 0 0 0-2-2H8a2 2 0 0 0-2 2v16"))),
	)
	return x.Svg(svgArgs...)
}
