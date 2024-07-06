package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type penelitian struct {
	ketua, anggota1, anggota2, anggota3, fakultas, judul, dana, luaran string
	tahun                                                              int
}
type arrPenelitian [1000]penelitian
type rekap struct {
	tahun   int
	jumlah1 int
}
type arrRekap [100]rekap

var T arrPenelitian
var n int

func header() {
	fmt.Println("=======================================")
	fmt.Println("            SELAMAT DATANG         ")
	fmt.Println("=======================================")
	fmt.Println("  APLIKASI TRI DHARMA PERGURUAN TINGGI")
	fmt.Println("=======================================")
	fmt.Println("          Universitas Telkom       ")
	fmt.Println("=======================================")
}
func login(p, s *string) {
	fmt.Println("LOGIN")
	fmt.Print("Username", "\t", ": ")
	fmt.Scan(p)
	fmt.Print("Password", "\t", ": ")
	fmt.Scan(s)
}
func bersihLayar() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
func showMenu() {
	var pilih int
	for pilih != 5 {
		fmt.Println("=======================================")
		fmt.Println("   APLIKASI TRI DHARMA PERGURUAN TINGGI    ")
		fmt.Println("           TELKOM UNIVERSITY            ")
		fmt.Println("=======================================")
		fmt.Println("Pilihan Menu:")
		fmt.Println("1. Menampilkan Data")
		fmt.Println("2. Menambahkan Data")
		fmt.Println("3. Mengedit Data")
		fmt.Println("4. Menghapus Data")
		fmt.Println("5. Exit")
		fmt.Println("=======================================")
		fmt.Print("Tentukan pilihan Anda (1-5): ")
		fmt.Scan(&pilih)
		switch pilih {
		case 1:
			bersihLayar()
			MenuTampilData()
		case 2:
			bersihLayar()
			menambahkanData(&T, &n)
		case 3:
			bersihLayar()
			mengeditData(&T, n)
		case 4:
			bersihLayar()
			menghapusData(&T, n)
		case 5:
			bersihLayar()
			farewell()
			os.Exit(0)
		}
	}
}
func menambahkanData(T *arrPenelitian, n *int) {
	*n++
	var pilih, x int
	var s string
	for {
		bersihLayar()

		fmt.Println("=======================================")
		fmt.Println("            MENAMBAHKAN DATA   ")
		fmt.Println("=======================================")

		fmt.Print("Nama ketua : ")
		fmt.Scan(&s)
		T[*n-1].ketua = s
		fmt.Print("Nama anggota 1 : ")
		fmt.Scan(&s)
		T[*n-1].anggota1 = s
		fmt.Print("Nama anggota 2 : ")
		fmt.Scan(&s)
		T[*n-1].anggota2 = s
		fmt.Print("Nama anggota 3 : ")
		fmt.Scan(&s)
		T[*n-1].anggota3 = s
		fmt.Print("Nama fakultas : ")
		fmt.Scan(&s)
		T[*n-1].fakultas = s
		fmt.Print("Judul : ")
		fmt.Scan(&s)
		T[*n-1].judul = s
		fmt.Print("Sumber dana (internal/eksternal) : ")
		fmt.Scan(&s)
		T[*n-1].dana = s
		fmt.Print("Luaran : ")
		fmt.Scan(&s)
		T[*n-1].luaran = s
		fmt.Print("Tahun : ")
		fmt.Scan(&x)
		T[*n-1].tahun = x
		fmt.Println("Penambahan data berhasil")
		fmt.Println("<<0.Tambah data lagi>>", "\t", "<<1.Kembali ke Menu>>")
		fmt.Scan(&pilih)
		if pilih == 1 {
			break
		} else {
			*n++
		}
	}
	bersihLayar()
	showMenu()
}

func MenuTampilData() {
	var B arrPenelitian
	var pilih, x, n2 int
	var p string
	var R arrRekap
	var menu bool = false
	for !menu {
		fmt.Println("===============================================")
		fmt.Println("        MENU PILIHAN MENAMPILKAN DATA   ")
		fmt.Println("===============================================")
		fmt.Println("Pilihan Menu:")
		fmt.Println("1. Menampilkan Seluruh Data")
		fmt.Println("2. Menampilkan Data Berdasarkan Tahun")
		fmt.Println("3. Menampilkan Data Berdasarkan Fakultas")
		fmt.Println("4. Menampilkan Data Berdasarkan Nama Ketua Peneliti")
		fmt.Println("5. Menampilkan Data Berdasarkan Luaran")
		fmt.Println("6. Menampilkan Data Berdasarkan Dana (Internal/Eksternal)")
		fmt.Println("7. Menampilkan Data Berdasarkan Judul Penelitian")
		fmt.Println("8. Rekapitulasi Kegiatan Per Tahun")
		fmt.Println("9. Kembali ke Menu")
		fmt.Println("===============================================")
		fmt.Print("Masukkan pilihan Anda (1-10): ")
		fmt.Scan(&pilih)
		switch pilih {
		case 1:
			bersihLayar()
			fmt.Println("Menampilkan data terurut secara :")
			fmt.Println("1. Descending")
			fmt.Println("2. Ascending")
			fmt.Scan(&pilih)
			for pilih != 1 && pilih != 2 {
				fmt.Println("Angka yang Anda masukkan tidak valid, silakan coba lagi ")
				fmt.Scan(&pilih)
			}
			if pilih == 1 {
				bersihLayar()
				totalSortDesc(&T, n)
				tampilData(T, n)
			} else if pilih == 2 {
				bersihLayar()
				totalSortAsc(&T, n)
				tampilData(T, n)
			}
			fmt.Println("<<0. Kembali ke Menu>>")
			fmt.Scan(&pilih)
			for pilih != 0 {
				fmt.Println("Angka yang Anda masukkan tidak valid")
				fmt.Scan(&pilih)
			}
			bersihLayar()
			if pilih == 0 {
				return
			}

		case 2:
			bersihLayar()
			cariTahun(T, &B, n, x)
			fmt.Println("<<0. Kembali ke Menu>>")
			fmt.Scan(&pilih)
			for pilih != 0 {
				fmt.Println("Angka yang Anda masukkan tidak valid")
				fmt.Scan(&pilih)
			}
			bersihLayar()
			if pilih == 0 {
				return
			}
		case 3:
			bersihLayar()
			fmt.Println("Menampilkan data terurut secara :")
			fmt.Println("1. Descending")
			fmt.Println("2. Ascending")
			fmt.Scan(&pilih)
			for pilih != 1 && pilih != 2 {
				fmt.Println("Angka yang Anda masukkan tidak valid, silakan coba lagi ")
				fmt.Scan(&pilih)
			}
			if pilih == 1 {
				bersihLayar()
				cariArray(T, &B, n, p, 3, pilih)
			} else if pilih == 2 {
				bersihLayar()
				cariArray(T, &B, n, p, 3, pilih)
			}
			fmt.Println("<<0. Kembali ke Menu>>")
			fmt.Scan(&pilih)
			for pilih != 0 {
				fmt.Println("Angka yang Anda masukkan tidak valid")
				fmt.Scan(&pilih)
			}
			bersihLayar()
			if pilih == 0 {
				return
			}
		case 4:
			bersihLayar()
			fmt.Println("Menampilkan data terurut secara :")
			fmt.Println("1. Descending")
			fmt.Println("2. Ascending")
			fmt.Scan(&pilih)
			for pilih != 1 && pilih != 2 {
				fmt.Println("Angka yang Anda masukkan tidak valid, silakan coba lagi ")
				fmt.Scan(&pilih)
			}
			if pilih == 1 {
				bersihLayar()
				cariArray(T, &B, n, p, 4, pilih)
			} else if pilih == 2 {
				bersihLayar()
				cariArray(T, &B, n, p, 4, pilih)
			}
			fmt.Println("<<0. Kembali ke Menu>>")
			fmt.Scan(&pilih)
			for pilih != 0 {
				fmt.Println("Angka yang Anda masukkan tidak valid")
				fmt.Scan(&pilih)
			}
			bersihLayar()
			if pilih == 0 {
				return
			}
		case 5:
			bersihLayar()
			fmt.Println("Menampilkan data terurut secara :")
			fmt.Println("1. Descending")
			fmt.Println("2. Ascending")
			fmt.Scan(&pilih)
			for pilih != 1 && pilih != 2 {
				fmt.Println("Angka yang Anda masukkan tidak valid, silakan coba lagi ")
				fmt.Scan(&pilih)
			}
			if pilih == 1 {
				bersihLayar()
				cariArray(T, &B, n, p, 5, pilih)
			} else if pilih == 2 {
				bersihLayar()
				cariArray(T, &B, n, p, 5, pilih)
			}
			fmt.Scan(&pilih)
			for pilih != 0 {
				fmt.Println("Angka yang Anda masukkan tidak valid")
				fmt.Scan(&pilih)
			}
			bersihLayar()
			if pilih == 0 {
				return
			}
		case 6:
			bersihLayar()
			fmt.Println("Menampilkan data terurut secara :")
			fmt.Println("1. Descending")
			fmt.Println("2. Ascending")
			fmt.Scan(&pilih)
			for pilih != 1 && pilih != 2 {
				fmt.Println("Angka yang Anda masukkan tidak valid, silakan coba lagi ")
				fmt.Scan(&pilih)
			}
			if pilih == 1 {
				bersihLayar()
				cariArray(T, &B, n, p, 6, pilih)
			} else if pilih == 2 {
				bersihLayar()
				cariArray(T, &B, n, p, 6, pilih)
			}
			fmt.Println("<<0. Kembali ke Menu>>")
			fmt.Scan(&pilih)
			for pilih != 0 {
				fmt.Println("Angka yang Anda masukkan tidak valid")
				fmt.Scan(&pilih)
			}
			bersihLayar()
			if pilih == 0 {
				return
			}
		case 7:
			bersihLayar()

			binarySearch(&T, n)
			fmt.Print("<<0.Kembali ke Menu>>", "\n")
			fmt.Scan(&pilih)
			for pilih != 0 && pilih != 1 {
				fmt.Println("Angka yang Anda masukkan tidak valid")
				fmt.Scan(&pilih)
			}
			fmt.Print("<<0.Kembali ke Menu>>", "\n")
			bersihLayar()
			if pilih == 0 {
				return
			}
			bersihLayar()

		case 8:
			bersihLayar()
			fmt.Println("Menampilkan data terurut secara :")
			fmt.Println("1. Descending")
			fmt.Println("2. Ascending")
			fmt.Scan(&pilih)
			for pilih != 1 && pilih != 2 {
				fmt.Println("Angka yang Anda masukkan tidak valid, silakan coba lagi ")
				fmt.Scan(&pilih)
			}
			Rekapitulasi(T, &R, &n2, n)
			SortRekap(&R, n2, pilih)
			printRekap(R, n2)
			fmt.Print("<<0.Kembali ke Menu>>", "\n")
			fmt.Scan(&pilih)
			for pilih != 0 {
				fmt.Println("Angka yang Anda masukkan tidak valid")
				fmt.Scan(&pilih)
			}
			bersihLayar()

			if pilih == 0 {
				return
			}
		case 9:
			bersihLayar()
			menu = true
		}
		if pilih == 9 {
			break
		}

	}
	showMenu()
}
func tampilData(T arrPenelitian, n int) {
	fmt.Println(strings.Repeat("=", 180))
	fmt.Printf("%-20s%-15s%-15s%-15s%-15s%-70s%-10s%-15s%-15s\n", "Ketua Peneliti", "Anggota1", "Anggota2", "Anggota3", "Fakultas", "Judul", "Dana", "Luaran", "Tahun")
	fmt.Println(strings.Repeat("=", 180))
	for i := 0; i <= n; i++ {
		if T[i].tahun != 0 {
			fmt.Printf("%-20s%-15s%-15s%-15s%-15s%-70s%-10s%-15s%-15d\n", T[i].ketua, T[i].anggota1, T[i].anggota2, T[i].anggota3, T[i].fakultas, T[i].judul, T[i].dana, T[i].luaran, T[i].tahun)
		}
	}

	fmt.Println(strings.Repeat("=", 180))
}

func main() {
	var p, s string
	var berhenti bool = false
	assignArray(&T, &n)
	x := 0
	header()
	login(&p, &s)
	berhenti = p == "admin" && s == "123456"
	if berhenti {
		bersihLayar()
		showMenu()
	}
	for x < 3 && !berhenti {
		x++
		bersihLayar()
		header()
		fmt.Println("Username/Password yang Anda masukkan salah, silakan coba lagi.")
		login(&p, &s)
		berhenti = p == "admin" && s == "123456"
		if berhenti {
			bersihLayar()
			showMenu()
		}
		if x == 2 {
			fmt.Println("Terlalu banyak gagal mencoba")
			break
		}
	}

}

func assignArray(T *arrPenelitian, n *int) {
	T[0].ketua, T[0].anggota1, T[0].anggota2, T[0].anggota3, T[0].fakultas, T[0].judul, T[0].dana, T[0].luaran, T[0].tahun = "Prasti_Eko", "Ando", "Budi", "Candra", "FIF", "METODE_PENDETEKSIAN_DYSLEXIA_PADA_ORANG_DEWASA", "Internal", "Publikasi", 2021
	T[1].ketua, T[1].anggota1, T[1].anggota2, T[1].anggota3, T[1].fakultas, T[1].judul, T[1].dana, T[1].luaran, T[1].tahun = "Prasti_Eko", "Ceko", "Doding", "Catur", "FIF", "DESAIN_EKSTRAKSI_CIRI_PADA_SISTEM_OTENTIKASI_PENGGUNA", "Internal", "Publikasi", 2021
	T[2].ketua, T[2].anggota1, T[2].anggota2, T[2].anggota3, T[2].fakultas, T[2].judul, T[2].dana, T[2].luaran, T[2].tahun = "Prasti_Eko", "Eko", "Fahmi", "Galang", "FIF", "MULTI-VOTERS_MULTI-COMMISSION_NEAREST_NEIGHBORS", "Internal", "Produk", 2021
	T[3].ketua, T[3].anggota1, T[3].anggota2, T[3].anggota3, T[3].fakultas, T[3].judul, T[3].dana, T[3].luaran, T[3].tahun = "Aji_Gautama", "Susi", "Santi", "Sandra", "FIF", "PENGEMBANGAN_IOT-BASED_CLINICAL_PATHWAY", "Eksternal", "Produk", 2017
	T[4].ketua, T[4].anggota1, T[4].anggota2, T[4].anggota3, T[4].fakultas, T[4].judul, T[4].dana, T[4].luaran, T[4].tahun = "Aji_Gautama", "Budi", "Budiman", "Astro", "FIF", "DESAIN_DAN_IMPLEMENTASI_MINI_SERVER_TOPUP_PULSA", "Eksternal", "Publikasi", 2018
	T[5].ketua, T[5].anggota1, T[5].anggota2, T[5].anggota3, T[5].fakultas, T[5].judul, T[5].dana, T[5].luaran, T[5].tahun = "Aji_Gautama", "Zulfahmi", "Zulfiana", "Zulfikar", "FIF", "SISTEM_PEMANTAUAN_PELAJAR_PENGGUNA_BUS_SEKOLAH", "Internal", "Produk", 2018
	T[6].ketua, T[6].anggota1, T[6].anggota2, T[6].anggota3, T[6].fakultas, T[6].judul, T[6].dana, T[6].luaran, T[6].tahun = "Aji_Gautama", "Dadang", "Wawan", "Maman", "FIF", "MULTI-APPLICATION_ENERGY_HARVESTING_PADA_IOT_AQUAPONIC", "Eksternal", "Produk", 2021
	T[7].ketua, T[7].anggota1, T[7].anggota2, T[7].anggota3, T[7].fakultas, T[7].judul, T[7].dana, T[7].luaran, T[7].tahun = "Dodi_Wisaksono", "Parman", "Yanto", "Sutris", "FIF", "PENGEMBANGAN_APLIKASI_MOBILE_UNTUK_UJI_KOMPETENSI", "Eksternal", "Produk", 2021
	T[8].ketua, T[8].anggota1, T[8].anggota2, T[8].anggota3, T[8].fakultas, T[8].judul, T[8].dana, T[8].luaran, T[8].tahun = "Dodi_Wisaksono", "Bambang", "Made", "Bagus", "FIF", "DESAIN_DAN_IMPLEMENTASI_PURWARUPA_ADVERTISED", "Internal", "Publikasi", 2020
	T[9].ketua, T[9].anggota1, T[9].anggota2, T[9].anggota3, T[9].fakultas, T[9].judul, T[9].dana, T[9].luaran, T[9].tahun = "Dodi_Wisaksono", "Deni", "Beni", "Reni", "FIF", "PERANCANGAN_DAN_IMPLEMENTASI_PENGENALAN_HAND_GESTURE", "Internal", "Publikasi", 2019
	T[10].ketua, T[10].anggota1, T[10].anggota2, T[10].anggota3, T[10].fakultas, T[10].judul, T[10].dana, T[10].luaran, T[10].tahun = "Dodi_Wisaksono", "Dono", "Jono", "Mulya", "FIF", "PERBANDINGAN_SISTEM_OTOMASI_KUNCI_PINTU_PENGENALAN_WICARA", "Internal", "Pbulikasi", 2018
	T[11].ketua, T[11].anggota1, T[11].anggota2, T[11].anggota3, T[11].fakultas, T[11].judul, T[11].dana, T[11].luaran, T[11].tahun = "Didit_Adiytia", "Endang", "Subaru", "Galih", "FIF", "PENDEKATAN_INVERSI_TSUNAMI_DENGAN_SIMULASI_NUMERIK", "Eksternal", "Publikasi", 2021
	T[12].ketua, T[12].anggota1, T[12].anggota2, T[12].anggota3, T[12].fakultas, T[12].judul, T[12].dana, T[12].luaran, T[12].tahun = "Didit_Adiytia", "Jaya", "Abdi", "Sastra", "FIF", "MODEL_PREDIKSI_GELOMBANG_LAUT_DENGAN_PENDEKATAN_ALGORITMA_ML", "Eksternal", "Publikasi", 2020
	T[13].ketua, T[13].anggota1, T[13].anggota2, T[13].anggota3, T[13].fakultas, T[13].judul, T[13].dana, T[13].luaran, T[13].tahun = "Didit_Adiytia", "Dika", "Wahyu", "Habibie", "FIF", "PENGEMBANGAN_SISTEM_DENOISING_SINYAL_ELECTROCARDIOGRAM", "Internal", "Publikasi", 2019
	T[14].ketua, T[14].anggota1, T[14].anggota2, T[14].anggota3, T[14].fakultas, T[14].judul, T[14].dana, T[14].luaran, T[14].tahun = "Didit_Adiytia", "Sandro", "Mirza", "Dana", "FIF", "WAVE_FORECASTING_SYSTEM", "Internal", "Produk", 2021
	T[15].ketua, T[15].anggota1, T[15].anggota2, T[15].anggota3, T[15].fakultas, T[15].judul, T[15].dana, T[15].luaran, T[15].tahun = "Suyanto", "Andi", "Bagas", "Cindy", "FIF", "PELATIHAN_TOOLS_UNTUK_PEMBELAJARAN_ONLINE_DAN_PEMBUATAN_MATERI", "Internal", "Pelatihan", 2021
	T[16].ketua, T[16].anggota1, T[16].anggota2, T[16].anggota3, T[16].fakultas, T[16].judul, T[16].dana, T[16].luaran, T[16].tahun = "Suyanto", "Dandi", "Sungganu", "Deka", "FIF", "WEBINAR_PENGENALAN_AI_DI_INDONESIA", "Internal", "Seminar", 2021
	T[17].ketua, T[17].anggota1, T[17].anggota2, T[17].anggota3, T[17].fakultas, T[17].judul, T[17].dana, T[17].luaran, T[17].tahun = "Suyanto", "Sigit", "Hadi", "Yudo", "FIF", "SISTEM_PENGENALAN_UCAPAN_UNTUK_RESEP_ELEKTRONIK", "Eksternal", "Produk", 2017
	T[18].ketua, T[18].anggota1, T[18].anggota2, T[18].anggota3, T[18].fakultas, T[18].judul, T[18].dana, T[18].luaran, T[18].tahun = "Suyanto", "Galang", "Gilang", "Gading", "FIF", "PENGEMBANGAN_IOT-BASED_CLINICAL_PATHWAY ", "Eksternal", "Produk", 2017
	T[19].ketua, T[19].anggota1, T[19].anggota2, T[19].anggota3, T[19].fakultas, T[19].judul, T[19].dana, T[19].luaran, T[19].tahun = "Achmad Rizal", "Bayu", "Sugondo", "Teno", "FTE", "EMPIRICAL_MODE_DECOMPOTISION_AND_HILBERT_SPECTRUM ", "Eksternal", "Publikasi", 2022
	T[20].ketua, T[20].anggota1, T[20].anggota2, T[20].anggota3, T[20].fakultas, T[20].judul, T[20].dana, T[20].luaran, T[20].tahun = "Casi_Setianingsih", "Ariana", "Ardianto", "Ogin", "FTE", "A_SEMI-SUPERVISED_ULTI-MODEL_PREDICTION_TECHNIQUE ", "Internal", "Publikasi", 2023
	T[21].ketua, T[21].anggota1, T[21].anggota2, T[21].anggota3, T[21].fakultas, T[21].judul, T[21].dana, T[21].luaran, T[21].tahun = "Casi_Setianingsih", "Werfyan", "Ardianto", "Diah", "FTE", "BIOSENSORS_FOR_THERAPEUTIC_DRUG_MONITORING", "Eksternal", "Publikasi", 2023
	T[22].ketua, T[22].anggota1, T[22].anggota2, T[22].anggota3, T[22].fakultas, T[22].judul, T[22].dana, T[22].luaran, T[22].tahun = "Fityanul_Akhyar", "Ying", "Chao", "Timothy", "FTE", "FDD:A_DEEP_LEARNING-BASED_STEEL_DEFECT_DETECTORS", "Internal", "Publikasi", 2023
	T[23].ketua, T[23].anggota1, T[23].anggota2, T[23].anggota3, T[23].fakultas, T[23].judul, T[23].dana, T[23].luaran, T[23].tahun = "Husneni_Mukhtar", "Manuel", "Fredy", "Anto", "FTE", "CHARACTERIZATION_OF_FUNCTIONAL_MATERIALS", "Eksternal", "Publikasi", 2023
	T[24].ketua, T[24].anggota1, T[24].anggota2, T[24].anggota3, T[24].fakultas, T[24].judul, T[24].dana, T[24].luaran, T[24].tahun = "Doan_Perdana", "Aji", "Maman", "Hilal", "FTE", "CIMA:A_NOVEL_CLASSIFICATION_MOVING_AVERAGE_MODEL", "Internal", "Publikasi", 2022
	T[25].ketua, T[25].anggota1, T[25].anggota2, T[25].anggota3, T[25].fakultas, T[25].judul, T[25].dana, T[25].luaran, T[25].tahun = "Prayitno_Abadi", "Slamet", "Hasanudin", "Dudy", "FTE", "CONSTRUCTION_OF_NOMINAL_IONOSPHERIC_GRADIENT", "Eksternal", "Publikasi", 2022
	T[26].ketua, T[26].anggota1, T[26].anggota2, T[26].anggota3, T[26].fakultas, T[26].judul, T[26].dana, T[26].luaran, T[26].tahun = "Nasrullah_Armi", "Omar", "Wajeb", "Andi", "FTE", "AN_E-COMMERCE_CONTROL_UNIT_FOR_ADDRESSING_ONLINE_TRANSACTIONS", "Internal", "Publikasi", 2022
	T[27].ketua, T[27].anggota1, T[27].anggota2, T[27].anggota3, T[27].fakultas, T[27].judul, T[27].dana, T[27].luaran, T[27].tahun = "Rustam", "Agus", "Made", "Tri", "FTE", "DATA_DIMENSIONALITY_REDUCTION_TECHNIQUE_FOR_CLUSTERING_PROBLEM", "Internal", "Publikasi", 2021
	T[28].ketua, T[28].anggota1, T[28].anggota2, T[28].anggota3, T[28].fakultas, T[28].judul, T[28].dana, T[28].luaran, T[28].tahun = "Achmad_Mansur", "Runik", "Cahyo", "Dwi", "FEB", "SISTEM_MONITORING_PENANAMAN,PEMUPUKAN,PENYIRAMAN,PEMANENAN", "Internal", "Publikasi", 2022
	T[29].ketua, T[29].anggota1, T[29].anggota2, T[29].anggota3, T[29].fakultas, T[29].judul, T[29].dana, T[29].luaran, T[29].tahun = "Achmad_Mansur", "Yudha", "Runik", "Eka", "FEB", "ANALISIS_IMPLEMENTASI_SYSTEM_INVENTORY_MANAGEMENT_DI_UMKM_JAKET_KEREN", "Internal", "Publikasi", 2021
	T[30].ketua, T[30].anggota1, T[30].anggota2, T[30].anggota3, T[30].fakultas, T[30].judul, T[30].dana, T[30].luaran, T[30].tahun = "Teguh_Widodo", "Achmad", "Soeparwoto", "Vany", "FEB", "TRANSFORMATIVE_ORGANIZATIONAL_LEARNING_MODEL_AND_ITS_IMPACT", "Eksternal", "Publikasi", 2020
	T[31].ketua, T[31].anggota1, T[31].anggota2, T[31].anggota3, T[31].fakultas, T[31].judul, T[31].dana, T[31].luaran, T[31].tahun = "Achmad_Mansur", "Yudha", "Yahya", "Enda", "FEB", "MODEL_STRATEGI_MARKETING_KOMUNITAS_FASHION_BERBAHAN_KULIT_DI_GARUT", "Internal", "Publikasi", 2020
	T[32].ketua, T[32].anggota1, T[32].anggota2, T[32].anggota3, T[32].fakultas, T[32].judul, T[32].dana, T[32].luaran, T[32].tahun = "Achmad_Mansur", "Runik", "Yudha", "Pramoto", "FEB", "MODEL_PEMASARAN_UMK_KOMUNITAS_FASHION_KULIT_DI_ERA_DIGITAL_MARKETING", "Internal", "Publikasi", 2020
	T[33].ketua, T[33].anggota1, T[33].anggota2, T[33].anggota3, T[33].fakultas, T[33].judul, T[33].dana, T[33].luaran, T[33].tahun = "Teguh_Widodo", "Achmad", "Soeparwoto", "Vany", "FEB", "STUDY_ON_BUILDING_STRATEGIC_INTELLECTUAL_CAPITAL_IN_IT-BASED_COMPANY", "Internal", "Publikasi", 2020
	T[34].ketua, T[34].anggota1, T[34].anggota2, T[34].anggota3, T[34].fakultas, T[34].judul, T[34].dana, T[34].luaran, T[34].tahun = "Muhammad_Yahya", "Achmad", "Runik", "Machfiroh", "FEB", "ANALISA_PERUBAHAN_SOSIAL_BUDAYA_DALAM_PEMBENTUKKAN_DESA_WISATA", "Internal", "Publikasi", 2019
	T[35].ketua, T[35].anggota1, T[35].anggota2, T[35].anggota3, T[35].fakultas, T[35].judul, T[35].dana, T[35].luaran, T[35].tahun = "Adhi_Prasetio", "Ferinia", "Tanjung", "Purba", "FEB", "PERILAKU_KONSUMEN_KEPARIWISTAAN", "Internal", "Publikasi", 2021
	T[36].ketua, T[36].anggota1, T[36].anggota2, T[36].anggota3, T[36].fakultas, T[36].judul, T[36].dana, T[36].luaran, T[36].tahun = "Adhi_Prasetio", "Hurriyati", "Sari", "Dani", "FEB", "SOCIAL_CAPITAL_AND_ELECTRONICS_WORD-OF-MOUTH_(eWOM)", "Internal", "Publikasi", 2017
	T[37].ketua, T[37].anggota1, T[37].anggota2, T[37].anggota3, T[37].fakultas, T[37].judul, T[37].dana, T[37].luaran, T[37].tahun = "Adhi_Prasetio", "Sharif", "Perdana", "Alamanda", "FEB", "ANALYZING_THE_IMPACT_OF_TRAFFIC_SOURCE_ON_VISIT_DURATION", "Internal", "Publikasi", 2017
	T[38].ketua, T[38].anggota1, T[38].anggota2, T[38].anggota3, T[38].fakultas, T[38].judul, T[38].dana, T[38].luaran, T[38].tahun = "Adhi_Prasetio", "Islami", "Angga", "Putra", "FEB", "ANALISIS_DAN_PREDIKSI_PENJUALAN_PADA_MARKETPLACE", "Eksternal", "Publikasi", 2018

	*n = 39
}

func fakultasSort(T *arrPenelitian, n int) {
	pass := 1
	for pass < n {
		i := pass
		temp := T[pass]
		for i > 0 && temp.fakultas < T[i-1].fakultas {
			T[i] = T[i-1]
			i--
		}
		T[i] = temp
		pass++
	}
}
func tahunSortDesc(T *arrPenelitian, n int) {
	pass := 1
	for pass <= n-1 {
		idx := pass - 1
		i := pass
		for i < n {
			if T[idx].tahun < T[i].tahun && T[idx].fakultas == T[i].fakultas {
				idx = i
			}
			i++
		}
		T[pass-1], T[idx] = T[idx], T[pass-1]
		pass++
	}

}
func tahunSortAsc(T *arrPenelitian, n int) {
	pass := 1
	for pass <= n-1 {
		idx := pass - 1
		i := pass
		for i < n {
			if T[idx].tahun > T[i].tahun && T[idx].fakultas == T[i].fakultas {
				idx = i
			}
			i++
		}
		T[pass-1], T[idx] = T[idx], T[pass-1]
		pass++
	}

}

func namaSort(T *arrPenelitian, n int) {
	pass := 1
	for pass <= n-1 {
		idx := pass - 1
		i := pass
		for i < n {
			if T[idx].ketua > T[i].ketua && T[idx].fakultas == T[i].fakultas && T[i].tahun == T[idx].tahun {
				idx = i
			}
			i++
		}
		T[pass-1], T[idx] = T[idx], T[pass-1]
		pass++
	}

}
func namaSortBinary(T *arrPenelitian, n int) {
	pass := 1
	for pass <= n-1 {
		idx := pass - 1
		i := pass
		for i < n {
			if T[idx].judul > T[i].judul {
				idx = i
			}
			i++
		}
		T[pass-1], T[idx] = T[idx], T[pass-1]
		pass++
	}

}

func totalSortDesc(T *arrPenelitian, n int) {
	fakultasSort(T, n)
	tahunSortDesc(T, n)
	namaSort(T, n)
}

func totalSortAsc(T *arrPenelitian, n int) {
	fakultasSort(T, n)
	tahunSortAsc(T, n)
	namaSort(T, n)
}
func cariTahun(T arrPenelitian, B *arrPenelitian, n int, x int) {
	var hasil bool = false
	fmt.Println("Masukkan tahun yang ingin ditampilkan :")
	fmt.Scan(&x)
	j := 0
	for i := 0; i <= n; i++ {
		if T[i].tahun == x {
			B[j] = T[i]
			j++
			hasil = true
		}
	}
	if !hasil {
		fmt.Println("Maaf, data yang Anda cari tidak ditemukan.")
	} else {
		fakultasSort(B, j)
		namaSort(B, j)
		tampilData(*B, j)
	}
}

func cariArray(T arrPenelitian, B *arrPenelitian, n int, x string, y int, m int) {
	var hasil bool = false
	j := 0
	if y == 3 {
		fmt.Println("Masukkan fakultas yang ingin ditampilkan :")
		fmt.Scan(&x)
		for i := 0; i < n; i++ {
			if T[i].fakultas == x {
				B[j] = T[i]
				j++
				hasil = true
			}
		}
		if !hasil {
			fmt.Println("Maaf, data yang Anda cari tidak ditemukan.")
		} else {
			if m == 1 {
				tahunSortDesc(B, j)
				namaSort(B, j)
				tampilData(*B, j)
			} else {
				tahunSortAsc(B, j)
				namaSort(B, j)
				tampilData(*B, j)
			}

		}
	} else if y == 4 {
		fmt.Println("Masukkan nama ketua peneliti yang ingin ditampilkan :")
		fmt.Scan(&x)
		for i := 0; i < n; i++ {
			if T[i].ketua == x {
				B[j] = T[i]
				j++
				hasil = true
			}
		}
		if !hasil {
			fmt.Println("Maaf, data yang Anda cari tidak ditemukan.")
		} else {
			if m == 1 {
				fakultasSort(B, j)
				tahunSortDesc(B, j)
				tampilData(*B, j)
			} else {
				fakultasSort(B, j)
				tahunSortAsc(B, j)
				tampilData(*B, j)
			}

		}
	} else if y == 5 {
		fmt.Println("Masukkan jenis luaran yang ingin ditampilkan :")
		fmt.Scan(&x)
		for i := 0; i < n; i++ {
			if T[i].luaran == x {
				B[j] = T[i]
				j++
				hasil = true
			}
		}
		if !hasil {
			fmt.Println("Maaf, data yang Anda cari tidak ditemukan.")
		} else {
			if m == 1 {
				fakultasSort(B, j)
				tahunSortDesc(B, j)
				namaSort(B, j)
				tampilData(*B, j)
			} else {
				fakultasSort(B, j)
				tahunSortAsc(B, j)
				namaSort(B, j)
				tampilData(*B, j)
			}
		}
	} else if y == 6 {
		fmt.Println("Masukkan jenis sumber dana yang ingin ditampilkan :")
		fmt.Scan(&x)
		for i := 0; i < n; i++ {
			if T[i].dana == x {
				B[j] = T[i]
				j++
				hasil = true
			}
		}
		if !hasil {
			fmt.Println("Maaf, data yang Anda cari tidak ditemukan.")
		} else {
			if m == 1 {
				fakultasSort(B, j)
				tahunSortDesc(B, j)
				namaSort(B, j)
				tampilData(*B, j)
			} else {
				fakultasSort(B, j)
				tahunSortAsc(B, j)
				namaSort(B, j)
				tampilData(*B, j)
			}
		}
	}
}
func cariRekapTahun(T2 arrRekap, n int, xtahun int) int {
	for i := 0; i < n; i++ {
		if T2[i].tahun == xtahun {
			return i
		}
	}
	return -1
}

func Rekapitulasi(T1 arrPenelitian, T2 *arrRekap, n2 *int, n int) {

	for i := 0; i < n; i++ {
		idx := cariRekapTahun(*T2, *n2, T1[i].tahun)
		if idx == -1 {
			T2[*n2].tahun = T1[i].tahun
			T2[*n2].jumlah1 = 1
			*n2++
		} else {
			T2[idx].jumlah1++
		}

	}

}
func SortRekap(T *arrRekap, n2 int, x int) {
	pass := 1
	for pass <= n2-1 {
		idx := pass - 1
		i := pass
		if x == 1 {
			for i < n2 {
				if T[idx].jumlah1 < T[i].jumlah1 {
					idx = i
				}
				i++
			}
			T[pass-1], T[idx] = T[idx], T[pass-1]
			pass++
		} else {
			for i < n2 {
				if T[idx].jumlah1 > T[i].jumlah1 {
					idx = i
				}
				i++
			}
			T[pass-1], T[idx] = T[idx], T[pass-1]
			pass++
		}
	}
}
func printRekap(T arrRekap, n2 int) {
	fmt.Println("Tahun   | Jumlah Penelitian")
	fmt.Println("--------|-----------------")
	for i := 0; i < n2; i++ {
		fmt.Printf("%-8d| %d\n", T[i].tahun, T[i].jumlah1)
	}
}
func sequential_search(T arrPenelitian, n int, x string) int {
	var i, j int
	j = -1
	for i <= n && j == -1 {
		if T[i].judul == x {
			j = i
		}
		i++
	}
	return j
}
func mengeditData(T *arrPenelitian, n int) {
	var pilih, x, idx int
	var s, judul, opsibatal, opsi, bagian string
	var temp penelitian
	for {
		bersihLayar()

		fmt.Println("=======================================")
		fmt.Println("            MENGUBAH DATA   ")
		fmt.Println("=======================================")
		fmt.Print("Judul dari data yang ingin diubah : ")
		fmt.Scan(&judul)
		idx = sequential_search(*T, n, judul)
		if idx != -1 {
			temp = T[idx]

			fmt.Println(strings.Repeat("=", 200))
			fmt.Printf("%-20s%-15s%-15s%-15s%-15s%-70s%-10s%-15s%-15s\n", "KetuaPeneliti", "Anggota1", "Anggota2", "Anggota3", "Fakultas", "Judul", "Dana", "Luaran", "Tahun")
			fmt.Printf("%-20s%-15s%-15s%-15s%-15s%-70s%-10s%-15s%-15d\n", T[idx].ketua, T[idx].anggota1, T[idx].anggota2, T[idx].anggota3, T[idx].fakultas, T[idx].judul, T[idx].dana, T[idx].luaran, T[idx].tahun)
			fmt.Println(strings.Repeat("=", 200))
			fmt.Println("1. Ubah sebagian || 0. Seluruh ")
			fmt.Print("Pilihan: ")
			fmt.Scan(&opsi)
			fmt.Print("\n")
			for opsi == "1" {
				fmt.Print("Ubah bagian : ")
				fmt.Scan(&bagian)
				if bagian == "KetuaPeneliti" {
					fmt.Print("Nama ketua : ")
					fmt.Scan(&s)
					T[idx].ketua = s
				} else if bagian == "Anggota1" {
					fmt.Print("Nama anggota 1 : ")
					fmt.Scan(&s)
					T[idx].anggota1 = s
				} else if bagian == "Anggota2" {
					fmt.Print("Nama anggota 2 : ")
					fmt.Scan(&s)
					T[idx].anggota2 = s
				} else if bagian == "Anggota3" {
					fmt.Print("Nama anggota 3 : ")
					fmt.Scan(&s)
					T[idx].anggota3 = s
				} else if bagian == "Fakultas" {
					fmt.Print("Nama fakultas : ")
					fmt.Scan(&s)
					T[idx].fakultas = s
				} else if bagian == "Judul" {
					fmt.Print("Judul : ")
					fmt.Scan(&s)
					T[idx].judul = s
				} else if bagian == "Dana" {
					fmt.Print("Sumber dana (internal/eksternal) : ")
					fmt.Scan(&s)
					T[idx].dana = s
				} else if bagian == "Luaran" {
					fmt.Print("Luaran : ")
					fmt.Scan(&s)
					T[idx].luaran = s
				} else if bagian == "Tahun" {
					fmt.Print("Tahun : ")
					fmt.Scan(&x)
					T[idx].tahun = x
				}
				fmt.Println("1.Ubah bagian lain || 2.Selesai ")
				fmt.Print("Pilihan: ")
				fmt.Scan(&opsi)
				fmt.Print("\n")

			}
			if opsi == "0" {
				fmt.Print("Nama ketua : ")
				fmt.Scan(&s)
				T[idx].ketua = s

				fmt.Print("Nama anggota 1 : ")
				fmt.Scan(&s)
				T[idx].anggota1 = s

				fmt.Print("Nama anggota 2 : ")
				fmt.Scan(&s)
				T[idx].anggota2 = s

				fmt.Print("Nama anggota 3 : ")
				fmt.Scan(&s)
				T[idx].anggota3 = s

				fmt.Print("Nama fakultas : ")
				fmt.Scan(&s)
				T[idx].fakultas = s

				fmt.Print("Judul : ")
				fmt.Scan(&s)
				T[idx].judul = s

				fmt.Print("Sumber dana (internal/eksternal) : ")
				fmt.Scan(&s)
				T[idx].dana = s

				fmt.Print("Luaran : ")
				fmt.Scan(&s)
				T[idx].luaran = s

				fmt.Print("Tahun : ")
				fmt.Scan(&x)
				T[idx].tahun = x

			}
			fmt.Println(strings.Repeat("=", 200))
			fmt.Printf("%-20s%-15s%-15s%-15s%-15s%-70s%-10s%-15s%-15s\n", "Ketua Peneliti", "Anggota1", "Anggota2", "Anggota3", "Fakultas", "Judul", "Dana", "Luaran", "Tahun")
			fmt.Printf("%-20s%-15s%-15s%-15s%-15s%-70s%-10s%-15s%-15d\n", T[idx].ketua, T[idx].anggota1, T[idx].anggota2, T[idx].anggota3, T[idx].fakultas, T[idx].judul, T[idx].dana, T[idx].luaran, T[idx].tahun)
			fmt.Println(strings.Repeat("=", 200))

			fmt.Print("Konfirmasi pengubahan data : Ya/Tidak ", "\n")
			fmt.Scan(&opsibatal)
			if opsibatal == "Ya" {
				fmt.Println("Pengubahan data berhasil")
			} else {
				T[idx] = temp
				fmt.Println("Pengubahan data dibatalkan")
			}
		} else {
			fmt.Println("Judul tidak ditemukan")
		}
		fmt.Println("<<0.Ubah data lain>>", "\t", "<<1.Kembali ke Menu>>")
		fmt.Scan(&pilih)
		if pilih == 1 {
			break
		}
	}
	bersihLayar()
	showMenu()

}
func menghapusData(T *arrPenelitian, n int) {
	var pilih int
	var judul, opsibatal, bagian, opsi string
	var idx int
	var temp penelitian

	for {
		bersihLayar()

		fmt.Println("=======================================")
		fmt.Println("            MENGHAPUS DATA   ")
		fmt.Println("=======================================")
		fmt.Print("Judul dari data yang ingin dihapus : ")
		fmt.Scan(&judul)
		idx = sequential_search(*T, n, judul)
		if idx != -1 {
			temp = T[idx]

			fmt.Println(strings.Repeat("=", 200))
			fmt.Printf("%-20s%-15s%-15s%-15s%-15s%-70s%-10s%-15s%-15s\n", "Ketua Peneliti", "Anggota1", "Anggota2", "Anggota3", "Fakultas", "Judul", "Dana", "Luaran", "Tahun")
			fmt.Printf("%-20s%-15s%-15s%-15s%-15s%-70s%-10s%-15s%-15d\n", T[idx].ketua, T[idx].anggota1, T[idx].anggota2, T[idx].anggota3, T[idx].fakultas, T[idx].judul, T[idx].dana, T[idx].luaran, T[idx].tahun)
			fmt.Println(strings.Repeat("=", 200))
			fmt.Println("1. Hapus sebagian || 0. Seluruh ")
			fmt.Print("Pilihan: ")
			fmt.Scan(&opsi)
			fmt.Print("\n")
			for opsi == "1" {
				fmt.Print("Ubah bagian : ")
				fmt.Scan(&bagian)
				if bagian == "KetuaPeneliti" {
					T[idx].ketua = "-"
				} else if bagian == "Anggota1" {
					T[idx].anggota1 = "-"
				} else if bagian == "Anggota2" {
					T[idx].anggota2 = "-"
				} else if bagian == "Anggota3" {
					T[idx].anggota3 = "-"
				} else if bagian == "Fakultas" {
					T[idx].fakultas = "-"
				} else if bagian == "Judul" {
					T[idx].judul = "-"
				} else if bagian == "Dana" {
					T[idx].dana = "-"
				} else if bagian == "Luaran" {
					T[idx].luaran = "-"
				} else if bagian == "Tahun" {
					T[idx].tahun = -1
				}
				fmt.Println(strings.Repeat("=", 190))
				fmt.Printf("%-20s%-15s%-15s%-15s%-15s%-70s%-10s%-15s%-15s\n", "Ketua Peneliti", "Anggota1", "Anggota2", "Anggota3", "Fakultas", "Judul", "Dana", "Luaran", "Tahun")
				fmt.Printf("%-20s%-15s%-15s%-15s%-15s%-70s%-10s%-15s%-15d\n", T[idx].ketua, T[idx].anggota1, T[idx].anggota2, T[idx].anggota3, T[idx].fakultas, T[idx].judul, T[idx].dana, T[idx].luaran, T[idx].tahun)
				fmt.Println(strings.Repeat("=", 190))
				fmt.Println("1.Hapus bagian lain || 2.Selesai ")
				fmt.Print("Pilihan: ")
				fmt.Scan(&opsi)
				fmt.Print("\n")
			}
			if opsi == "0" {
				for i := idx; i <= n; i++ {
					T[i] = T[i+1]
				}
			}
			fmt.Print("Konfirmasi penghapusan data : Ya/Tidak ", "\n")
			fmt.Scan(&opsibatal)
			if opsibatal == "Ya" {
				fmt.Println("Penghapusan data berhasil")
			} else {
				T[idx] = temp
				fmt.Println("Penghapusan data dibatalkan")
			}
		} else {
			fmt.Println("Judul tidak ditemukan")
		}
		fmt.Println("<<0.Hapus data lain>>", "\t", "<<1.Kembali ke Menu>>")
		fmt.Scan(&pilih)
		if pilih == 1 {
			break
		}
	}
	bersihLayar()
	showMenu()

}
func binarySearch(t *arrPenelitian, n int) {
	var x string
	namaSortBinary(&*t, n)
	fmt.Println("Cari Data berdasarkan Judul :")

	fmt.Scan(&x)
	kr := 0
	kn := n - 1
	found := -1
	for kr <= kn && found == -1 {
		mid := (kr + kn) / 2
		if x > t[mid].judul {
			kr = mid + 1
		} else if x < t[mid].judul {
			kn = mid - 1
		} else {
			found = mid
		}
	}

	if found == -1 {
		fmt.Println("Data tidak ditemukan")
	} else {
		fmt.Println(strings.Repeat("=", 190))
		fmt.Printf("%-20s%-15s%-15s%-15s%-15s%-70s%-10s%-15s%-15s\n", "Ketua Peneliti", "Anggota1", "Anggota2", "Anggota3", "Fakultas", "Judul", "Dana", "Luaran", "Tahun")
		fmt.Println(strings.Repeat("=", 190))
		fmt.Printf("%-20s%-15s%-15s%-15s%-15s%-70s%-10s%-15s%-15d\n", T[found].ketua, T[found].anggota1, T[found].anggota2, T[found].anggota3, T[found].fakultas, T[found].judul, T[found].dana, T[found].luaran, T[found].tahun)

	}
	fmt.Println(strings.Repeat("=", 190))

}

func farewell() {
	fmt.Println("==========================================================================")
	fmt.Println("  Terima kasih telah menggunakan aplikasi Tri Dharma Perguruan Tinggi!")
	fmt.Println("==========================================================================")
}
