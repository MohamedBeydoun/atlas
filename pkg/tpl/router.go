package tpl

// RouterTemplate ...
func RouterTemplate() []byte {
	return []byte(`import { Router } from "express";
import { {{ .Name }}Controller } from "../controllers/{{ .Name }}";

const {{ .Name }}Router: Router = Router();
{{ range $method, $url := .Routes }}
{{ $.Name }}Router.{{ $method }}("{{ $url }}", {{ $.Name }}Controller["CHANGE_ME"]);
{{ end }}
export { {{ .Name }}Router };
`)
}
