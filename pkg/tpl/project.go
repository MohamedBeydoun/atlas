package tpl

// ReadmeTemplate ...
func ReadmeTemplate() []byte {
	return []byte(`# {{ .Name }}

`)
}

// GitignoreTemplate ...
func GitignoreTemplate() []byte {
	return []byte(`node_modules
.nyc_output
dist/
`)
}

// PackageJSONTemplate ...
func PackageJSONTemplate() []byte {
	return []byte(`{
    "name": "{{ .Name }}",
    "version": "1.0.0",
    "description": "",
    "main": "server.ts",
    "scripts": {
        "start": "nodemon --watch 'src/**/*.ts' --exec 'ts-node' src/server.ts",
        "build": "tsc -p tsconfig.json",
        "watch": "tsc -w",
        "lint": "tslint -c tslint.json -p tsconfig.json"
    },
    "author": "",
    "license": "ISC",
    "devDependencies": {
        "@babel/preset-typescript": "^7.3.3",
        "nodemon": "^1.19.0",
        "source-map-support": "^0.5.12",
        "ts-node": "^8.1.0"
    },
    "dependencies": {
        "@types/cors": "^2.8.6",
        "@types/express": "^4.16.1",
        "@types/mongodb": "^3.1.26",
        "@types/mongoose": "^5.5.0",
        "@types/node": "^11.13.10",
        "@babel/core": "^7.0.0",
        "body-parser": "^1.19.0",
        "cors": "^2.8.5",
        "express": "^4.16.4",
        "mongoose": "^5.5.7",
        "tslint": "^5.16.0",
        "typescript": "^3.4.5"
    }
}
`)
}

// TSLintTemplate ...
func TSLintTemplate() []byte {
	return []byte(`{
    "rules": {
        "class-name": true,
        "comment-format": [
            true,
            "check-space"
        ],
        "indent": [
            true,
            "spaces"
        ],
        "one-line": [
            true,
            "check-open-brace",
            "check-whitespace"
        ],
        "no-var-keyword": true,
        "quotemark": [
            true,
            "double",
            "avoid-escape"
        ],
        "semicolon": [
            true,
            "always",
            "ignore-bound-class-methods"
        ],
        "whitespace": [
            true,
            "check-branch",
            "check-decl",
            "check-operator",
            "check-module",
            "check-separator",
            "check-type"
        ],
        "typedef-whitespace": [
            true,
            {
                "call-signature": "nospace",
                "index-signature": "nospace",
                "parameter": "nospace",
                "property-declaration": "nospace",
                "variable-declaration": "nospace"
            },
            {
                "call-signature": "onespace",
                "index-signature": "onespace",
                "parameter": "onespace",
                "property-declaration": "onespace",
                "variable-declaration": "onespace"
            }
        ],
        "no-internal-module": true,
        "no-trailing-whitespace": true,
        "prefer-const": true
    }
}
`)
}

// TSConfigTemplate ...
func TSConfigTemplate() []byte {
	return []byte(`{
    "compilerOptions": {
        "allowSyntheticDefaultImports": true,
        "sourceMap": true,
        "esModuleInterop": true,
        "module": "commonjs",
        "baseUrl": "./src",
        "rootDir": "src",
        "outDir": "dist",
        "resolveJsonModule": true,
        "target": "es6",
        "lib": [
            "esnext",
            "es2018",
            "dom"
        ]
    },
    "include": [
        "src/**/*.ts"
    ],
    "exclude": [
        "node_modules",
    ]
}
`)
}

// AppTemplate ...
func AppTemplate() []byte {
	return []byte(`import express from "express";
import { Application, Request, Response } from "express";
import bodyParser from "body-parser";
import cors from "cors";

export const port: Number = parseInt(process.env.PORT) || {{ .Port }};
const app: Application = express();

app.use(cors());
app.use(bodyParser.json());

app.use((req: Request, res: Response) => {
    res.status(404).send({
        status: 404,
        message: "Invalid route"
    });
});

export { app };
`)
}

// ServerTemplate ...
func ServerTemplate() []byte {
	return []byte(`import { app, port } from "./app";
import mongoose from "mongoose";

let dbUrl = "";

(process.env.DB_URL)
    ? dbUrl = process.env.DB_URL
    : dbUrl = "{{ .DBURL }}";

mongoose.connect(dbUrl, {
    useNewUrlParser: true,
    useFindAndModify: false,
    useUnifiedTopology: true,
    useCreateIndex: true
});
mongoose.set("useCreateIndex", true);

const server = app.listen(port, async () => {
    console.log("Server listening on port " + port);
});

export { server };
`)
}
