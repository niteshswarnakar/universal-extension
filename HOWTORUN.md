# To create universal-extension zipper

### Pre-requisites

- You need to have "extension.json" file with proper information in the extension-plugin repo at root level

### Build command

```bash
go build -o zipper
```

### Run command

- you need to copy this "zipper" executable file to the root level path of extension-plugin where your extension.json lies"

### Generate ZipFile 

- From the target extension-plugin repo, use this command

```bash
make build
```

- And then your zipfile "extension.zip" inside ./zippedfile is ready to be used as "extension" library in for ros-server