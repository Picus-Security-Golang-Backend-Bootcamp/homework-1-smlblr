package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

func main() {
	args := os.Args

	// os.Args sizlere komut satırında yazılan değerleri bir string slice olarak verir
	// ilk argüman her zaman programın kendisi yani çalıştırılabilir hali oluyor.

	// Eğer argümanların uzunluğu 1 ise kullanıcı uygulamayı çalıştırıp hiçbir komut belirtmemiş demektir.
	if len(args) == 1 {
		//args listesinin 0. elemanı her zaman programın kendisidir.
		// path.Base parametre olarak verilen dosya yolundaki son değeri geriye döner.
		projectName := path.Base(args[0])
		fmt.Printf("%s uygulamasında kullanabileceğiniz komutlar : \n search => arama işlemi için \n list => listeleme işlemi için\n", projectName)
		return
	}

	// i değerini 1'den başlatma sebebimiz args slice'ındaki 0. elemanın bize programımızın adını/yolunu vermesi
	for i := 1; i < len(args); i++ {
		fmt.Printf("Kullanıcı tarafından girilen %d. değer : %s\n", i, args[i])
	}

	fileContent, err := ioutil.ReadFile("books.txt")
	if err != nil {
		fmt.Printf("failed reading data from file: %s\n", err)
		return
	}
	// fmt.Println(fileContent)

	bookList := strings.Split(string(fileContent), "\n")
	if len(args) == 2 && args[1] == "list" {
		fmt.Println("Kitaplar listeleniyor...")
		for i := range bookList {
			fmt.Printf("%d. %s\n", i+1, bookList[i])
		}
	}

	if len(args) == 2 && args[1] == "search" {
		fmt.Println("\"search\" argümanından sonra aramak istediğiniz şeyi yazınız.")
	}

	if len(args) > 2 && args[1] == "search" {
		searchItem := strings.Join(args[2:], " ")
		fmt.Printf("Aranan değer: \"%s\"\n", searchItem)
		searchItem = strings.ToLower(searchItem)
		for i := range bookList {
			if strings.Contains(strings.ToLower(bookList[i]), searchItem) {
				fmt.Printf("%d. %s\n", i+1, bookList[i])
			}
		}
	}
}
