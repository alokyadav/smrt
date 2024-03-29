# SMRT Service
Design SMRT system as a web application where you need to support two APIs:
* As admin, I should be able to add line with unique line-id and stations on that line with unique station-id along with travel time between each station on line.
* As user, I should be able to search optimal path between source and destination stations with one of the below optimality criteria - quickest travel time, least line switches, least number of stations



## How To Run This Project
```bash
git clone https://github.com/alokyadav/smrt.git

docker-compose up db

#Open new tab

docker-compose build smrt-service

docker-compose up smrt-service


```

## Note
* Install docker
* Install docker compose


## API end points:

### Add Line 
To add a new line to system 
*localhost:8080/addline*

```
Method = POST
RequestBody
{
    "line": {
        "id": "3", #string
        "name": "R", # strig 
        "stations": [
            {
                "id": "8", #string
                "name": "queenstown" # string
            },
            {
                "id": "2", #string
                "name": "vista" #string
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
        "times": [
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
    "source": "1", #string
    "destination": "11", #string
    "criteria": "LEAST_TIME" #Possible values - LEAST_TIME, LEAST_SWITCH and LEAST_STATION_NUMBER
}

Success Response = {
    "path" : Array of paths containing station ids
}
```


## Future Improvement
* Use external Cache (Redis etc ) to store the searh data structures .
* Separate out both apis into two independent service as search api will have more call as compare to add line.

