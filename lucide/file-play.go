package lucide

import x "github.com/plainkit/html"

// FilePlay creates a File Play Lucide icon.
func FilePlay(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-file-play", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M14 2v4a2 2 0 0 0 2 2h4"))),
		x.Child(x.Path(x.D("M15 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7z"))),
		x.Child(x.Path(x.D("M15.033 13.44a.647.647 0 0 1 0 1.12l-4.065 2.352a.645.645 0 0 1-.968-.56v-4.704a.645.645 0 0 1 .967-.56z"))),
	)
	return x.Svg(svgArgs...)
}
