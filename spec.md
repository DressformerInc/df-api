### api.dressformer.com

#### `/v2/garments`
__Methods:__ 

- GET
- POST
- PUT
- DELETE

__Parameters:__

- `ids` coma-separated list of ids  

__Result:__ Array of objects

e.g.:

```
GET /v2/garments?ids=93e92e72-1bdb-436f-bdb9-52dba6c16176

[
	{},
]

```

#####  Garment Model

```javascript
[
	{
		// Public fields
		
		"id"    : "43a5dbed-4abb-452a-8f82-9a86e555930b", // Garment Id as a uuid
		
		"name"  : "Some name",
		
		"size"  : "M", // (String) Any size representation as a string, e.g. 
		               // "XL", "42", etc..., 

		"sizes" : [   // (Array of Objects)
			{"S" : "43a5dbed-4abb-452a-8f82-9a86e555930b"}, // name of size : garment id
			{"X" : "33a6dbed-4abb-452a-8f82-1a16e666930x"}  
		],

		"color" : "", // @todo
		
		"assets" : {
			"model" : "assets.dressformer.com/models/43a5dbed-4abb-452a-8f82-9a86e555930b" // Garment model, by default in obj format
			"diffuse" : "assets.dressformer.com/53b54559fcb05d3238000002" // Diffuse map
			"normal"  : "assets.dressformer.com/53b61050eff01c1008000001" // Normal map
		}
		
		// Private fields (available for admin-user)
		
		"sources" : { // Array of arrays of objects. Object is a morph-target - weight pair.
			[
				[
					{"53d11d10fcb05d8ed2000001"  : 83.0}, // id of morph-target : weight
					{"53d11d10fcb05d8ed2000002"  : 110.0}
					{"53d11d10fcb05d8ed2000003"  : 105.0}
				],
				[
					{"53d11d10fcb05d8ed2000011" : 190.0}, 
					{"53d11d10fcb05d8ed2000022" : 155.0}
				],
				[
					{"53d11d10fcb05d8ed2000031" : 90.0}, 
					{"53d11d10fcb05d8ed2000042" : 60.0}
				]
			]
		}
	},	
]
```

#### `/v2/user`
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
		"model" : "assets.dressformer.com/model/53d11d10fcb05d8ed2000042" // Some base mannequin
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


### assets.dressformer.com

#### `/model`
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

Gets mannequin morphed to waist 90.0

```sh
GET assets.dressformer.com/model/53d11d10fcb05d8ed2000042?waist=90.0
```

Gets garment with id `43a5dbed-4abb-452a-8f82-9a86e555930b` morphed to waist 90

```sh
GET assets.dressformer.com/model/43a5dbed-4abb-452a-8f82-9a86e555930b?waist=90.0
```

#### `/:id` 
	
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

	


