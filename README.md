# Atlas

Atlas is a command-line tool that helps you initialize and develop your Express-Typescript applications.

## Installation

### Precompiled binaries

Precompiled binaries can be found in the [releases](https://github.com/MohamedBeydoun/atlas/releases) section of this project. Once the appropriate binary has been downloaded run the following commands:

**Linux and macOS**
```bash
chmod +x <downloaded_atlas_binary>
sudo mv <downloaded_atlas_binary> /usr/local/bin/atlas
atlas
```

**Windows**
```powershell
#Open powershell as admin
mv <path_to_downloaded_atlas_binary> C:\WINDOWS\system32\atlas.exe 
atlas
```

## Usage

**atlas help**

Provides a description and usage instructions for the provided command.

```bash
$ atlas help <command>
```
NOTE: The --help (-h) flag can be used with any command for the same effect.

**atlas create**

Creates a new express-typescript project.

```bash
$ atlas create <name> [options]
```

Arguments:
* name: name of the project

Options:
* --port, -p: the default port of the server
* --db-url: the default mongodb url

**atlas generate**

Generates a new resrouce of specified type. Currently, we support:
* model
* controller
* router

```bash
$ atlas generate <resource> [options]
```

**atlas generate model**

Generates the model, interface, and basic database interactions files for a mongodb model.

```bash
$ atlas generate model <name> [options]
```

Arguments:
* name: name of the model (`MUST` be singular lowecase)

Options:
* --fields, -f: a list of fields with their respective types e.g. name=string,toys=\[srting\] (surround type with brakets for an array of the type)

**atlas generate controller**

Generates the files for an express controller.

```bash
$ atlas generate controller <name> -f [options]
```

Arguments:
* name: name of the controller (`MUST` be singular lowecase)

Options:
* --functions, -f: a list of functions to be created in the controller e.g. index,create,show,delete

**atlas generate router**

Generates the files for an express router.

```bash
$ atlas generate router <name> [options]
```
NOTE: The user will have to manually choose the controller functions that go with each route by modifying the "CHANGE_ME" in the router file.
User will also have to use the router in their app.ts.

Arguments:
* name: name of the router (`MUST` be singular lowecase)

Options:
* --routes, -r: a list of routes with their http method and url e.g. post="/users",get="/users/:userId"

## License

Apache License 2.0. see [LICENSE](https://github.com/MohamedBeydoun/atlas/blob/master/LICENSE)