package main

import "fmt"

func main() {
	//soal 1
	// var k byte;
	// var cekK bool;
	// fmt.Scan(&k);

	// cekK = (k >= 0 && k <= 9) || (k >= 'a' && k <= 'z') || (k >= 'A' && k <= 'Z');
	// fmt.Println(cekK);


	//soal 2
	// var tahun int;
	// var cekTahun bool;
	// fmt.Scan(&tahun);

	// cekTahun = (tahun % 4 == 0 && tahun % 100 != 0) || (tahun % 400 == 0);
	// fmt.Print(cekTahun);


	//soal 3
	// var sudutA, sudutB, sudutC int;
	// var cekSegitiga bool;
	// fmt.Scanln(&sudutA, &sudutB, &sudutC);

	// cekSegitiga = sudutA+sudutB > sudutC && sudutA+sudutC > sudutB && sudutB+sudutC > sudutA;
	// fmt.Println(cekSegitiga);


	//soal 4
	// var total int
    // var inginKartu, dapatKartu, dapatDiskon, dapatCashback, Diskon, Cashback bool

    // fmt.Scan(&total)
    // fmt.Scan(&inginKartu)

    // Diskon = total >= 100000
    // Cashback = total >= 200000 && inginKartu

    // dapatKartu = inginKartu
    // dapatDiskon = Diskon
    // dapatCashback = Cashback

    // fmt.Println("Kartu?", dapatKartu)
    // fmt.Println("Diskon?", dapatDiskon)
    // fmt.Println("Cashback?", dapatCashback)


	//soal 5
	var a, b, c int;
	var cek bool;

	cek = true;

	cek = (a + b) / 2 == c || (a + c) / 2 == b || (b + c) / 2 == a;
	fmt.Println(cek);
}