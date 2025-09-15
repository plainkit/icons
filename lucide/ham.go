package lucide

import x "github.com/plainkit/html"

// Ham creates a Ham Lucide icon.
func Ham(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-ham", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M13.144 21.144A7.274 10.445 45 1 0 2.856 10.856"))),
		x.Child(x.Path(x.D("M13.144 21.144A7.274 4.365 45 0 0 2.856 10.856a7.274 4.365 45 0 0 10.288 10.288"))),
		x.Child(x.Path(x.D("M16.565 10.435 18.6 8.4a2.501 2.501 0 1 0 1.65-4.65 2.5 2.5 0 1 0-4.66 1.66l-2.024 2.025"))),
		x.Child(x.Path(x.D("m8.5 16.5-1-1"))),
	)
	return x.Svg(svgArgs...)
}
