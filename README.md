# SMRT Service


## API end points:**

### Add Line *localhost:8080/addline*
To add a new line to system 
```
Method = POST
RequestBody
{
    "line": {
        "id": "3",
        "name": "R",
        "stations": [
            {
                "id": "8",
                "name": "queenstown"
            },
            {
                "id": "2",
                "name": "vista"
            },
            {
                "id": "6",
                "name": "Dover"
            },
            {
                "id": "9",
                "name": "Dover"
            }
            
        ],
        "distances": [
            2,
            8,
            2
        ]
    }
}

Succees Response = {
    "v":"Added Successfully"
    }

ErrorResponse = {
    "v":"",
    "err": "Error Message"
}
```

### Search Between *localhost:8080/searhpath*

```
Method = POST
RequestBody = {
    "source": "1",
    "destination": "11",
    "criteria": "LEAST_TIME"
}

Success Response = {
    "path" : Array of paths
}
```




