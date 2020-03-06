package tpl

// ModelTemplate ...
func ModelTemplate() []byte {
	return []byte(`import { Schema, Document, Model, model } from "mongoose";
import { I{{ .Name | ToTitle }} } from "../../interfaces/I{{ .Name | ToTitle }}";

export interface I{{ .Name | ToTitle }}Model extends I{{ .Name | ToTitle }}, Document {}

const {{ .Name }}Schema: Schema = new Schema({
{{ range $field, $type := .Fields }}{{ $isArray:=index $type 0 }}    {{ $field }}: {{ if eq $isArray 91 }}[{
        type: {{ $type | TrimLeft | TrimLeft | ToTitle }}
    }],
{{else}}{
        type: {{ $type | ToTitle }}
    },
{{ end }}{{ end }}});

const {{ .Name | ToTitle }}: Model<I{{ .Name | ToTitle }}Model> = model<I{{ .Name | ToTitle }}Model>("{{ .Name | ToTitle }}", {{ .Name }}Schema);

export { {{ .Name | ToTitle }} };
`)
}

// InterfaceTemplate ...
func InterfaceTemplate() []byte {
	return []byte(`export interface I{{ .Name | ToTitle }} {
{{ range $field, $type := .Fields }}    {{ $isArray:=index $type 0 }}{{ if eq $isArray 91 }}{{ $field }}: {{ $type | TrimLeft | TrimLeft }}[]{{ else }}{{ $field }}: {{ $type}}{{ end }};
{{ end }}}
`)
}

// InteractionsTemplate ...
func InteractionsTemplate() []byte {
	return []byte(`import { I{{ .Name | ToTitle }} } from "../../interfaces/I{{ .Name | ToTitle }}";
import { {{ .Name | ToTitle }}, I{{ .Name | ToTitle }}Model } from "../models/{{ .Name }}";

export const {{ .Name }}DBInteractions = {

    create: ({{ .Name }}: I{{ .Name | ToTitle }}): Promise<I{{ .Name | ToTitle }}Model> => {
        return {{ .Name | ToTitle }}.create({{ .Name }});
    },

    all: (): Promise<I{{ .Name | ToTitle }}Model[]> => {
        return {{ .Name | ToTitle }}.find().exec();
    },

    find: ({{ .Name }}Id: string): Promise<I{{ .Name | ToTitle }}Model> => {
        return {{ .Name | ToTitle }}.findOne({ _id: {{ .Name }}Id }).exec();
    },

    update: ({{ .Name }}Id: string, new{{ .Name | ToTitle }}: I{{ .Name | ToTitle }}): Promise<I{{ .Name | ToTitle }}Model> => {
        return {{ .Name | ToTitle }}.findByIdAndUpdate({{ .Name }}Id, new{{ .Name | ToTitle }}, { new: true }).exec();
    },

    delete: ({{ .Name }}Id: string): Promise<I{{ .Name | ToTitle }}Model> => {
        return {{ .Name | ToTitle }}.findByIdAndDelete({{ .Name }}Id).exec();
    },
};
`)
}
