# Story KK-INV-03: Pencatatan Barang Masuk (Restock) & Otomatisasi HPP/COGS

* **Modul**: Inventory & Master Data
* **Target Service**: `inventory-service` (`db_inventory`)
* **Persona**: Owner (*Decision Maker*)

---

## 1. Narrative
* **As an** Owner (*Decision Maker*),
* **I want to** mencatat penerimaan barang belanjaan supplier/grosir (Restock) dalam satuan apa pun (Dus/Pack/Pcs) beserta harga belinya,
* **So that** stok fisik di gudang otomatis bertambah sesuai satuan terkecil dan harga modal (COGS) per satuan eceran otomatis terhitung tanpa kalkulator manual.

---

## 2. Scope Boundary (Out of Scope)
* Pengurangan stok karena barang rusak/hilang/dimakan rayap ➔ Ditangani di [KK-INV-05](KK-INV-05.md) *(Stock Adjustment)*.
* Integrasi pembayaran hutang tempo ke supplier ➔ Ditangani di modul Keuangan *(Finance Service)*.

---

## 3. Acceptance Criteria (AC)

### Scenario 1: Restock Satuan Kemasan Dus Otomatis Konversi ke Eceran (Happy Path)
* **Given** Produk *"Indomie Goreng"* memiliki stok fisik `10 Pcs`, dan rumus kemasan `1 Dus = 40 Pcs`,
* **When** Owner mencatat restock: memilih satuan `Dus`, kuantitas `5 Dus`, dan harga beli `Rp 116.000 / Dus` (Total nota: Rp 580.000), lalu menyimpan,
* **Then**:
  1. Sistem otomatis menambahkan `+200 Pcs` (5 Dus x 40 Pcs) ke stok fisik, sehingga total stok menjadi `210 Pcs`.
  2. Sistem otomatis menghitung harga modal eceran baru: `Rp 116.000 / 40 = Rp 2.900 / Pcs`.
  3. Sistem mencatat riwayat mutasi masuk ke tabel audit stok `stock_movements`.

### Scenario 2: Deteksi Kenaikan Harga Modal Supplier (Smart Notification)
* **Given** Harga modal eceran lama adalah `Rp 2.800 / Pcs` dan harga jual eceran adalah `Rp 3.500 / Pcs`,
* **When** Owner melakukan restock dengan harga baru yang menghasilkan modal eceran `Rp 2.900 / Pcs` (modal naik Rp 100),
* **Then** Sistem memperbarui harga modal produk dan menampilkan notifikasi info:
  *"Harga modal naik menjadi Rp 2.900/Pcs (Margin keuntungan turun menjadi 17.1%). Apakah Anda ingin memperbarui harga jual eceran?"* dengan opsi tombol cepat untuk ubah harga jual.

### Scenario 3: Validasi Input Restock (Negative Path)
* **Given** Owner sedang berada di form restock,
* **When** Owner menginput jumlah barang `<= 0` atau harga beli minus,
* **Then** Sistem menolak penyimpanan dengan error `400 Bad Request`: *"Jumlah barang masuk dan harga beli harus lebih besar dari 0"*.

---

## 4. Business Rules & Edge Cases
- [ ] **Stock Ledger (Audit Trail)**: Setiap restock wajib membuat baris baru di tabel `stock_movements` (mencatat: tipe `RESTOCK`, referensi satuan yang dibeli, kuantitas satuan dasar yang masuk, harga beli, dan siapa yang menginput).
- [ ] **Restock Satuan Dasar**: Owner juga bisa mencatat restock langsung dalam satuan dasar `Pcs` (misal untuk barang non-dus seperti kerupuk).
- [ ] **Nama Supplier / Catatan Nota**: Form restock menyediakan kolom opsional untuk mencatat nomor nota atau nama toko grosir asal belanjaan.
- [ ] **Data Reference**: Detail field database mengacu pada tabel `stock_movements` di [data-dictionary.md](data-dictionary.md).
