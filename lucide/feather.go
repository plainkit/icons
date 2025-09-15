package lucide

import x "github.com/plainkit/html"

// Feather creates a Feather Lucide icon.
func Feather(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-feather", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12.67 19a2 2 0 0 0 1.416-.588l6.154-6.172a6 6 0 0 0-8.49-8.49L5.586 9.914A2 2 0 0 0 5 11.328V18a1 1 0 0 0 1 1z"))),
		x.Child(x.Path(x.D("M16 8 2 22"))),
		x.Child(x.Path(x.D("M17.5 15H9"))),
	)
	return x.Svg(svgArgs...)
}
