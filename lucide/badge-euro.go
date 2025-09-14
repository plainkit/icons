package lucide

import x "github.com/bloxui/blox"

// BadgeEuro creates a Badge Euro Lucide icon.
func BadgeEuro(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-badge-euro", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M3.85 8.62a4 4 0 0 1 4.78-4.77 4 4 0 0 1 6.74 0 4 4 0 0 1 4.78 4.78 4 4 0 0 1 0 6.74 4 4 0 0 1-4.77 4.78 4 4 0 0 1-6.75 0 4 4 0 0 1-4.78-4.77 4 4 0 0 1 0-6.76Z"))),
		x.Child(x.Path(x.D("M7 12h5"))),
		x.Child(x.Path(x.D("M15 9.4a4 4 0 1 0 0 5.2"))),
	)
	return x.Svg(svgArgs...)
}
