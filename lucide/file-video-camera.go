package lucide

import x "github.com/plainkit/html"

// FileVideoCamera creates a File Video Camera Lucide icon.
func FileVideoCamera(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-file-video-camera", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M4 22h14a2 2 0 0 0 2-2V7l-5-5H6a2 2 0 0 0-2 2v4"))),
		x.Child(x.Path(x.D("M14 2v4a2 2 0 0 0 2 2h4"))),
		x.Child(x.Rect(x.RectWidth("8"), x.RectHeight("6"), x.X("2"), x.Y("12"), x.Rx("1"))),
		x.Child(x.Path(x.D("m10 13.843 3.033-1.755a.645.645 0 0 1 .967.56v4.704a.645.645 0 0 1-.967.56L10 16.157"))),
	)
	return x.Svg(svgArgs...)
}
