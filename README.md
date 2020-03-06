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

**atlas console**

Loads the express application into a node console for easier debugging.

```bash
atlas console [options]
```

Options:
* --db-url: the mongodb url to which the console will connect

**atlas generate**

Generates a new resrouce of specified type. Currently, we support:
* model
* router
* route

```bash
$ atlas generate <resource> [options]
```

NOTE: Preferably, resource names should be singular as the cli will change to plural as needed.

**atlas generate model**

Generates the model, interface, and basic database interactions files for a mongodb model.

```bash
$ atlas generate model <name> [options]
```

Arguments:
* name: name of the model

Options:
* --fields, -f: a list of fields with their respective types e.g. name=string,toys=srting\[\] (this flag can be used repeatedly instead of being comma separated e.g. -f name=string -f toys=string\[\])

**atlas generate router**

Generates the files for an express router along with it's controller.

```bash
$ atlas generate router <name> [options]
```

Arguments:
* name: name of the router

**atlas generate route**

Populates the router and controller files with the given route information.

```bash
$ atlas generate route [options]
```

Options:
* --router, -r: the name of the router to which this route is associated
* --method, -m: HTTP method for the route e.g. get, post, etc
* --url, -u: the route's endpoint
* --controller, -c: the name of the controller function associated with this route e.g. index, show, delete, etc

## License

Apache License 2.0. see [LICENSE](https://github.com/MohamedBeydoun/atlas/blob/master/LICENSE)
