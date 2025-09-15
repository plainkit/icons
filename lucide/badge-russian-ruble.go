package lucide

import x "github.com/plainkit/html"

// BadgeRussianRuble creates a Badge Russian Ruble Lucide icon.
func BadgeRussianRuble(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-badge-russian-ruble", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M3.85 8.62a4 4 0 0 1 4.78-4.77 4 4 0 0 1 6.74 0 4 4 0 0 1 4.78 4.78 4 4 0 0 1 0 6.74 4 4 0 0 1-4.77 4.78 4 4 0 0 1-6.75 0 4 4 0 0 1-4.78-4.77 4 4 0 0 1 0-6.76Z"))),
		x.Child(x.Path(x.D("M9 16h5"))),
		x.Child(x.Path(x.D("M9 12h5a2 2 0 1 0 0-4h-3v9"))),
	)
	return x.Svg(svgArgs...)
}
