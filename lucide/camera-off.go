package lucide

import x "github.com/plainkit/html"

// CameraOff creates a Camera Off Lucide icon.
func CameraOff(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-camera-off", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M14.564 14.558a3 3 0 1 1-4.122-4.121"))),
		x.Child(x.Path(x.D("m2 2 20 20"))),
		x.Child(x.Path(x.D("M20 20H4a2 2 0 0 1-2-2V9a2 2 0 0 1 2-2h1.997a2 2 0 0 0 .819-.175"))),
		x.Child(x.Path(x.D("M9.695 4.024A2 2 0 0 1 10.004 4h3.993a2 2 0 0 1 1.76 1.05l.486.9A2 2 0 0 0 18.003 7H20a2 2 0 0 1 2 2v7.344"))),
	)
	return x.Svg(svgArgs...)
}
