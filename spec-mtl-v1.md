#####  V1. Client side .mtl parsing


__Material model__

```javascript
{	
	"id"        : "22a6dbed-1aab-452b-8f81-2a16e994121a",
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
		"orig_name" : "12345_specular.jpg",
		"options"   : "-o 1 1 1"
	}
}
```

__Examples__  

```sh
curl -XPOST -H "Content-Type:application/json" -d '
[{
    "name"  : "Test Material",
    "ka"    :   "0.0435 0.0435 0.0435",
    "kd"    :   "0.1086 0.1086 0.1086",
    "ks"    :   "0.0000 0.0000 0.0000",
    "tf"    :   "0.9885 0.9885 0.9885",
    "illum" :   6,
    "d"     :   "-halo 0.6600",
    "ns"    :   "10.0000",
    "ni"    :   "1.19713",

    "map_ka": {
        "id"        : "54390be70000000000000001",
        "options"   : "-s 1 1 1 -o 0 0 0 -mm 0 1",
        "orig_name" : "test_300.jpg"
    },

    "map_kd": {
        "id"        : "54390be70000000000000001",
        "options"   : "-s 1 1 1 -o 0 0 0 -mm 0 1",
        "orig_name" : "test_300.jpg"
    }
}]' http://localhost:5500/materials
```

Result:

```json
[
  "4717c27a-9fb0-4d09-ad70-2ae01f85deeb"
]
```

__Materials API__

Порядок работы: 

- заливаем в `assets/` все файлы, которые выбрал пользователь (карты нормалей, диффузы, итд)
- получаем айдишники
- парсим .mtl на клиенте и делаем маппинг известных нам полей + маппинг `orig_name` -> `id`
- делаем POST `/material` - отправляем массив объектов, в результате получаем массив id
- делаем POST или PUT с заполненным масивом айдишников `{"materials":[{"id":"$id"}]}` `/garments`