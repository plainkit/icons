package lucide

import x "github.com/plainkit/html"

// BadgeCheck creates a Badge Check Lucide icon.
func BadgeCheck(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-badge-check", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M3.85 8.62a4 4 0 0 1 4.78-4.77 4 4 0 0 1 6.74 0 4 4 0 0 1 4.78 4.78 4 4 0 0 1 0 6.74 4 4 0 0 1-4.77 4.78 4 4 0 0 1-6.75 0 4 4 0 0 1-4.78-4.77 4 4 0 0 1 0-6.76Z"))),
		x.Child(x.Path(x.D("m9 12 2 2 4-4"))),
	)
	return x.Svg(svgArgs...)
}
