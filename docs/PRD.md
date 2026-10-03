# Product Requirements Document (PRD): KiosKecil

* **Status**: In Progress (Drafting Module-by-Module)
* **Version**: 0.1
* **Target Audience**: Penjaga Warung (*Daily Operator*) & Pemilik Warung (*Decision Maker*)

---

## 1. Product Overview & Vision
**KiosKecil** adalah sistem kasir (POS) dan manajemen inventaris cerdas yang dirancang khusus untuk operasional warung/toko kelontong. 

Tujuan utama KiosKecil adalah mentransformasi operasional warung yang saat ini masih manual, terfragmentasi, dan bergantung pada ingatan/buku catatan fisik menjadi sistem terpadu (*single source of truth*) yang menyinkronkan stok, harga jual, dan arus kas secara akurat dan *real-time*.

---

## 2. Core Problem & Impact

### A. Masalah Inti (Core Problems)
1. **Tidak Ada Single Source of Truth**: Harga barang dan jumlah stok hanya diingat di kepala atau tercecer di berbagai catatan fisik.
2. **Kebutaan Finansial (Financial Invisibility)**: Tidak ada kejelasan apakah hari ini warung untung atau rugi. Arus kas harian dan bulanan tidak terukur.
3. **Pencatatan Utang & Biaya Manual**: Kasbon pelanggan, hutang ke supplier, dan biaya harian (listrik, kantong plastik) masih manual, sering lupa dicatat, atau catatannya hilang.

### B. Dampak Nyata (Impacts)
* **Harga Tidak Konsisten**: Kasir sering salah sebut harga (terlalu murah membuat margin bocor/rugi, terlalu mahal memicu komplain pelanggan).
* **Keputusan Bisnis "Buta Data"**: Sulit mengevaluasi laba bersih, menentukan strategi harga, atau memilih barang mana yang perlu dipromosikan.
* **Restock Tanpa Data**: Modal warung "mati" karena membeli barang yang belum habis, sementara barang yang laku keras justru kehabisan stok.
* **Miskomunikasi Utang/Piutang**: Terjadi perselisihan dengan penagih supplier atau kecanggungan saat menagih kasbon tetangga/pelanggan.

---

## 3. Target Solusi MVP (Minimum Viable Product)
Membangun satu pusat data terpadu (*Single Source of Truth*) yang mampu:
1. **Mengunci Harga Jual & COGS**: Menyimpan Harga Pokok Penjualan (HPP/COGS) dan harga jual baku per satuan agar konsisten.
2. **Otomatisasi Stok & Konversi Multi-Satuan**: Otomatis potong stok saat penjualan dan tambah stok saat restock dengan konversi satuan fleksibel (misal: 1 Dus = 40 Pcs).
3. **Laporan Untung-Rugi (P&L) Real-Time**: Memberikan visibilitas instan atas laba kotor dan laba bersih harian/bulanan.

---

## 4. User Personas

### 👤 Persona 1: The Daily Operator (Penjaga Warung / Kasir)
* **Profil**: Siapa pun yang bertugas menjaga shift warung.
* **Karakteristik**:
  * *Multi-tasking*: Melayani pembeli sambil membongkar muatan kurir supplier.
  * *Mobilitas tinggi*: Sering bolak-balik antara meja kasir dan rak gudang.
  * *Speed-oriented*: Butuh kecepatan akses data saat antrean sedang ramai.
* **Tanggung Jawab Operasional**:
  * Input transaksi retail maupun grosir dengan cepat.
  * Input penerimaan barang masuk (*stock receiving*).
  * Mencatat kasbon/hutang pelanggan secara instan.
  * Mengecek ketersediaan stok barang secara cepat.
* **Kebutuhan Utama (Needs)**:
  * Antarmuka (UI) yang nyaman di HP/Tablet dengan tombol yang mudah ditekan.
  * Pencarian cerdas (ketik *"Indo"* langsung muncul semua varian Indomie).
  * Konversi otomatis (sistem otomatis tahu 1 Dus = 40 Pcs tanpa perlu kalkulator manual).
  * Fitur **Hold / Resume Transaksi** saat ada pelanggan yang menyela antrean.
* **Titik Lelah (Pain Points)**:
  * Repot membuka buku catatan fisik saat toko sedang ramai.
  * Lupa mencatat orang yang kasbon karena sibuk melayani pembeli lain.
  * Tekanan saat harus membatalkan keranjang belanja pembeli A hanya untuk melayani pembeli B yang "nyerobot" antrean.
  * Takut salah hitung kembalian atau salah memberikan total harga.

---

### 👤 Persona 2: The Decision Maker (Pemilik Warung / Analyst)
* **Profil**: Pengelola bisnis, pemilik modal, dan pemegang keputusan akhir.
* **Karakteristik**:
  * Fokus pada angka akhir (*bottom line*).
  * Memikirkan strategi penetapan harga dan perencanaan belanja modal.
* **Tanggung Jawab Manajerial**:
  * *Price Control*: Menetapkan harga jual pusat agar margin keuntungan aman dan konsisten.
  * *Expense Tracking*: Mencatat pengeluaran operasional non-stok (listrik, plastik kresek, keamanan).
  * *Profit-Loss Analysis*: Meninjau kesehatan finansial: *"Omzet bulan ini besar, tapi uang fisiknya ke mana?"*
  * *Stock Strategy*: Memutuskan barang mana yang harus didiskon (*slow-moving*) dan barang mana yang wajib di-restock (*best-seller*).
* **Kebutuhan Utama (Needs)**:
  * Dashboard ringkas: Total Omzet, Total Laba Kotor, Laba Bersih, dan Total Piutang belum lunas.
  * Filter rentang waktu laporan (Harian, Mingguan, Bulanan).
  * Keamanan dan integritas data (role-based access, riwayat transaksi tidak bisa dimanipulasi kasir).
* **Titik Lelah (Pain Points)**:
  * Miskomunikasi saat harga modal supplier naik tapi kasir masih menjual dengan harga lama.
  * Ketidaksesuaian antara uang tunai fisik di laci kasir vs catatan penjualan di sistem.

---

## 5. Rencana Modul MVP
Sistem KiosKecil terbagi ke dalam 4 modul utama yang akan didefinisikan secara bertahap bersama Acceptance Criteria-nya:
1. **Modul 1: Inventory & Master Data** (Master Produk, Multi-satuan Dus/Pcs, Pricing, Penyesuaian Stok)
2. **Modul 2: Sales & POS / Transaction** (Katalog/Pencarian, Keranjang, Hold/Resume, Split Payment, Struk)
3. **Modul 3: Finance & Debt** (Buku Kasbon/Piutang Pelanggan, Pencatatan Biaya Operasional)
4. **Modul 4: Reporting & Analytics** (Dashboard Real-time, Laporan P&L, Analisis Best/Slow Moving)

## 6. Spesifikasi Modul Detail (Modular PRD)
Setiap modul memiliki dokumen spesifikasi mandiri yang mencakup User Stories, Acceptance Criteria (AC), dan Data Dictionary:

* 📦 **Modul 1: Inventory & Master Data** ➔ Baca [docs/prd/inventory/README.md](prd/inventory/README.md) *(Sedang Disusun)*
* 🛒 **Modul 2: Sales & POS (Transaction)** ➔ Baca `docs/prd/02-sales-pos.md` *(Upcoming)*
* 💰 **Modul 3: Finance & Debt** ➔ Baca `docs/prd/03-finance-debt.md` *(Upcoming)*
* 📊 **Modul 4: Reporting & Analytics** ➔ Baca `docs/prd/04-reporting.md` *(Upcoming)*

---

## 7. Functional Requirements (FR)

Kebutuhan fungsional mendefinisikan kapabilitas dan fitur sistem yang harus disediakan oleh platform KiosKecil:

### 📦 A. Modul Inventory & Master Data (`inventory-service`)
Bertanggung jawab atas pengelolaan data barang, harga, dan pergerakan stok sebelum dijual:
* **`[FR-INV-01]` Manajemen Produk**: Sistem harus dapat menyimpan data pokok produk (Nama, Barcode/SKU unik, Kategori).
* **`[FR-INV-02]` Multi-Unit Conversion**: Sistem harus mendukung minimal 2 level satuan (misal: Dus dan Pcs) dengan satu faktor konversi tetap (`> 1`).
* **`[FR-INV-03]` Pricing Logic**: Sistem harus menyimpan Harga Modal (COGS/HPP) dan Harga Jual resmi untuk setiap level satuan.
* **`[FR-INV-04]` Stock Management**: Sistem harus otomatis menambah stok saat *restock* dan mengurangi stok saat *sales checkout* (berdasarkan konversi ke satuan dasar).
* **`[FR-INV-05]` Stock Adjustment**: Sistem harus menyediakan fitur manual untuk mengoreksi kuantitas stok fisik dengan alasan terstandarisasi (*Rusak*, *Expired*, *Pemakaian Pribadi/Owner Use*, *Selisih Opname*).
* **`[FR-INV-06]` Low Stock Alert**: Sistem harus memberikan tanda visual jika kuantitas stok berada di bawah ambang batas minimum yang ditentukan Owner.

### 🛒 B. Modul Sales & POS Transaction (`pos-service`)
Mesin transaksi utama yang digunakan pada meja kasir operasional harian:
* **`[FR-SAL-01]` Point of Sale Interface**: Sistem harus menyediakan pencarian produk secara cepat (pencarian instan nama atau scan barcode fisik).
* **`[FR-SAL-02]` Cart Management**: Sistem harus dapat menambah, mengubah kuantitas, dan menghapus item dalam keranjang belanja sebelum checkout.
* **`[FR-SAL-03]` Hold / Resume Transaction**: Sistem harus mampu menyimpan minimal 5 transaksi "Draft" sekaligus untuk menangani interupsi antrean pembeli.
* **`[FR-SAL-04]` Hybrid Calculation**: Sistem harus otomatis menghitung total harga jika dalam satu keranjang terdapat kombinasi satuan kemasan (misal: 1 Dus + 5 Pcs Indomie).
* **`[FR-SAL-05]` Flexible Pricing (Service / Pulsa)**: Sistem harus mendukung produk bertipe "Jasa" / non-fisik di mana harga jual (*Open Price*) dan biaya admin dapat diinput fleksibel saat transaksi tanpa memotong stok.
* **`[FR-SAL-06]` Payment Processing**: Sistem harus mendukung berbagai metode bayar: Tunai (dengan hitungan kembalian otomatis), Kasbon / Hutang Pelanggan, dan Campuran (*Split Payment*).
* **`[FR-SAL-07]` Receipt Generation**: Sistem harus menghasilkan ringkasan transaksi (Struk) yang dapat dicetak fisik maupun dibagikan secara digital.

### 💰 C. Modul Finance & Debt (`finance-service`)
Mencatat arus kas masuk dan keluar di luar transaksi langsung penjualan barang dagangan:
* **`[FR-FIN-01]` Expense Logging**: Sistem harus mampu mencatat biaya operasional warung non-stok berdasarkan kategori (Listrik, Air, Kantong Plastik, Keamanan, Kebersihan, dll).
* **`[FR-FIN-02]` Debt Ledger**: Sistem harus mencatat buku besar piutang per pelanggan dan menyediakan fitur pencatatan cicilan atau pelunasan hutang.
* **`[FR-FIN-03]` Profit Calculation**: Sistem harus menghitung Laba Kotor (*Gross Profit = Sales - COGS*) secara otomatis setiap transaksi selesai.

### 📊 D. Modul Reporting & Analytics (`reporting-service`)
Pengolahan data transaksi menjadi informasi dan wawasan bisnis bagi pemilik warung:
* **`[FR-REP-01]` Real-time Dashboard**: Menampilkan ringkasan harian performa warung (Total Omzet, Total Laba Kotor, dan Kasbon baru hari ini).
* **`[FR-REP-02]` P&L Statement**: Menghasilkan laporan laba rugi komprehensif berkala (*Laba Bersih = Laba Kotor - Beban Operasional*).
* **`[FR-REP-03]` Rank Analysis**: Menampilkan peringkat produk terlaris (*Top 10 Best Seller*) dan produk macet (*Top 10 Slow-Moving items > 30 hari*).
* **`[FR-REP-04]` Inventory Turnover Insights**: Menampilkan estimasi perputaran stok dan sisa hari sebelum stok habis berdasarkan rata-rata penjualan harian.

---

## 8. Non-Functional Requirements (NFR)

Kebutuhan non-fungsional mendefinisikan batasan teknis, kualitas performa, keamanan, dan keandalan arsitektur KiosKecil:

### ⚡ 1. Performance (Kecepatan & Responsivitas)
* **`[NFR-PER-01]` Response Time**: Pencarian produk (*Search / Barcode Scan*) harus memberikan hasil dalam waktu **< 500 ms**, agar kasir tidak terhambat saat melayani antrean panjang.
* **`[NFR-PER-02]` Startup Time**: Service backend Go harus dapat *booting* dan siap menerima request dalam waktu **< 5 detik** saat server dinyalakan.

### 🛡️ 2. Reliability & Availability (Keandalan & Ketahanan Data)
* **`[NFR-REL-01]` Data Persistence (ACID)**: Setiap transaksi yang berstatus lunas (*Completed*) wajib tersimpan permanen di database PostgreSQL secara atomik. Jika listrik padam atau sistem mati mendadak, data transaksi tidak boleh rusak (*corrupt*) atau hilang.
* **`[NFR-REL-02]` Offline-First Mindset**: Sistem dirancang untuk operasional warung lokal, sehingga harus dapat berjalan stabil di jaringan lokal / localhost tanpa ketergantungan koneksi internet publik.

### 🔒 3. Security (Keamanan Akses & Integritas)
* **`[NFR-SEC-01]` Role-Based Access Control (RBAC)**: Terdapat pemisahan wewenang yang tegas antara:
  * **Kasir (`cashier`)**: Hanya diizinkan mengakses transaksi kasir, lookup katalog, dan riwayat shift kasir hari ini.
  * **Pemilik (`owner`)**: Memiliki kontrol penuh atas master data, konfigurasi harga modal (COGS), penyesuaian stok, pengeluaran, dan seluruh laporan finansial.
* **`[NFR-SEC-02]` Strict Input Validation**: Sistem wajib menolak input angka negatif atau harga nol pada entri yang tidak sah untuk mencegah *human error* dan bug integritas database.

### 🏗️ 4. Maintainability & Observability (Kemudahan Pemeliharaan)
* **`[NFR-MAI-01]` Clean Layered Architecture**: Kode microservice Go wajib mengikuti pola *Layered Clean Architecture* (`handler` -> `service` -> `repository`) dengan *Database-per-Service* agar modular dan mudah diperluas (misal: integrasi pembayaran QRIS di masa depan).
* **`[NFR-MAI-02]` Structured Logging**: Seluruh error operasional dan mutasi penting wajib dicatat menggunakan structured logger bawaan Go (`log/slog` via `common/logger`) berpasangan *key-value*, tanpa menggunakan `fmt.Println` atau `panic()`.

### ⌨️ 5. Usability (Kemudahan Penggunaan Operasional)
* **`[NFR-USA-01]` Keyboard-Friendly Navigation**: Antarmuka kasir dirancang ergonomis dengan dukungan *shortcut keyboard* untuk navigasi cepat (pencarian barang, tambah kuantitas, tahan antrean, dan pemicu pembayaran) tanpa mewajibkan penggunaan mouse/trackpad.


