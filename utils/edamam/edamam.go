package edamam

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type DetailNutrients struct {
	Enerc_kcal float64 `json:"ENERC_KCAL"`
	Procnt     float64 `json:"PROCNT"`
	Fat        float64 `json:"FAT"`
	Chocdf     float64 `json:"CHOCDF"`
	Fibtg      float64 `json:"FIBTG"`
}

type DetailFood struct {
	FoodId        string          `json:"foodId"`
	Label         string          `json:"label"`
	Nutrients     DetailNutrients `json:"nutrients"`
	Category      string          `json:"category"`
	CategoryLabel string          `json:"categoryLabel"`
	Image         string          `json:"image"`
}

type DetailMeasures struct {
	Uri    string `json:"uri"`
	Label  string `json:"label"`
	Weight int    `json:"weight"`
}

type Data struct {
	Food     DetailFood       `json:"food"`
	Measures []DetailMeasures `json:"measures"`
}

type Response struct {
	Text   string            `json:"text"`
	Parsed map[string]string `json:"parsed"`
	Hints  []Data            `json:"hints"`
}

func FoodThirdParty(s string) (Response, error) {

	// https://api.edamam.com/api/food-database/v2/parser?app_id=be2d6a07&app_key=28fd93ac7f43534e5a28ed8843adbfa7&ingr=a&nutrition-type=cooking
	url := fmt.Sprintf("https://api.edamam.com/api/food-database/v2/parser?app_id=be2d6a07&app_key=28fd93ac7f43534e5a28ed8843adbfa7&ingr=%s&nutrition-type=cooking", s)

	apiGet, err := http.Get(url)
	if err != nil {
		return Response{}, err
		// fmt.Print(err.Error())
		// os.Exit(1)
	}

	bodyData, err := ioutil.ReadAll(apiGet.Body)
	if err != nil {
		return Response{}, err
		// log.Fatal(err)
	}
	defer apiGet.Body.Close()

	response := Response{}
	// bodyString := bodyData
	// fmt.Println("API as a string:\n", response)

	json.Unmarshal(bodyData, &response)
	// fmt.Println("API as a string:\n", response)
	// for i := 0; i < len(response.Hints); i++ {
	// 	fmt.Print("============= ", response.Hints[i].Food.Nutrients)
	// }
	return response, nil

}
