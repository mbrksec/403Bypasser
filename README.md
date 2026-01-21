🛡️ 403Bypasser
403Bypasser, web uygulamalarındaki erişim engellerini (403 Forbidden) aşmak için çeşitli teknikleri otomatize eden, Go (Golang) ile yazılmış yüksek performanslı bir güvenlik aracıdır.

🚀 Özellikler
Verb Tampering: Hedef URL'ye farklı HTTP metotları (GET, POST, PUT, TRACE vb.) ile istek atarak yetkilendirme hatalarını test eder.

Header Manipulation: X-Forwarded-For, X-Rewrite-URL ve X-Remote-IP gibi kritik header bilgilerini manipüle ederek IP tabanlı kısıtlamaları atlatmaya çalışır.

Concurrency: Go'nun goroutine yapısını kullanarak tarama işlemlerini eşzamanlı ve hızlı bir şekilde gerçekleştirir.

Simple Interface: Komut satırı üzerinden kolayca kullanılabilir.

🛠️ Kurulum
Sisteminizde Go yüklü olduğundan emin olun, ardından aşağıdaki komutları izleyin:

Bash
# Depoyu klonlayın
git clone https://github.com/mbrksec/403Bypasser.git

# Proje dizinine gidin
cd 403Bypasser

# Bağımlılıkları başlatın ve derleyin
go mod init 403bypasser

go build main.go

📖 Kullanım
Aracı çalıştırmak için hedef URL'yi parametre olarak vermeniz yeterlidir:

go run main.go https://example.com/admin

⚖️ Yasal Uyarı
Bu araç sadece eğitim ve yasal sızma testi (pentesting) süreçleri için geliştirilmiştir. İzin alınmamış sistemler üzerinde kullanılması yasal sorumluluk doğurabilir. Kullanıcı, bu aracın kullanımından kaynaklanan tüm sorumluluğu kabul eder.
