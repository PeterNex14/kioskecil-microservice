# Story KK-INV-02: Konfigurasi Multi-Satuan & Harga Grosir (Packaging Formula)

* **Modul**: Inventory & Master Data
* **Target Service**: `inventory-service` (`db_inventory`)
* **Persona**: Owner (*Decision Maker*)

---

## 1. Narrative
* **As an** Owner (*Decision Maker*),
* **I want to** mengonfigurasi satu atau lebih satuan kemasan bertingkat (misal: Renceng, Box, Dus) dengan rasio konversi ke satuan dasar, barcode kemasan opsional, dan harga jual grosir khusus,
* **So that** rumus konversi terkunci satu kali di awal dan sistem otomatis memotong stok satuan dasar serta membedakan harga eceran vs grosir tanpa perlu hitung ulang manual.

---

## 2. Scope Boundary (Out of Scope)
* Pencatatan transaksi belanja kulakan dari supplier ➔ Ditangani di [KK-INV-03](KK-INV-03.md).
* Penghitungan keranjang hybrid/campuran saat checkout kasir (misal: 1 Dus + 5 Pcs) ➔ Ditangani di modul POS (`KK-SAL-04`).

---

## 3. Acceptance Criteria (AC)

### Scenario 1: Menambahkan Satuan Kemasan Bertingkat Berhasil (Happy Path)
* **Given** Produk *"Kopi Kapal Api"* sudah terdaftar dengan satuan dasar `Sachet` (Harga eceran: Rp 1.500/sachet),
* **When** Owner menambahkan dua satuan kemasan:
  1. Satuan `Renceng` (Faktor Konversi: 10 Sachet, Harga grosir: Rp 13.500).
  2. Satuan `Dus` (Faktor Konversi: 120 Sachet, Harga grosir: Rp 150.000),
* **Then** Sistem menyimpan kedua level kemasan tersebut dengan relasi ke produk utama, dan keduanya aktif serta siap digunakan di kasir.

### Scenario 2: Registrasi Barcode Khusus Kemasan Dus/Box
* **Given** Produk *"Beng-Beng"* memiliki barcode kemasan kardus/box dari pabrik,
* **When** Owner menginput barcode kemasan `899275322110` pada satuan `Box` (Isi 20 Pcs),
* **Then** Sistem menyimpan barcode tersebut sebagai barcode khusus satuan kemasan. Saat kasir men-scan barcode ini, sistem langsung mengenali produk sebagai 1 Box Beng-Beng dengan harga box.

### Scenario 3: Validasi Barcode Kemasan Duplikat (Negative Path)
* **Given** Barcode `8999999999` sudah terdaftar pada produk atau satuan lain,
* **When** Owner mencoba menggunakan barcode yang sama pada satuan kemasan baru,
* **Then** Sistem menolak penyimpanan dengan status `409 Conflict` dan pesan: *"Barcode kemasan sudah digunakan oleh produk lain"*.

### Scenario 4: Validasi Faktor Konversi Tidak Valid (Negative Path)
* **Given** Owner sedang menambahkan satuan kemasan baru,
* **When** Owner memasukkan faktor konversi bernilai `0`, negatif, atau `1` (misal 1 Dus = 1 Pcs),
* **Then** Sistem menolak dengan status `400 Bad Request` dan pesan: *"Faktor konversi kemasan harus bilangan bulat lebih besar dari 1"*.

---

## 4. Business Rules & Edge Cases
- [ ] **One-Time Formula Setup**: Rumus konversi di-setup satu kali di master produk. Setelah tersimpan, Owner tidak perlu lagi menginput rasio konversi saat restock atau saat kasir berjualan.
- [ ] **Single Source of Physical Stock**: Total stok fisik produk **selalu disimpan dan dihitung dalam satuan terkecil (*base unit*)** di database. Jika kasir menjual 1 Dus Indomie (isi 40 Pcs), sistem otomatis mengurangi 40 Pcs dari total stok fisik.
- [ ] **Barcode Bersifat Opsional**: Satuan kemasan (seperti Renceng yang diikat karet manual) boleh tidak memiliki barcode. Kasir tetap bisa memilihnya melalui antarmuka kasir.
- [ ] **Independent Packaging Price**: Harga jual kemasan tidak otomatis mengalikan harga eceran, melainkan ditentukan bebas oleh Owner (agar harga grosir bisa lebih murah dari eceran).
- [ ] **Deletion Safety Constraint**: Satuan kemasan tidak boleh dihapus jika sudah pernah memiliki riwayat transaksi penjualan (hanya boleh dinonaktifkan / *soft delete*).
- [ ] **Data Reference**: Detail field database mengacu pada tabel `product_units` di [data-dictionary.md](data-dictionary.md).
