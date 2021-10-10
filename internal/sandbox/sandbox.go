package sandbox

import (
	"encoding/json"
	"fmt"
)

type myStruct struct {
	// Это поле структуры будет представлено в JSON объекте как "myName"
	Field1 int `json:"myNameField1"`
	// То же самое, что и в предыдущем примере, только если поле пустое, то его не будет в JSON объекте.
	Field2 int `json:"myNameField2,omitempty"`
	// Если не указать название поля, то будет использовано имя по умолчанию - "Field"
	Field3 int `json:",omitempty"`
	// Поле будет проигнорировано при сериализации этой структуры.
	Field4 int `json:"-"`
	// Поле будет представлено с ключом "-"
	Fiel5 int `json:"-,"`
	// Если указать опцию string, то поле будет представлено как UTF-8 строка JSON.
	Int64String int64 `json:",string"`
}

func Marshaling() {

	marshInst := myStruct{1, 2, 3, 4, 5, 6}

	js, _ := json.Marshal(marshInst)

	fmt.Printf("%v", string(js))

	var unmarshInst myStruct

	_ = json.Unmarshal(js, &unmarshInst)

}
