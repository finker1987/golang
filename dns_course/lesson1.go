//Задание
//Нужно подготовить массив товаров полученных из апи сервиса хайповых товаров в том виде, в котором он отдается.
//Ссылка на апи http://hype-products-counter-master.hc-k8s.dns-shop.ru/api/v1/get-all
//Ключ авторизации: Bearer zXSgQ4(!9.b[zg}[

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Готовим структуру ответа. (data парсить не стал, для пример хватит)
type config struct {
	isSuccess bool   `json:"isSuccess"`
	errors    string `json:"errors"`
	data      string `json:"data"`
}

func main() {
	// Создаем http клиента
	client := http.Client{}

	// Готовим запрос
	req, err := http.NewRequest("GET", "http://sc-chat-js1.dns-shop.ru:8080/api/get-settings", nil)
	if err != nil {
		panic(err)
	}

	// Выполняем запрос
	r, err := client.Do(req)

	// Получаем тело ответо в байтах, тело можно вывести на экран через fmt преобразовав в string.
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	// Создаем переменную нашего типа и распаковываем в нее json ответа.
	var cfg config
	err = json.Unmarshal(body, &cfg)
	if err != nil {
		panic(err)
	}

	// Выводим на экран одно из полей.
	fmt.Println(cfg.isSuccess)

}
