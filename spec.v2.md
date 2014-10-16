Base API
============================

### Common
__Content negotiation:__  
By default all data is sent and received as JSON. _"Content-Type"_ header should be set to `application/json`  

__Pagination:__  
Requests that return multiple items will be paginated to 25 items by default. You can control it with `start` and `limit` parameters. If `limit` is more than 100, default value will be set.  

__Errors:__  

- If item is not found, `200 OK` and empty result: `{}` or `[]` will be sent.
- Sending invalid JSON will result in a `400 Bad Request` response.
- Sending the wrong type of JSON values will result in a `400 Bad Request` response.
- Sending invalid fields will result in a `422 Unprocessable Entity` response.

Every 4XX response contains JSON with an error object like

```
{
	"errors": [
		{
			"classification": "DeserializationError",
			"message": "invalid character '\\\\' looking for beginning of value"
		}
	]
}
```

### /garments 
Get garments list. Create new garment.  

__Endpoint:__ `http://v2.dressformer.com/api/v2/garments`  

__Methods:__ 

- GET, POST

__Parameters:__

- `ids`   Coma separated list of ids.
- `start` Skip "n" records.
- `limit` Limit selection. Default value is 50.

__Result:__ Array of garment objects

e.g.:

```
GET /v2/garments?ids=93e92e72-1bdb-436f-bdb9-52dba6c16176

[
	{ 
		/* Garment Object */ 
	},
]

```

### /garments/:id
__Endpoint:__ `http://v2.dressformer.com/api/v2/garments/:id`  

__Methods:__ 
Get, update and delete certain garment.

- GET, PUT, DELETE

#####  Garment Model

```javascript
[
	{
		
		"id"        : "43a5dbed-4abb-452a-8f82-9a86e555930b",     // Garment Id as a uuid
		"gid"       : "43a5dbed-4abb-452a-8f82-123123123123",     // Group id
		"dummy_id"  : "0ae99696-0e13-4c54-8ad7-d1488dffbf65",
		"name"      : "Some name",
		"size_name" : "M",         // (String) Size, e.g. "XL", "42", etc...

		"sizes" : [
			{
				"id"        : "43a5dbed-4abb-452a-8f82-9a86e555930b",
				"size_name" : "S"
			},
			{
				"id"        : "33a6dbed-4abb-452a-8f82-1a16e666930a",
				"size_name" : "X"
			}  
		],

		"color" : "", // @todo		
		"slot"  : "some string",
		"layer" : 1.0,

		"url_prefix": "//v2.dressformer.com/assets/v2/",

		"assets" : {
			"geometry" : {
				"id"   : "22a6dbed-1aab-452b-8f81-2a16e994120b"
			},
			"placeholder" : {
				"id"      : "542d25c0000000000000000d"
			}			
		},
		
		"materials" : [
			{
				// material object
			},
			
			{
				// material object
			}
		]
		
	},	
]
```

### /user
__Endpoint:__ `http://v2.dressformer.com/api/v2/user`  


__Methods:__ 

- GET, POST, PUT

__Parameters:__


__Result:__ Object

##### User Model

```javascript
{
	"dummy": {
		"id": "0ae99696-0e13-4c54-8ad7-d1488dffbf65",
		"default": true,
		"assets": {
			"geometry": {
				"id" : "e748f388-36f8-47a2-b012-61f1083b80e7"
			}
		},
		
		"body": {
			"height"    : 174.0,
			"chest"     : 95.0,
			"underbust" : 72.0,
			"waist"     : 61.5,
			"hips"      : 89.0
		},
	},
	
	"history" : [ 
		{
			"id"       : "e748f388-36f8-47a2-b012-61f1083b80e7", 
			"gid"      : "43a5dbed-4abb-452a-8f82-123123123123",
			"dummy_id" : "0ae99696-0e13-4c54-8ad7-d1488dffbf65",
			"name"     : "Some name",
			
				// ... full garment model ...
		},
	],

	// Personal settings for objects. Control with /user/objects API
	/* Unused
	"objects" : [
		{
			"id" : "e748f388-36f8-47a2-b012-61f1083b80e7",
			"assets" : {
				"placeholder" : {
					"id"  : "53b54559fcb05d3238000002",
				}
			}
		},
		{
			"id"     : "f848f388-41f1-43b3-b012-31f1023b11e1",
			"assets" : {
				"placeholder" : {
					"id"  : "53b54559fcb05d3238000002",
				}
			}
		}		
	]
	*/
}

```

To get morphed mannequin, we should add all users body parameters to the dummy link, e.g.:

```
GET v2.dressformer.com/assets/v2/22a6dbed-1aab-452b-8f81-2a16e994120b?height=174.0&chest=95.0&underbust=72.0&waist=61.5&hips=89.0
```

#### Placeholders and History
@ todo description


### /dummies
Create and Get methods for dummies

__Endpoint:__ `http://v2.dressformer.com/api/v2/dummies`  


__Methods:__ 

- GET, POST

__Parameters:__


__Result:__ Array of Objects or Object for POST

##### Dummy Model

```json

{
	"id"         : "0ae99696-0e13-4c54-8ad7-d1488dffbf65",
	"name"       : "default dummy",
	"default"    : true,
	"url_prefix" : "//v2.dressformer.com/assets/v2/",  

	"assets": {
		"geometry": {
			"id"  : "e748f388-36f8-47a2-b012-61f1083b80e7"
		}
	},
	
	"body": {
		"chest"     : 91.154,
		"underbust" : 77.13,
		"waist"     : 66.71,
		"hips"      : 88.88,
		"height"    : 170
	}
}

```

### /dummies/:id
Update and Delete methods for dummy object

__Endpoint:__ `http://v2.dressformer.com/api/v2/dummies/:id`  


__Methods:__ 

- GET, PUT, DELETE


__Result:__ Dummy object  

__Example:__

```sh
curl -XPOST -H "Content-Type:application/json" -d '
{
	"name"    : "default dummy", 
	"default" : true, 
	"assets"  : {
		"geometry" : {
			"id" : "e748f388-36f8-47a2-b012-61f1083b80e7"
		}
	}
}' http://v2.dressformer.com/api/v2/dummies
```

Result:

```json
{
	"id"         : "0ae99696-0e13-4c54-8ad7-d1488dffbf65",
	"name"       : "default dummy",
	"default"    : true,
	"url_prefix" : "//v2.dressformer.com/assets/v2/",  
	"assets": {
		"geometry": {
			"id"  : "e748f388-36f8-47a2-b012-61f1083b80e7"
		}
	},
	"body": {}
}
```

### /materials
Create and Get methods for materials object

__Endpoint:__ `http://v2.dressformer.com/api/v2/materials`  


__Methods:__ 

- GET, POST

__Result:__ Array of objects  

__Expected POST data__: Array of objects


__Material Object__  

```json
{
    "id"   : "4717c27a-9fb0-4d09-ad70-2ae01f85deeb",
    "name" : "Test Material",
    "ka"   : "0.0435 0.0435 0.0435",
    "kd"   : "0.1086 0.1086 0.1086",
    "ks"   : "0.0000 0.0000 0.0000",
    "illum": 6,
    "d"    : "-halo 0.6600",
    "ns"   : "10.0000",
    "ni"   : "1.19713",
    
    "map_ka" : {
    	"id"       : "54390be70000000000000001",
    	"orig_name": "test_300.jpg",
    	"options"  : "-s 1 1 1 -o 0 0 0 -mm 0 1"
    },
    
    "map_kd" : {
    	"id"         : "54390be70000000000000001",
    	"orig_name"  : "test_300.jpg",
    	"options"    : "-s 1 1 1 -o 0 0 0 -mm 0 1"
    }
}
```

__1__. Create Material  

POST /materials `[ {some material data}, {some material data} ]` 

As a result, you will get an array of generated ids. E.g.

```
[
  "48e1f441-1c50-4f17-9c39-360cf5ecc99c",
  "fab465a2-19f0-4067-8216-b456a7742b8f"
]
```

__2__. Use these ids array as a value for `materials` field of a garment object, e.g.: 

```
PUT /garments/4717c27a-9fb0-4d09-ad70-2ae01f85deeb 
{
	"materials" : [
		"48e1f441-1c50-4f17-9c39-360cf5ecc99c",
		"fab465a2-19f0-4067-8216-b456a7742b8f"
	]
}
```

File API
============================

### /
_Upload files to asset._

__Endpoint:__ `http://v2.dressformer.com/assets/v2`  
  
__Methods:__

- POST

__Expected data:__

- MultiPart Form Data

__Example:__   
Following `curl` command uploads selected files 

```sh
curl                                                                                       \
	-i -XPOST -H "ContentType:multipart/form-data"                                         \
	-F name=Base.obj                         -F filedata=@Base.obj                         \
	-F name=Chest_max.obj                    -F filedata=@Chest_max.obj                    \
	-F name=KPL_201407_0020_0005_diffuse.jpg -F filedata=@KPL_201407_0020_0005_diffuse.jpg \
	-F name=KPL_201407_0020_0005_normal.jpg  -F filedata=@KPL_201407_0020_0005_normal.jpg  \
http://v2.dressformer.com/assets/
```
and returns

```json
[
	{
		"id"        : "53f622eb0000000000000001",
		"orig_name" : "Base.obj"
	},
	{
		"id"        : "53f622eb0000000000000002",
		"orig_name" : "Chest_max.obj"
	},
	{
		"id"        : "53f735c10000000000000001",
		"orig_name" : "KPL_201407_0020_0005_diffuse.jpg"
	},
	{
		"id"        : "53f735c10000000000000002",
		"orig_name" : "KPL_201407_0020_0005_normal.jpg"
	}	
]	
```

### /:id
_Get file from asset. Uniform method for every content types._

__Endpoint:__ `http://v2.dressformer.com/assets/v2`  
  
__Methods:__

- GET

__Expected parameters:__

- UUID  
	- __Accept:__
		- `application/octet-stream` default for GET method
		- `application/json` default for POST, PUT, DELETE mothods
	- __Result:__
		- `application/octet-stream` 
		- `application/json`
	- __Parameters:__
		- `height`    (float)
		- `chest`     (float)
		- `underbust` (float)
		- `waist`     (float)
		- `hips`      (float)

- ObjectId  
	__Parameters:__

	- `scale` Scaling image to dimensions  
		Prototype: `([0-9]+x) or (x[0-9]+) or ([0-9]+) or (0.[0-9]+)`  
		E.g.:  
  			+ `800x` scale to width 800px, height will be calculated  
		  	+ `x600` scale to height 600px, width will be calculated  
		  	+ `640`  maximum dimension is 640px, e.g. original 1024x768 pixel image will be scaled
  		  	   to 640x480, same option applied for 900x1600 image results 360x640  
		  	+ `0.5`  50% of original dimensions, e.g. 1024x768 = 512x384
	- `q` 0-100 image quality
	- `format` Image format — `png` or `jpg`. Jpeg is default one.
  
