# Story KK-INV-01: Registrasi Master Produk & Penetapan Harga Dasar

* **Modul**: Inventory & Master Data
* **Target Service**: `inventory-service` (`db_inventory`)
* **Persona**: Owner (*Decision Maker*)

---

## 1. Narrative
* **As an** Owner (*Decision Maker*),
* **I want to** mendaftarkan produk baru dengan nama, barcode/SKU, kategori, serta menetapkan harga modal (COGS) dan harga jual dasarnya,
* **So that** sistem memiliki referensi harga resmi yang terkunci sehingga kasir tidak salah menyebut harga dan margin keuntungan tetap aman.

---

## 2. Scope Boundary (Out of Scope)
* Pengaturan satuan kemasan sekunder (seperti Dus/Pack/Renceng) ➔ Ditangani di [KK-INV-02](KK-INV-02.md).
* Input kuantitas stok fisik awal atau penerimaan barang supplier (*Restock*) ➔ Ditangani di [KK-INV-03](KK-INV-03.md).

---

## 3. Acceptance Criteria (AC)

### Scenario 1: Registrasi Produk Baru Berhasil (Happy Path)
* **Given** Owner telah terautentikasi dan mengakses antarmuka tambah produk,
* **When** Owner menginput nama produk valid (misal: *"Indomie Goreng Original"*), barcode valid, kategori (misal: *"Makanan Instan"*), harga modal `Rp 2.800`, dan harga jual `Rp 3.500`, lalu menyimpan,
* **Then** Sistem menyimpan produk baru dengan status `is_active = true`, mengembalikan status `201 Created`, dan produk langsung terindeks untuk pencarian kasir.

### Scenario 2: Registrasi Produk Tanpa Barcode (Produk Tradisional / Non-Pabrik)
* **Given** Owner mendaftarkan produk yang tidak memiliki barcode fisik dari pabrik (misal: *"Kerupuk Kaleng Putih"* atau *"Es Batu"*),
* **When** Owner mengosongkan kolom Barcode dan melengkapi data lainnya,
* **Then** Sistem otomatis men-generate kode SKU lokal unik (format: `KIOS-PRD-XXXXX`) dan produk berhasil disimpan tanpa error.

### Scenario 3: Validasi Barcode Duplikat (Negative Path)
* **Given** Barcode `8992775211114` sudah terdaftar pada produk *"Teh Botol Sosro"*,
* **When** Owner mencoba menyimpan produk baru menggunakan barcode yang sama,
* **Then** Sistem menolak permintaan, mengembalikan status `409 Conflict`, dan menampilkan pesan: *"Barcode sudah terdaftar pada produk Teh Botol Sosro"*.

### Scenario 4: Validasi Harga Jual Tidak Masuk Akal (Negative Path)
* **Given** Owner sedang mengisi harga produk,
* **When** Owner memasukkan harga jual bernilai `0` atau minus (`<= 0`),
* **Then** Sistem menolak input dengan status `400 Bad Request` dan pesan: *"Harga jual harus lebih besar dari 0"*.

### Scenario 5: Peringatan Jual Rugi / Margin Minus (Soft Warning)
* **Given** Owner memasukkan harga modal `Rp 30.000` dan harga jual `Rp 3.000` (terindikasi salah ketik / typo),
* **When** Owner menekan tombol simpan,
* **Then** Sistem menampilkan dialog konfirmasi: *"Peringatan: Harga jual lebih rendah dari harga modal (Margin minus Rp 27.000). Tetap simpan produk?"*, dan hanya memproses penyimpanan jika Owner mengonfirmasi "Ya".

---

## 4. Business Rules & Edge Cases
- [ ] **Role Authorization**: Hanya pengguna dengan role `owner` yang memiliki izin membuat, mengubah, atau menonaktifkan master produk. Role `cashier` hanya memiliki izin baca (*read-only*).
- [ ] **Kategori Default**: Kategori bersifat opsional. Jika tidak diisi, sistem otomatis menyematkan kategori *"Umum"* (*Uncategorized*).
- [ ] **Data Sanitization**: Nama produk otomatis dibersihkan dari spasi berlebih (*trim*) di awal dan akhir teks.
- [ ] **COGS Privacy**: Kolom harga modal (`cost_price`) tidak boleh disertakan dalam respons API yang diakses oleh kasir (*data masking*).
- [ ] **Data Reference**: Detail field database mengacu pada tabel `products` di [data-dictionary.md](data-dictionary.md).
