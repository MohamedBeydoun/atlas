package tpl

// RouterTemplate ...
func RouterTemplate() []byte {
	return []byte(`import { Router } from "express";
import { {{ .Name }}Controller } from "../controllers/{{ .Name }}";

const {{ .Name }}Router: Router = Router();

export { {{ .Name }}Router };
`)
}
