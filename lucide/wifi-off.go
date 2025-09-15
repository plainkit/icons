package lucide

import x "github.com/plainkit/html"

// WifiOff creates a Wifi Off Lucide icon.
func WifiOff(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-wifi-off", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 20h.01"))),
		x.Child(x.Path(x.D("M8.5 16.429a5 5 0 0 1 7 0"))),
		x.Child(x.Path(x.D("M5 12.859a10 10 0 0 1 5.17-2.69"))),
		x.Child(x.Path(x.D("M19 12.859a10 10 0 0 0-2.007-1.523"))),
		x.Child(x.Path(x.D("M2 8.82a15 15 0 0 1 4.177-2.643"))),
		x.Child(x.Path(x.D("M22 8.82a15 15 0 0 0-11.288-3.764"))),
		x.Child(x.Path(x.D("m2 2 20 20"))),
	)
	return x.Svg(svgArgs...)
}
