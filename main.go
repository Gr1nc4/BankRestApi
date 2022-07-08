package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type account struct {
	Client  client  `json:"Client"`
	Balance float64 `json:"Balance"`
	//	haveLoanLimit bool    `json:"HaveLoanLimit"`
}

type client struct {
	Id      int    `json:"Id"`
	Name    string `json:"Name"`
	Surname string `json:"Surname"`
}

var accStore = []account{}

var DataSt = map[int]account{}

func main() {
	//для теста
	// DataStore := make(map[int]account)

	// cl1 := addNewClient(1, "Evgeny", "Grishchuk")
	// acc1 := addNewAccount(cl1, 123.21, true)
	// balanceAdd(&acc1, 200.0)

	// cl2 := addNewClient(2, "Kirill", "Abramenko")
	// acc2 := addNewAccount(cl2, 50.0, false)

	// fmt.Println("Сумма акканутов до перевода", acc1.Balance, acc2.Balance)

	// transfer(&acc2, &acc1, 2000)

	// fmt.Println("Сумма аккаунтов после перевода", acc1.Balance, acc2.Balance)

	// addToDataStore(DataStore, &acc1)
	// addToDataStore(DataStore, &acc2)

	// fmt.Println(DataStore)

	mux := http.NewServeMux()

	mux.HandleFunc("/create_account", createAccount)
	mux.HandleFunc("/", getAccountById)
	mux.HandleFunc("/transaction", transaction)
	err := http.ListenAndServe(":8081", mux)
	log.Fatal(err)
}

func createAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		fmt.Println(" incorret method", http.StatusMethodNotAllowed)
		return
	}
	var acc account
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(w, " error with reading json", http.StatusNotAcceptable)
		return
	}
	err = json.Unmarshal(reqBody, &acc)
	if err != nil {
		fmt.Println("не удалось распарсить")
	}
	accStore = append(accStore, acc)
	DataSt[acc.Client.Id] = acc
	w.WriteHeader(http.StatusCreated)
	fmt.Println("создан авкаунт клиента", acc.Client.Name)
	fmt.Println(DataSt)
}

func getAccountById(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	AccId, _ := strconv.Atoi(id)

	fmt.Println(AccId, DataSt[AccId])
}
func transaction(w http.ResponseWriter, r *http.Request) {
	firstId := r.URL.Query().Get("id1")
	SecondId := r.URL.Query().Get("id2")
	S := r.URL.Query().Get("sum")

	id1, _ := strconv.Atoi(firstId)
	id2, _ := strconv.Atoi(SecondId)
	sum, _ := strconv.ParseFloat(S, 64)
	acc1 := DataSt[id1]
	acc2 := DataSt[id2]

	transfer(&acc1, &acc2, sum)
}

//Добавление аккаунта в БД
// func addToDataStore(data map[int]account, acc *account) map[int]account {
// 	data[acc.Client.Id] = *acc
// 	return data
// }

// //Создание аккаунта для клиента
// func addNewAccount(cl client, bal float64, limit bool) account {
// 	acc := account{
// 		Client:  cl,
// 		Balance: bal,
// 		//haveLoanLimit: limit,
// 	}

// 	return acc
// }

// //Создание клиента
// func addNewClient(id int, name string, surname string) client {
// 	cl := client{
// 		Id:      id,
// 		Name:    name,
// 		Surname: surname,
// 	}
// 	return cl
// }

//Пополнение баланса
func balanceAdd(acc *account, sum float64) {
	acc.Balance += sum
}

//Уменьшение баланса
func balanceDecrease(acc *account, sum float64) {
	acc.Balance -= sum
}

//Перевод средств с одного аккаунта на другой
func transfer(acc1, acc2 *account, sum float64) {

	// if !acc2.haveLoanLimit && acc2.balance < sum {
	// 	fmt.Println("Ошибка перевода, недостаточно средств у", acc2.client.name)
	// } else {
	balanceAdd(acc1, sum)
	balanceDecrease(acc2, sum)
	fmt.Println("транзкация успешна")

}

/*
{
    "Client": {
        "ID":"1",
        "Name":"Evgeny",
        "Surname":"Grishchuk"
    },
    "Balance":"100.21",
    "HaveLoanLimit":"false"

}
*/
