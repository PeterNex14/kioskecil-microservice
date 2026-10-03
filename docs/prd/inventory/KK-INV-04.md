# Story KK-INV-04: Pengaturan Batas Stok Minimum & Peringatan Stok Menipis (Low Stock Alert)

* **Modul**: Inventory & Master Data
* **Target Service**: `inventory-service` (`db_inventory`)
* **Persona**: Owner (*Decision Maker*)
* **Referensi FR**: `[FR-INV-06]` (Low Stock Alert)
* **Referensi NFR**: `[NFR-SEC-01]` (RBAC), `[NFR-SEC-02]` (Input Validation)

---

## 1. Narrative
* **As an** Owner (*Decision Maker*),
* **I want to** mengatur ambang batas stok minimum (*low stock threshold*) pada setiap produk dan melihat rekap daftar barang yang persediaannya menipis,
* **So that** saya mendapat peringatan visual tepat waktu sebelum barang habis total dan dapat segera menyusun daftar belanjaan kulakan ke pasar/supplier secara cepat tanpa memeriksa rak fisik satu per satu.

---

## 2. Scope Boundary (Out of Scope)
* Pengurangan kuantitas stok saat transaksi penjualan di kasir ➔ Ditangani di modul POS (`KK-CSH-02`, `KK-CSH-05`).
* Penambahan kuantitas stok fisik saat belanjaan kulakan datang (*Restock*) ➔ Ditangani di [KK-INV-03](KK-INV-03.md).
* Pengiriman notifikasi eksternal (misal: SMS Gateway, WhatsApp blast, atau push notification ke supplier) ➔ *Out of scope* untuk MVP warung.

---

## 3. Acceptance Criteria (AC)

### Scenario 1: Konfigurasi Batas Stok Minimum Berhasil (Happy Path)
* **Given** Owner telah terautentikasi dan membuka form ubah produk *"Beras Ramos 5kg"*,
* **When** Owner memasukkan nilai `5` pada kolom `min_stock_alert` dan menekan simpan,
* **Then** Sistem menyimpan nilai ambang batas tersebut di database dengan status `200 OK`, dan produk kini memiliki batas peringatan sebesar `5 Pcs`.

### Scenario 2: Validasi Input Angka Negatif (Negative Path)
* **Given** Owner sedang mengatur batas minimum stok produk,
* **When** Owner memasukkan angka negatif (misal: `-2`),
* **Then** Sistem menolak penyimpanan dengan status `400 Bad Request` dan pesan error: *"Batas minimum stok tidak boleh bernilai negatif"*.

### Scenario 3: Menonaktifkan Peringatan untuk Produk Tertentu (Optional Alert)
* **Given** Produk musiman atau produk titipan yang tidak rutin distok (misal: *"Sirup Marjan"* di luar bulan puasa),
* **When** Owner mengatur nilai `min_stock_alert = 0` (atau mengosongkan kolom),
* **Then** Sistem memperlakukan fitur peringatan stok untuk barang tersebut sebagai **Non-Aktif** (status produk tidak akan pernah memicu status `LOW_STOCK`, hanya berpindah dari `IN_STOCK` ke `OUT_OF_STOCK` jika stok `<= 0`).

### Scenario 4: Logika Klasifikasi Status Stok (Tri-State Status Logic)
* **Given** Suatu produk memiliki batas minimum `min_stock_alert = 10`,
* **When** Sistem mengevaluasi kuantitas stok fisik saat ini (`current_stock`):
  * **Kondisi A (`IN_STOCK`)**: Jika `current_stock > 10` ➔ Status = `IN_STOCK` (Badge Hijau Normal).
  * **Kondisi B (`LOW_STOCK`)**: Jika `1 <= current_stock <= 10` ➔ Status = `LOW_STOCK` (Badge Oranye Peringatan).
  * **Kondisi C (`OUT_OF_STOCK`)**: Jika `current_stock <= 0` ➔ Status = `OUT_OF_STOCK` (Badge Merah Habis).

### Scenario 5: Filter Rekap Daftar Belanja Kulakan (*Shopping List Query*)
* **Given** Di warung terdapat 7 produk berstatus `LOW_STOCK` dan 3 produk berstatus `OUT_OF_STOCK`,
* **When** Owner membuka filter atau menu *"Daftar Belanja / Stok Menipis"*,
* **Then** Sistem menyajikan daftar ke-10 produk tersebut yang terurut dari kondisi paling kritis (`OUT_OF_STOCK` di urutan teratas, diikuti `LOW_STOCK` dengan rasio sisa stok terkecil), dengan menampilkan data:
  1. Nama Produk & Barcode/SKU.
  2. Sisa Stok Fisik Saat Ini (dalam `base_unit`).
  3. Nilai Batas Minimum (`min_stock_alert`).
  4. Rekomendasi Satuan Kulakan (jika produk memiliki satuan grosir di `product_units`, tampilkan rasio konversi kemasan, misal: *"1 Dus = 40 Pcs"*).

---

## 4. Business Rules & Edge Cases
- [ ] **Satuan Terkecil Mutlak**: Nilai `min_stock_alert` wajib selalu diukur dalam satuan terkecil (`base_unit`, misal: Pcs, Sachet, Butir) untuk menjamin akurasi dan menghindari kerancuan konversi desimal.
- [ ] **Role Authorization**: Hanya user dengan role `owner` yang berhak melihat dan mengubah angka `min_stock_alert`. User role `cashier` hanya diizinkan melihat badge indikator status visual di antarmuka kasir tanpa melihat angka konfigurasi batasnya.
- [ ] **Data Reference**: Definisi field database mengacu pada tabel `products` di [data-dictionary.md](data-dictionary.md).
