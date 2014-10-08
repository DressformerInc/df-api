#####  V1. Client side .mtl parsing


__Garment model__

```javascript
{

	"geometry" : {
		"url"  : "//v2.dressformer.com/geometry/22a6dbed-1aab-452b-8f81-2a16e994120b",
		"id"   : "22a6dbed-1aab-452b-8f81-2a16e994120b"
	},
	
	"materials" : [
		{	
			// writable field
			"id"        : "22a6dbed-1aab-452b-8f81-2a16e994121a",
			
			// following fields are auto generated
			"url"       : "//v2.dressformer.com/material/22a6dbed-1aab-452b-8f81-2a16e994121a"
			"name"      : "same name as on newmtl line",
		    "Ka"        : [1.000 1.000 1.000],
			"Kd"        : [1.000 1.000 1.000],
		    "Ks"        : [0.000 0.000 0.000],
		    
			"map_Kd"    : {
				"id"        : "53b54559fcb05d323800000a",
				"url"       : "//v2.dressformer.com/53b54559fcb05d323800000a",
				"orig_name" : "12345_diffuse.jpg"
			},
			
			"map_Ks"    : {
				"id"        : "53b5454fcb05d323800000b",
				"url"       : "//v2.dressformer.com/53b54559fcb05d323800000b",
				"orig_name" : "12345_specular.jpg"
			}
		},
	],
}
```

__Materials API__

Порядок работы: 

- заливаем в `assets/` все файлы, которые выбрал пользователь (карты нормалей, диффузы, итд)
- получаем айдишники
- парсим .mtl на клиенте и делаем маппинг известных нам полей + маппинг `orig_name` -> `id`
- делаем POST `/material` - отправляем массив объектов, в результате получаем массив id
- делаем POST или PUT с заполненным масивом айдишников `{"materials":[{"id":"$id"}]}` `/garments`