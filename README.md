# Quago

A super lightweight JSON configuration tool

### Download package

Download using the following command:
```
$ go get github.com/agilasolutions/quago

```

### Set Configuration

Set environment variable to determine which config to reference (dev/uat/prod):
```
quago.SetEnvironment("") // default will lookup to dev.json
```
Set config file format (current only support .json):
```
quago.SetFormat(".json")

```
Set path where config file is located:
```
quago.SetPath("/configs")
```
### Get Configuration Value

Then finally, source the config file and get value by key:
```
config, _ := quago.Source()
config.Get("URL") // Get value where key is equal to URL
