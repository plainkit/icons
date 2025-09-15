package lucide

import x "github.com/plainkit/blox"

// PointerOff creates a Pointer Off Lucide icon.
func PointerOff(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-pointer-off", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10 4.5V4a2 2 0 0 0-2.41-1.957"))),
		x.Child(x.Path(x.D("M13.9 8.4a2 2 0 0 0-1.26-1.295"))),
		x.Child(x.Path(x.D("M21.7 16.2A8 8 0 0 0 22 14v-3a2 2 0 1 0-4 0v-1a2 2 0 0 0-3.63-1.158"))),
		x.Child(x.Path(x.D("m7 15-1.8-1.8a2 2 0 0 0-2.79 2.86L6 19.7a7.74 7.74 0 0 0 6 2.3h2a8 8 0 0 0 5.657-2.343"))),
		x.Child(x.Path(x.D("M6 6v8"))),
		x.Child(x.Path(x.D("m2 2 20 20"))),
	)
	return x.Svg(svgArgs...)
}
