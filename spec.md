## api.dressformer.com

### /garments
__Methods:__ 

- GET

__Parameters:__

- `ids` coma-separated list of ids  

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

### /garment/:id
__Methods:__ 

- GET
- POST
- PUT
- DELETE


#####  Garment Model

```javascript
[
	{
		// Public fields
		
		"id"    : "43a5dbed-4abb-452a-8f82-9a86e555930b", // Garment Id as a uuid
		
		"name"  : "Some name",
		
		"size_name"  : "M", // (String) Any size representation as a string, e.g. 
		               // "XL", "42", etc..., 

		"sizes" : [   // (Array of Objects)
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
			"geometry" : "assets.dressformer.com/geometry/22a6dbed-1aab-452b-8f81-2a16e994120b"
			
			"diffuse"  : "assets.dressformer.com/image/53b54559fcb05d3238000002"
			
			"normal"   : "assets.dressformer.com/image/53b61050eff01c1008000001"
		}		
	},	
]
```

### /user
__Methods:__ 

- GET
- POST
- PUT

__Parameters:__


__Result:__ Object

##### User Model

```javascript
{
	// Not authorized, guest user with default settings

	"avatar" : {
		"model" : "assets.dressformer.com/geometry/22a6dbed-1aab-452b-8f81-2a16e994120b" // Some base mannequin
	},
	
	"body" : {
		"height"    : 170.0,
		"chest"     : 90.0,
		"underbust" : 70.0,
		"waist"     : 60.0,
		"hips"      : 90.0
	}
}
```

## assets.dressformer.com

### /
_Upload files to asset._
  
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
http://assets.dressformer.com/
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

```json
	"id"            : "22a6dbed-1aab-452b-8f81-2a16e994120b",
	"base"          : "53f879c40000000000000002",
	"morph_targets" : [
		{
			"section" : "chest"
			"sources" : [
				{"id" : "53f879c40000000000000002", "weight" : 116.4},
				{"id" : "53f879c40000000000000003", "weight" : 130.0},
				{"id" : "53f879c40000000000000004", "weight" : 80.0}
			]
		}
```
Supported sections: `chest` `waist` `hips` `height` `underbust`  

__Example:__  
Get morphed geometry

```sh
GET assets.dressformer.com/geometry/22a6dbed-1aab-452b-8f81-2a16e994120b?waist=95.0
```

### /image/:id
	
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
curl -o test.jpg "http://assets.dressformer.com/image/53f735c10000000000000001?q=50"
```

Get an image resized to 500x500px (original 2048x2048 becames 500x500) with 50% quality

```sh
curl -o test.jpg "http://assets.dressformer.com/image/53f735c10000000000000001?q=50&scale=500"
```

### More examples

#### Create mannequin

Upload all files to asset.

```sh
curl                                                                                       \
	-i -XPOST -H "ContentType:multipart/form-data"                                         \
	-F name=base.obj             -F filedata=@base.obj                                     \
	-F name=chest_1164.obj       -F filedata=@chest_1164.obj                               \
	-F name=chest_1300.obj       -F filedata=@chest_1300.obj                               \
	-F name=chest_800.obj        -F filedata=@chest_800.obj                                \
	-F name=height_1550.obj      -F filedata=@height_1550.obj                              \
	-F name=height_1900.obj      -F filedata=@height_1900.obj                              \
	-F name=hips_1100.obj        -F filedata=@hips_1100.obj                                \
	-F name=hips_1248.obj        -F filedata=@hips_1248.obj                                \
	-F name=hips_840.obj         -F filedata=@hips_840.obj                                 \
	-F name=underchest_730.obj   -F filedata=@underchest_730.obj                           \
	-F name=underchest_780.obj   -F filedata=@underchest_780.obj                           \
	-F name=waist_586.obj        -F filedata=@waist_586.obj                                \
	-F name=waist_900.obj        -F filedata=@waist_900.obj                                \
http://assets.dressformer.com 
```

Result:

```json
[
    {
        "id": "53f879c40000000000000001",
        "orig_name": "base.obj"
    },
    {
        "id": "53f879c40000000000000002",
        "orig_name": "chest_1164.obj"
    },
    
				... cutted ...
]
```

Create geometry object for uploaded files.

```sh
curl -X POST -H 'Content-Type:application/json' -d '
{
	"base" : "53f879c40000000000000001",
	"morph_targets" : [
		{
			"section" : "chest",
			"sources" : [
				{"id" : "53f879c40000000000000002", "weight" : 116.4},
				{"id" : "53f879c40000000000000003", "weight" : 130.0},
				{"id" : "53f879c40000000000000004", "weight" : 80.0}
			]
		},
		{
			"section" : "height",
			"sources" : [
				{"id" : "53f879c40000000000000005", "weight" : 155.0},
				{"id" : "53f879c40000000000000006", "weight" : 190.0}
			]
		},
		{
			"section" : "hips",
			"sources" : [
				{"id" : "53f879c40000000000000007", "weight" : 110.0},
				{"id" : "53f879c40000000000000008", "weight" : 124.8},
				{"id" : "53f879c40000000000000009", "weight" : 84.0}
			]
		},
		{
			"section" : "underbust",
			"sources" : [
				{"id" : "53f879c4000000000000000a", "weight" : 73.0},
				{"id" : "53f879c4000000000000000b", "weight" : 78.0}
			]
		},
		{
			"section" : "waist",
			"sources" : [
				{"id" : "53f879c4000000000000000c", "weight" : 58.6},
				{"id" : "53f879c4000000000000000d", "weight" : 90.0}
			]
		}
	]
}' http://assets.dressformer.com/geometry
```

Result:

```json
{
    "id"   : "d537757e-9b95-42d0-8d44-b769b4ece0b4",  
    "base" : "53f879c40000000000000001",
    "morph_targets": [
        {
            "section" : "chest",
            "sources" : [
                {
                    "id"     : "53f879c40000000000000002",
                    "weight" : 116.4
                },
                {
                    "id"     : "53f879c40000000000000003",
                    "weight" : 130
                },
                {
                    "id"     : "53f879c40000000000000004",
                    "weight" : 80
                }
            ]
        },
        
                    ... cutted ...
```