## Base API

### /garments 
Get garments list. Create new garment.  

__Endpoint:__ `http://v2.dressformer.com/api/garments`  

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
__Endpoint:__ `http://v2.dressformer.com/api/garments/:idv2.`  

__Methods:__ 
Update and delete garment with `id`

- PUT, DELETE


#####  Garment Model

```javascript
[
	{
		
		"id"    : "43a5dbed-4abb-452a-8f82-9a86e555930b", // Garment Id as a uuid

		"gid"   : "43a5dbed-4abb-452a-8f82-123123123123", // Group id
				
		"dummy_id" : "0ae99696-0e13-4c54-8ad7-d1488dffbf65",
		
		"name"  : "Some name",
		
		"size_name"  : "M", // (String) Size, e.g. "XL", "42", etc...

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
		
		"assets" : {
			"geometry" : "assets.dressformer.com/geometry/22a6dbed-1aab-452b-8f81-2a16e994120b",
			
			"diffuse"  : "assets.dressformer.com/image/53b54559fcb05d3238000002",
			
			"normal"   : "assets.dressformer.com/image/53b61050eff01c1008000001",
			
			"specular" : "assets.dressformer.com/image/53b61050eff01c1008000003"
		}		
	},	
]
```

### /user
__Endpoint:__ `http://v2.dressformer.com/api/user`  


__Methods:__ 

- GET, POST, PUT

__Parameters:__


__Result:__ Object

##### User Model

Guest user model

```javascript
{
	// default manequin
	"dummy" : "v2.dressformer.com/assets/geometry/22a6dbed-1aab-452b-8f81-2a16e994120b"
}

```

Authorized user contains body settings

```javascript
{
	// base dummy
	"dummy" : "v2.dressformer.com/assets/geometry/22a6dbed-1aab-452b-8f81-2a16e994120b"

	// Body object contains only those parameters, which are different from the base one.	
	"body" : {
		"height"    : 174.0,
		"chest"     : 95.0,
		"underbust" : 72.0,
		"waist"     : 61.5,
		"hips"      : 89.0
	}
}
```

To get morphed mannequin, we should add all users body parameters to the dummy link, e.g.:

```
GET v2.dressformer.com/assets/geometry/22a6dbed-1aab-452b-8f81-2a16e994120b?height=174.0&chest=95.0&underbust=72.0&waist=61.5&hips=89.0
```

### /dummies
Create and Get methods for dummies

__Endpoint:__ `http://v2.dressformer.com/api/dummies`  


__Methods:__ 

- GET, POST

__Parameters:__


__Result:__ Array of Objects or Object for POST

##### Dummy Model

```json

{
	"id": "0ae99696-0e13-4c54-8ad7-d1488dffbf65",
	
	"name": "default dummy",
	
	"default": true,
	
	"assets": {
		"geometry": "//localhost:6500/geometry/e748f388-36f8-47a2-b012-61f1083b80e7"
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

__Endpoint:__ `http://v2.dressformer.com/api/dummies/:id`  


__Methods:__ 

- PUT, DELETE


__Result:__ Dummy object  

__Example:__

```sh
curl -XPOST -H "Content-Type:application/json" -d '
{
	"name"    : "default dummy", 
	"default" : true, 
	"assets"  : {
		"geometry" : "//localhost:6500/geometry/e748f388-36f8-47a2-b012-61f1083b80e7"
	}
}' http://v2.dressformer.com/api/dummies
```

Result:

```json
{
	"id"      : "0ae99696-0e13-4c54-8ad7-d1488dffbf65",
	"name"    : "default dummy",
	"default" : true,
	"assets": {
		"geometry": "//localhost:6500/geometry/e748f388-36f8-47a2-b012-61f1083b80e7"
	},
	"body": {}
}
```



## File API

### /
_Upload files to asset._

__Endpoint:__ `http://v2.dressformer.com/assets/`  
  
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


### /geometry/:id
__Endpoint:__ `http://v2.dressformer.com/assets/geometry/:id`  

__Methods:__

- GET, POST, PUT, DELETE

__Accept:__  

- `application/octet-stream` default for GET method
- `application/json` default for POST, PUT, DELETE mothods

__Result:__

- Content-Type: 
	- `application/octet-stream` 
	- `application/json`
- Default Format: Wavefront Obj

__Parameters:__

- `height`    (float)
- `chest`     (float)
- `underbust` (float)
- `waist`     (float)
- `hips`      (float)

Geometry object structure:

```javascript
{
	"id"            : "22a6dbed-1aab-452b-8f81-2a16e994120b",
	
	"base"          : "53f879c40000000000000002",
	
	"name"          : "Some optional name",
	
	"morph_targets" : [
		{
			"section" : "chest",
			"sources" : [
				{"id" : "53f879c40000000000000002", "weight" : 116.4},
				{"id" : "53f879c40000000000000003", "weight" : 130.0},
				{"id" : "53f879c40000000000000004", "weight" : 80.0}
			]
		}
}
```
Supported sections: `chest` `waist` `hips` `height` `underbust`  

__Example:__  
Get morphed geometry

```sh
GET v2.dressformer.com/assets/geometry/22a6dbed-1aab-452b-8f81-2a16e994120b?waist=95.0
```

### /image/:id
__Endpoint:__ `http://v2.dressformer.com/assets/image/:id`  
	
__Methods:__

- GET

__Result:__

- Content-Type: `image/jpeg` `image/png`
- Image, jpeg by default

__Parameters:__

- `scale` Scaling image to dimensions  
	Prototype: `([0-9]+x) or (x[0-9]+) or ([0-9]+) or (0.[0-9]+)`  
	E.g.:
  	+ `800x` scale to width 800px, height will be calculated
  	+ `x600` scale to height 600px, width will be calculated
  	+ `640`  maximum dimension is 640px, e.g. original 1024x768 pixel image will be scaled to 640x480,
           same option applied for 900x1600 image results 360x640
  	+ `0.5`  50% of original dimensions, e.g. 1024x768 = 512x384
- `q` 0-100 image quality
- `format` Image format — `png` or `jpg`. Jpeg is default one.

__Examples:__  

Get an image with 50% quality

```sh
curl -o test.jpg "http://v2.dressformer.com/assets/image/53f735c10000000000000001?q=50"
```

Get an image resized to 500x500px (original 2048x2048 becames 500x500) with 50% quality

```sh
curl -o test.jpg "http://v2.dressformer.com/assets/image/53f735c10000000000000001?q=50&scale=x500"
```

### More examples

#### Create mannequin

Upload all files to asset.

```sh
curl \
    -i -XPOST -H "ContentType:multipart/form-data" \
    -F name=447_base.obj -F filedata=@447_base.obj \
    -F name=447_chest_max.obj       -F filedata=@447_chest_max.obj \
    -F name=447_chest_min.obj       -F filedata=@447_chest_min.obj \
    -F name=447_height_max.obj      -F filedata=@447_height_max.obj \
    -F name=447_height_min.obj      -F filedata=@447_height_min.obj \
    -F name=447_hips_max.obj        -F filedata=@447_hips_max.obj \
    -F name=447_hips_min.obj        -F filedata=@447_hips_min.obj \
    -F name=447_underchest_max.obj  -F filedata=@447_underchest_max.obj \
    -F name=447_underchest_min.obj  -F filedata=@447_underchest_min.obj \
    -F name=447_waist_max.obj       -F filedata=@447_waist_max.obj \
    -F name=447_waist_min.obj       -F filedata=@447_waist_min.obj \
http://v2.dressformer.com/assets/
```

Result:

```json
[
    {
        "id": "53fcc20d0000000000000001",
        "orig_name": "447_base.obj"
    },
    {
        "id": "53fcc20d0000000000000002",
        "orig_name": "447_chest_max.obj"
    },
    
						...
]
```

Create geometry object for uploaded files.

```sh
curl -XPOST -H 'Content-Type:application/json' -d '
{
    "base" : "53fcc20d0000000000000001",
    "name" : "Base dummy",
    "morph_targets" : [
        {
            "section" : "height",
            "sources" : [
                {"id" : "53fcc20d0000000000000004", "weight" : 190.0},
                {"id" : "53fcc20d0000000000000005", "weight" : 155.0}
            ]
        },
        {
            "section" : "chest",
            "sources" : [
                {"id" : "53fcc20d0000000000000002", "weight" : 130.0},
                {"id" : "53fcc20d0000000000000003", "weight" : 80.0}
            ]
        },
        {
            "section" : "underbust",
            "sources" : [
                {"id" : "53fcc20d0000000000000008", "weight" : 130.0},
                {"id" : "53fcc20d0000000000000009", "weight" : 80.0}
            ]
        },
        {
            "section" : "waist",
            "sources" : [
                {"id" : "53fcc20d000000000000000a", "weight" : 90.0},
                {"id" : "53fcc20d000000000000000b", "weight" : 60.0}
            ]
        },
        {
            "section" : "hips",
            "sources" : [
                {"id" : "53fcc20d0000000000000006", "weight" : 110.0},
                {"id" : "53fcc20d0000000000000007", "weight" : 84.0}
            ]
        }
    ]
}' http://v2.dressformer.com/assets/geometry
```

Result:

```json
{
    "id": "e488c579-af46-45d3-8647-af5279dc1f86",
    "base": "53fcc20d0000000000000001",
    "morph_targets": [
        {
            "section": "height",
            "sources": [
                {
                    "id": "53fcc20d0000000000000004",
                    "weight": 190
                },
                {
                    "id": "53fcc20d0000000000000005",
                    "weight": 155
                }
            ]
        },
        
							...
							
```

Get created object

```sh
curl -XGET http://v2.dressformer.com/assets/geometry/e488c579-af46-45d3-8647-af5279dc1f86
```

Result:

```sh
HTTP/1.1 200 OK
Server: nginx
Date: Tue, 26 Aug 2014 19:14:34 GMT
Content-Type: application/octet-stream
Content-Length: 4095278

                             ... obj data ...
```