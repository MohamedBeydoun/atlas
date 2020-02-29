package tpl

// ConsoleTemplate ...
func ConsoleTemplate() []byte {
	return []byte(`const App =  require("./dist/app.js").app
const mongoose = require("mongoose")

{{ range $model := . }}
const {{ $model | ToTitle }} = require("./dist/database/interactions/{{ $model }}.js").{{ $model }}DBInteractions

{{ end }}

mongoose.connect("mongodb://localhost:27017/test", {
    useNewUrlParser: true,
    useFindAndModify: false,
    useUnifiedTopology: true,
    useCreateIndex: true
});
mongoose.set("useCreateIndex", true);
`)
}
