package tpl

// ControllerTemplate ...
func ControllerTemplate() []byte {
	return []byte(`import { Request, Response } from "express";
import { statusCodes } from "../util/statusCodes";

const {{ .Name }}Controller = {

};

export { {{ .Name }}Controller };
`)
}
