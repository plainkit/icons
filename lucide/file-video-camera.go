package lucide

import (
	html "github.com/plainkit/html"
)

// FileVideoCamera creates a File Video Camera Lucide icon.
func FileVideoCamera(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-file-video-camera", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M4 22h14a2 2 0 0 0 2-2V7l-5-5H6a2 2 0 0 0-2 2v4"))),
		html.Child(html.SvgPath(html.AD("M14 2v4a2 2 0 0 0 2 2h4"))),
		html.Child(html.SvgRect(html.AWidth("8"), html.AHeight("6"), html.AX("2"), html.AY("12"), html.ARx("1"))),
		html.Child(html.SvgPath(html.AD("m10 13.843 3.033-1.755a.645.645 0 0 1 .967.56v4.704a.645.645 0 0 1-.967.56L10 16.157"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
