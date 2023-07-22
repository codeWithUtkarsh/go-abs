# Go-Abs

Go-Abs is a Golang Database Abstraction layer that provides plug and play support of any new database implementation.
Here is code for single implementation(NFT Storage):
https://github.com/codeWithUtkarsh/go-abs

Here is code for multiple implementation(Postgres and NFT Storage):-
https://github.com/codeWithUtkarsh/go-abs/tree/chore/multi-implementation

## How implementation is decided on Application startup?

The choice of implementation is determines by this environment variable "DATA_SOURCE". If you have set this environment variable to "NFT", the NFT implemntation will be considered. The availaible values for multi-source implemnetation are "postgres" and "NFT".

# Run your Code locally with NFT Storage
## Run your Code without creating image
- Set environment variable `DATA_SOURCE` as `NFT`.
- Set additional environment variable like `ENDPOINT` and `API_KEY`
- Run application using `go run .\main.go`. (NOTE: Make sure you are in go-abs directory)

## Run your Code locally with NFT Storage using Docker
- Build docker file using `docker build -t go-abs .`
- Run image created by above command using `docker run --env=GOLANG_VERSION=1.18.10 --env=GOPATH=/go --env=DATA_SOURCE=NFT --env=ENDPOINT=https://api.nft.storage --env=API_KEY=eJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJkaWQ6ZXRocjoweDk3ODMwMGIzNUMzOEM3NzMxYWNDNjk4NDFiODU4NDNiRmIxRjExN0QiLCJpc3MiOiJuZnQtc3RvcmFnZSIsImlhdCI6MTY0ODY2MjM1MDcwNiwibmFtZSI6Im9ydGVsaXVzX3Nib21fbGVkZ2VyX2RlbW8ifQ.egQzHqDB83UoK7ynE4fn4dhgIKWLhPP1B9E6qey4KHM -p 8080:10000 -d go-abs`

## Run your Code locally with PostgreSQL using Docker Compose
- Please make sure pre-requisites are done and verified.
- Run docker compose file using `docker-compose -f .\docker-compose-with-nft.yml up`
- 
## Testing your Application 


API: `/status`
Mandatory Parameter: `cid`
Sample Request:- 
```json
curl --location --request GET 'http://localhost:8080/status?cid=bafkreiajt4fghhgqblhulr5f45zp4wfowmyrhgvgecxenknrfmtkkgg7d4' \
--header 'Content-Type: application/json'
```
Sample Response:- 
```json
{
    "result": "success",
    "metadata": {
        "ok": true,
        "value": {
            "cid": "bafkreihn7cjn5gxhtdwp4ufjqdpozqxlhtcgv6z5ww5swivq22ix46x7uy",
            "deals": [
                {
                    "batchRootCid": "bafybeiar4kr7mmxszipg2wp5ogjsccsl6z7pf6du6loc6ybtid6mvykmve",
                    "chainDealID": 44394724,
                    "datamodelSelector": "Links/216/Hash/Links/689/Hash/Links/0/Hash",
                    "dealActivation": "2023-07-06T07:16:00+00:00",
                    "dealExpiration": "2024-12-19T07:16:00+00:00",
                    "lastChanged": "2023-07-22T12:20:03.549937+00:00",
                    "miner": "f097777",
                    "pieceCid": "baga6ea4seaql7fh3m3f6h62a2tuekysmgjfgzycrlueoeisqtmvnzb2on3iosba",
                    "status": "active",
                    "statusText": null
                },
                {
                    "batchRootCid": "bafybeiar4kr7mmxszipg2wp5ogjsccsl6z7pf6du6loc6ybtid6mvykmve",
                    "chainDealID": 42497455,
                    "datamodelSelector": "Links/216/Hash/Links/689/Hash/Links/0/Hash",
                    "dealActivation": "2023-06-24T14:48:30+00:00",
                    "dealExpiration": "2024-12-07T14:48:30+00:00",
                    "lastChanged": "2023-07-22T12:20:03.549937+00:00",
                    "miner": "f01163272",
                    "pieceCid": "baga6ea4seaql7fh3m3f6h62a2tuekysmgjfgzycrlueoeisqtmvnzb2on3iosba",
                    "status": "active",
                    "statusText": null
                },
                {
                    "batchRootCid": "bafybeiar4kr7mmxszipg2wp5ogjsccsl6z7pf6du6loc6ybtid6mvykmve",
                    "chainDealID": 41800472,
                    "datamodelSelector": "Links/216/Hash/Links/689/Hash/Links/0/Hash",
                    "dealActivation": "2023-06-20T10:20:30+00:00",
                    "dealExpiration": "2024-12-03T10:20:30+00:00",
                    "lastChanged": "2023-07-22T12:20:03.549937+00:00",
                    "miner": "f0717969",
                    "pieceCid": "baga6ea4seaql7fh3m3f6h62a2tuekysmgjfgzycrlueoeisqtmvnzb2on3iosba",
                    "status": "active",
                    "statusText": null
                },
                {
                    "batchRootCid": "bafybeiar4kr7mmxszipg2wp5ogjsccsl6z7pf6du6loc6ybtid6mvykmve",
                    "chainDealID": 41419453,
                    "datamodelSelector": "Links/216/Hash/Links/689/Hash/Links/0/Hash",
                    "dealActivation": "2023-06-17T07:26:00+00:00",
                    "dealExpiration": "2024-11-30T07:26:00+00:00",
                    "lastChanged": "2023-07-22T12:20:03.549937+00:00",
                    "miner": "f01137229",
                    "pieceCid": "baga6ea4seaql7fh3m3f6h62a2tuekysmgjfgzycrlueoeisqtmvnzb2on3iosba",
                    "status": "active",
                    "statusText": null
                },
                {
                    "batchRootCid": "bafybeiar4kr7mmxszipg2wp5ogjsccsl6z7pf6du6loc6ybtid6mvykmve",
                    "chainDealID": 41317941,
                    "datamodelSelector": "Links/216/Hash/Links/689/Hash/Links/0/Hash",
                    "dealActivation": "2023-06-17T08:54:00+00:00",
                    "dealExpiration": "2024-11-30T08:54:00+00:00",
                    "lastChanged": "2023-07-22T12:20:03.549937+00:00",
                    "miner": "f01771403",
                    "pieceCid": "baga6ea4seaql7fh3m3f6h62a2tuekysmgjfgzycrlueoeisqtmvnzb2on3iosba",
                    "status": "active",
                    "statusText": null
                },
                {
                    "batchRootCid": "bafybeiar4kr7mmxszipg2wp5ogjsccsl6z7pf6du6loc6ybtid6mvykmve",
                    "chainDealID": 40750482,
                    "datamodelSelector": "Links/216/Hash/Links/689/Hash/Links/0/Hash",
                    "dealActivation": "2023-06-13T18:45:00+00:00",
                    "dealExpiration": "2024-11-26T18:45:00+00:00",
                    "lastChanged": "2023-07-22T12:20:03.549937+00:00",
                    "miner": "f02095132",
                    "pieceCid": "baga6ea4seaql7fh3m3f6h62a2tuekysmgjfgzycrlueoeisqtmvnzb2on3iosba",
                    "status": "active",
                    "statusText": null
                }
            ],
            "pin": {
                "cid": "bafkreihn7cjn5gxhtdwp4ufjqdpozqxlhtcgv6z5ww5swivq22ix46x7uy",
                "created": "2023-06-06T17:40:29.11+00:00",
                "size": 27,
                "status": "pinned"
            }
        }
    }
}
```

API: `/upload`
Body:
NOTE: The Object inside payload will be saved as nft object.
```json
{
  "payload": {
    "username": "utkarshSharma",
    "email": "utkarsh@gmail.com"
  }
}
```
Sample Request:- 
```curl
curl -X 'POST' \
  'http://localhost:8080/upload' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{"payload":{"username":"utkarshSharma","email":"utkarsh@gmail.com"}}'
```
Sample Response:-
```json
{
  "result": "success",
  "metadata": {
    "ok": true,
    "value": {
      "cid": "bafkreiajt4fghhgqblhulr5f45zp4wfowmyrhgvgecxenknrfmtkkgg7d4",
      "created": "2023-07-22T12:17:31.15+00:00",
      "deals": [],
      "files": [],
      "name": "Upload at 2023-07-22T12:17:31.150Z",
      "pin": {
        "cid": "bafkreiajt4fghhgqblhulr5f45zp4wfowmyrhgvgecxenknrfmtkkgg7d4",
        "created": "2023-07-22T12:17:31.15+00:00",
        "size": 68,
        "status": "pinned"
      },
      "scope": "ortelius_sbom_ledger_demo",
      "size": 68,
      "type": ""
    }
  }
}
```

# Run your Code locally with PostgreSQL

## Pre-requisites
- Download a database client software to correct to db locally. eg:- https://dbeaver.io/download/
- Create Database that you want to use using `CREATE database db;` where db is my database name.
- Make sure you are using the correct database using this command `select current_database()`.
- Create tables for your models entities.
  eg- `CREATE TABLE public.account (
    id SERIAL,
    username varchar(50) NOT NULL,
    email varchar(50) NOT NULL,
    PRIMARY KEY (id)
)` 
- Verify if the table is created properly using `SELECT * from public.account`.
  
## Run your Code locally with PostgreSQL
- Set environment variable `DATA_SOURCE` as `postgres`.
- Set additional environment variable like `DB_HOST`, `DB_PORT`, `DB_USERNAME`,and `DB_PASSWORD`
- Run application using `go run .\main.go`. (NOTE: Make sure you are in go-abs directory)
  
## Run your Code locally with PostgreSQL using Docker Compose
- Please make sure pre-requisites are done and verified.
- Run docker compose file using `docker-compose -f .\docker-compose-with-postgresql.yml up`


## Testing your Application 

API: `/status`
Mandatory Parameter: `cid`
Sample Request:- NA for Relational Database
Sample Response:- NA for Relational Database 


API: `/upload`
Body:
```json
{
  "entity": "Account",
  "payload": {
    "username": "utkarshSharma",
    "email": "utkarsh@gmail.com"
  }
}
```
Sample Request:- 
```curl
curl -X 'POST' \
  'http://localhost:8080/upload' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{"entity":"Account","payload":{"username":"utkarshSharma","email":"utkarsh@gmail.com"}}'
```
Sample Response:-
```json
{
  "result": "success",
  "metadata": {
    "message": "The User has been inserted successfully! Id: 2"
  }
}
```