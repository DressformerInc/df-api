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
			// Garment model, by default in obj format
			"geometry" : "assets.dressformer.com/model/53b54559fcb05d3238000012" // Base model
			"diffuse" : "assets.dressformer.com/53b54559fcb05d3238000002" // Diffuse map
			"normal"  : "assets.dressformer.com/53b61050eff01c1008000001" // Normal map
		}
		
		// Private fields (available for admin-user)
		
		"sources" : { // Array of array of objects. Object is a morph-target - weight pair.
			[	
				[
					{ 
						"id" : "53d11d10fcb05d8ed2000001",
						"weight" : 83.0
					},
					{ 
						"id" : "53d11d10fcb05d8ed2000002",
						"weight" : 113.0
					},					
					
				],
				[
					{
						"id" : "53d11d10fcb05d8ed2000001",  
						"weight" : 83.0
					},
					{
						"id" : "53d11d10fcb05d8ed2000002",
						"weight"  : 110.0
					},
					{
						"id" : "53d11d10fcb05d8ed2000003",
						"weight"  : 105.0
					}
				]
			]
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
	// Not authorized, guest user

	"avatar" : {
		"model" : "assets.dressformer.com/geometry/53d11d10fcb05d8ed2000042" // Some base mannequin
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
Following `curl` command

```sh
curl -i -XPOST -H "ContentType:multipart/form-data"     \
	-F name=Base.obj      -F filedata=@Base.obj         \
	-F name=Chest_max.obj -F filedata=@Chest_max.obj    \
http://assets.dressformer.com/
```
returns

```json
[
	{
		"id"        : "53f622eb0000000000000001",
		"orig_name" : "Base.obj"
	},
	{
		"id"        : "53f622eb0000000000000002",
		"orig_name" : "Chest_max.obj"
	}
]	
```


### /geometry/:id
__Methods:__

- GET

__Result:__

- Content-Type: `application/octet-stream` 
- Default Format: Wavefront Obj

__Parameters:__

- `id` uuid or oid
- `height`    (float)
- `chest`     (float)
- `underbust` (float)
- `waist`     (float)
- `hips`      (float)

E.g.:

Gets geometry morphed to waist 90.0

```sh
GET assets.dressformer.com/geometry/53d11d10fcb05d8ed2000042?waist=90.0
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
- `quality` 0-100 image quality
- `format` Image format — `png` or `jpg`. Jpeg is default one.

	


