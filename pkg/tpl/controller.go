package tpl

// ControllerTemplate ...
func ControllerTemplate() []byte {
	return []byte(`import { Request, Response } from "express";

const {{ .Name }}Controller = {
{{ range $function := .Functions }}
    {{ $function }}: async (req: Request, res: Response) => {
        try {
            res.status(500).send({ msg: "Not Implemented" });
        } catch (error) {
            res.status(500).send(error);
        }
    },
{{ end }}};

export { {{ .Name }}Controller };
`)
}
