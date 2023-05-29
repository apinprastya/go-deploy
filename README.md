# Go-Deploy

## Features
- Upload binary app to server
- Start app on server

## Actions
A set of action that need to be run when calling the go-deploy. Available actions:
### Set Live
Command set live json:
```
{
    "command": "setLive",
    "options": {
        "excludedFolders": [],
        "excludedFiles": []   
    }
}
```
When set any version to live, it will copy all data to production folder. It possible that on production has generated folder like logs that we don't need to sync, just add the folder to `excludedFolders`
### Run
Command run will have json:
```
{
    "command": "run",
    "options": {
        "executable": "your app executable",
        "envs": {
            "ANY_ENVIRONTMENT": "VARIABLES"
        }
        "args": ["--production"]
    }
}
```
### Restart
Command restart will have json:
```
{
    "command": "restart"
}
```
Restart command will stop the last run command and start again
### Stop
Command stop:
```
{
    "command": "stop"
}
```

## Client App
Client App will need a valid key to authenticate with backend

### Client Config
Client App will find the config.yaml for read the config. The yaml value will look like below:
```
SECRET_KEY: [your server key]
PROJECT_NAME: [your project name]
```

### Upload
```
godeploy upload --folder [folder to sync to server] --version [version number] --folderExclude [folder to exclude with comma separated]
```

### Action
Action will need an app config to be run
#### App Config
When uploading a folder to server using `Client App`, make sure to have the `App Config`, this config consist of named commands. Example of config in json:
```
{
    "deployAndRun": [
        {
            "command": "setLive",
            "description": "set the application to live",
            "options": {
                "excludedFolders": ["logs"],
                "excludedFiles": ["log.log"]
            }
        },
        {
            "command": "run",
            "description": "database migrations",
            "options": {
                "executable": "migration",
                "envs": {
                    "ANY_ENVIRONTMENT": "VARIABLES"
                }
                "args": ["--folder", "myfolder"]
            }
        },
        {
            "command": "run",
            "description": "run the application",
            "options": {
                "executable": "yourAppName",
                "envs": {
                    "ANY_ENVIRONTMENT": "VARIABLES"
                }
                "args": ["--production"]
            }
        }
    ],
    "restart": [
        {
            "command": "restart"
        }
    ]
}
```

After has the `App Config` in the uploaded folder, then the action command can be run:
```
godeploy action --version [your version] --action deployAndRun
godeploy action --version [your version] --action restart
```