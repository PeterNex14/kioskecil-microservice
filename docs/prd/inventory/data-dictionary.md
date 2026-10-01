# Data Dictionary: Inventory Domain (`db_inventory`)

Dokumen ini adalah **pusat data tunggal (*single source of truth*)** untuk seluruh entitas dan field bisnis di dalam database `db_inventory` (`inventory-service`).

---

## 1. Tabel: `products` (Master Produk & Satuan Dasar)
Menyimpan data identitas pokok barang dagangan dan satuan dasar terkecilnya (*base unit*).

| Field Name | Logical Data Type | Mandatory / Optional | Default Value | Aturan Validasi & Deskripsi Bisnis |
| :--- | :--- | :--- | :--- | :--- |
| `id` | Identifier (UUID) | **Mandatory** | Auto-generated | Kunci primer produk unik secara global. |
| `name` | Text | **Mandatory** | - | Nama produk (Min. 3 karakter, Max. 255 karakter). |
| `barcode` | Text / Alphanumeric | Optional | Auto-generated SKU | Barcode fisik pabrik (EAN-13/UPC) atau SKU lokal unik. |
| `category` | Text | Optional | `"Umum"` | Kategori produk untuk memudahkan filter kasir. |
| `base_unit` | Text | **Mandatory** | `"Pcs"` | Satuan terkecil produk (misal: Pcs, Sachet, Butir, Bungkus). |
| `cost_price` | Currency (BigInt IDR) | **Mandatory** | `0` | Harga modal (HPP/COGS) per satuan dasar. Nilai `>= 0`. |
| `selling_price`| Currency (BigInt IDR) | **Mandatory** | - | Harga jual resmi kasir per satuan dasar. Nilai `> 0`. |
| `is_active` | Boolean | **Mandatory** | `true` | Status aktif produk. Jika `false`, disembunyikan dari kasir. |
| `created_by` | Identifier (UUID) | **Mandatory** | - | ID user (Owner) pembuat master produk. |
| `created_at` | Timestamp with TZ | **Mandatory** | Current Time | Waktu pembuatan data. |
| `updated_at` | Timestamp with TZ | **Mandatory** | Current Time | Waktu perubahan terakhir data. |

---

## 2. Tabel: `product_units` (Satuan Kemasan Bertingkat / Formula Grosir)
Menyimpan rumus kemasan sekunder (misal: Renceng, Dus, Pack, Box) terhadap satuan terkecil beserta harga grosirnya.

| Field Name | Logical Data Type | Mandatory / Optional | Default Value | Aturan Validasi & Deskripsi Bisnis |
| :--- | :--- | :--- | :--- | :--- |
| `id` | Identifier (UUID) | **Mandatory** | Auto-generated | Kunci primer unik satuan kemasan. |
| `product_id` | Identifier (UUID) | **Mandatory** | - | Relasi (*Foreign Key*) ke tabel `products.id`. |
| `unit_name` | Text | **Mandatory** | - | Nama kemasan (misal: "Dus", "Renceng", "Pack", "Box"). |
| `conversion_factor` | Integer | **Mandatory** | - | Rasio pengali ke satuan dasar (Wajib Integer `> 1`). |
| `barcode` | Text / Alphanumeric | Optional | `NULL` | Barcode khusus kardus/box pabrik (jika ada). Wajib unik jika diisi. |
| `cost_price` | Currency (BigInt IDR) | Optional | `0` | Harga modal grosir per kemasan saat kulakan. |
| `selling_price`| Currency (BigInt IDR) | **Mandatory** | - | Harga jual grosir kasir untuk kemasan ini. |
| `is_active` | Boolean | **Mandatory** | `true` | Status aktif satuan kemasan. |
| `created_at` | Timestamp with TZ | **Mandatory** | Current Time | Waktu pembuatan data. |
| `updated_at` | Timestamp with TZ | **Mandatory** | Current Time | Waktu perubahan terakhir data. |

---

## 3. Tabel: `stock_movements` (Buku Besar Audit Mutasi Stok)
Mencatat seluruh histori keluar-masuknya stok fisik barang secara permanen (*immutable ledger*).

| Field Name | Logical Data Type | Mandatory / Optional | Default Value | Aturan Validasi & Deskripsi Bisnis |
| :--- | :--- | :--- | :--- | :--- |
| `id` | Identifier (UUID) | **Mandatory** | Auto-generated | Kunci primer riwayat transaksi mutasi. |
| `product_id` | Identifier (UUID) | **Mandatory** | - | Relasi (*Foreign Key*) ke tabel `products.id`. |
| `movement_type`| Enum / Text | **Mandatory** | - | Tipe mutasi: `RESTOCK`, `SALES`, `ADJUSTMENT_DAMAGE`, `ADJUSTMENT_OWNER_USE`. |
| `quantity` | Integer | **Mandatory** | - | Perubahan kuantitas dalam **satuan terkecil (base unit)**. Positif = masuk, Negatif = keluar. |
| `unit_name` | Text | **Mandatory** | - | Nama satuan yang digunakan saat input (misal: "Dus" atau "Pcs"). |
| `unit_price` | Currency (BigInt IDR) | Optional | `0` | Harga satuan saat mutasi (harga beli saat restock). |
| `notes` | Text | Optional | `NULL` | Catatan (misal: nomor nota belanja atau nama supplier). |
| `created_by` | Identifier (UUID) | **Mandatory** | - | User yang melakukan transaksi mutasi. |
| `created_at` | Timestamp with TZ | **Mandatory** | Current Time | Waktu mutasi terjadi. |
