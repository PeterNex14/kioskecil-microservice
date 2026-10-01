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

