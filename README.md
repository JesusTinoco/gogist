# gogist

**gogist** is a command line tool that allows you to upload content to gists ([https://gist.github.com/](https://gist.github.com/)).

```console
→ gogist --help

                   _     _
                  (_)   | |
  __ _  ___   __ _ _ ___| |_
 / _  |/ _ \ / _  | / __| __|
| (_| | (_) | (_| | \__ \ |_
 \__, |\___/ \__, |_|___/\__|
  __/ |       __/ |
 |___/       |___/

 Command Line Tool to upload contents to Gist (https://gist.github.com/)
 Version: v0.0.1

  -d description
        Adds a description to the gist
  -help
        Help menu
  -l    Authenticate to GitHub
  -p    Sets the gist private
```

## Usage

To upload the content of a file:

`gogist file.go`

Also, you can upload the content of several files at the same time:

`gogist file1.go file2.go file3.go`

Use the `-p` option to make the gist private:

`gogist -p file.go`

Use the `-d` option to add a description:

`gogist -d "This is a description" file.go`

## Login

Rigth now, you only can create gists associated to your GitHub account, you before anything else, you will have to login. To do that, execute the following command:

```console
→ gogist -l

Obtaining OAuth2 access token from GitHub.
Enter GitHub username: JesusTinoco
Enter Github password: **************
Token created successfully. It will be safe at ~/gogist
Token saved successfully in  /Users/jesusrodrigueztinoco/.gogist
```
