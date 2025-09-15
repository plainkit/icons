package lucide

import x "github.com/plainkit/html"

// Wifi creates a Wifi Lucide icon.
func Wifi(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-wifi", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 20h.01"))),
		x.Child(x.Path(x.D("M2 8.82a15 15 0 0 1 20 0"))),
		x.Child(x.Path(x.D("M5 12.859a10 10 0 0 1 14 0"))),
		x.Child(x.Path(x.D("M8.5 16.429a5 5 0 0 1 7 0"))),
	)
	return x.Svg(svgArgs...)
}
