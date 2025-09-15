package lucide

import x "github.com/plainkit/blox"

// TrainFrontTunnel creates a Train Front Tunnel Lucide icon.
func TrainFrontTunnel(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-train-front-tunnel", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M2 22V12a10 10 0 1 1 20 0v10"))),
		x.Child(x.Path(x.D("M15 6.8v1.4a3 2.8 0 1 1-6 0V6.8"))),
		x.Child(x.Path(x.D("M10 15h.01"))),
		x.Child(x.Path(x.D("M14 15h.01"))),
		x.Child(x.Path(x.D("M10 19a4 4 0 0 1-4-4v-3a6 6 0 1 1 12 0v3a4 4 0 0 1-4 4Z"))),
		x.Child(x.Path(x.D("m9 19-2 3"))),
		x.Child(x.Path(x.D("m15 19 2 3"))),
	)
	return x.Svg(svgArgs...)
}
