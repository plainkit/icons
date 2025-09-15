package lucide

import x "github.com/plainkit/blox"

// Boxes creates a Boxes Lucide icon.
func Boxes(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-boxes", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M2.97 12.92A2 2 0 0 0 2 14.63v3.24a2 2 0 0 0 .97 1.71l3 1.8a2 2 0 0 0 2.06 0L12 19v-5.5l-5-3-4.03 2.42Z"))),
		x.Child(x.Path(x.D("m7 16.5-4.74-2.85"))),
		x.Child(x.Path(x.D("m7 16.5 5-3"))),
		x.Child(x.Path(x.D("M7 16.5v5.17"))),
		x.Child(x.Path(x.D("M12 13.5V19l3.97 2.38a2 2 0 0 0 2.06 0l3-1.8a2 2 0 0 0 .97-1.71v-3.24a2 2 0 0 0-.97-1.71L17 10.5l-5 3Z"))),
		x.Child(x.Path(x.D("m17 16.5-5-3"))),
		x.Child(x.Path(x.D("m17 16.5 4.74-2.85"))),
		x.Child(x.Path(x.D("M17 16.5v5.17"))),
		x.Child(x.Path(x.D("M7.97 4.42A2 2 0 0 0 7 6.13v4.37l5 3 5-3V6.13a2 2 0 0 0-.97-1.71l-3-1.8a2 2 0 0 0-2.06 0l-3 1.8Z"))),
		x.Child(x.Path(x.D("M12 8 7.26 5.15"))),
		x.Child(x.Path(x.D("m12 8 4.74-2.85"))),
		x.Child(x.Path(x.D("M12 13.5V8"))),
	)
	return x.Svg(svgArgs...)
}
