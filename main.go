package main

import (
	"fmt" // fmt paketi, formatlı I/O işlemleri için kullanılır
)

func main() {
	var bakiye float64 = 1000.0 // Başlangıç bakiyesi 1000 TL olarak tanımlanır
	var secim int               // Kullanıcının işlem seçimini tutar
	var miktar float64          // Kullanıcının işlem miktarını tutar

	fmt.Println("Hesabınızda 1000 TL var.")                                             // Başlangıç bakiyesini ekrana yazdırır
	fmt.Println("Para çekmek istiyorsanız 1'e, para yatırmak istiyorsanız 2'ye basın:") // Kullanıcıdan işlem seçmesini ister
	fmt.Scan(&secim)                                                                    // Kullanıcının seçimini alır

	switch secim {
	case 1:
		fmt.Println("Ne kadar çekmek istiyorsunuz?") // Para çekme miktarını sorar
		fmt.Scan(&miktar)                            // Kullanıcının çekmek istediği miktarı alır
		if miktar > bakiye {                         // Çekmek istenen miktar bakiyeden fazla ise
			fmt.Println("Bu işlem yapılamaz. Yetersiz bakiye.") // Uyarı mesajı gösterir
		} else if miktar < 0 { // Negatif miktar girilmişse
			fmt.Println("Bu işlem yapılamaz. Negatif tutar çekilemez.") // Uyarı mesajı gösterir
		} else {
			bakiye -= miktar                                       // Geçerli miktarı bakiyeden düşer
			fmt.Printf("Hesabınızda artık %.2f TL var.\n", bakiye) // Yeni bakiyeyi ekrana yazdırır
		}
	case 2:
		fmt.Println("Ne kadar yatırmak istiyorsunuz?") // Para yatırma miktarını sorar
		fmt.Scan(&miktar)                              // Kullanıcının yatırmak istediği miktarı alır
		if miktar < 0 {                                // Negatif miktar girilmişse
			fmt.Println("Bu işlem yapılamaz. Negatif tutar yatırılamaz.") // Uyarı mesajı gösterir
		} else {
			bakiye += miktar                                       // Geçerli miktarı bakiyeye ekler
			fmt.Printf("Hesabınızda artık %.2f TL var.\n", bakiye) // Yeni bakiyeyi ekrana yazdırır
		}
	default:
		fmt.Println("Geçersiz seçenek.") // Geçersiz bir seçim yapılmışsa uyarı mesajı gösterir
	}
}
