# Modul 1: Inventory & Master Data

* **Microservice Target**: `inventory-service`
* **Database Target**: `db_inventory`
* **Status**: In Progress (Drafting Stories & Acceptance Criteria)
* **Pusat Data Dictionary**: [data-dictionary.md](data-dictionary.md)

---

## 📖 Ringkasan Modul
Modul Inventory & Master Data bertanggung jawab sebagai **pusat data tunggal (*single source of truth*)** untuk seluruh informasi barang dagangan di warung. Modul ini mengelola pendaftaran produk, barcode/SKU, relasi konversi multi-satuan (Dus ke Pcs), penetapan harga modal (COGS) dan harga jual, serta pelacakan pergerakan dan penyesuaian stok fisik.

---

## 📋 Daftar User Stories & Acceptance Criteria

| ID Story | Judul Fitur | Status | File Dokumen |
| :--- | :--- | :--- | :--- |
| **`KK-INV-01`** | Registrasi Master Produk & Penetapan Harga Dasar | ✅ Complete | [KK-INV-01.md](KK-INV-01.md) |
| **`KK-INV-02`** | Konfigurasi Multi-Satuan & Harga Grosir (Formula) | ✅ Complete | [KK-INV-02.md](KK-INV-02.md) |
| **`KK-INV-03`** | Pencatatan Barang Masuk (Restock) & Otomatisasi HPP | ✅ Complete | [KK-INV-03.md](KK-INV-03.md) |
| **`KK-INV-04`** | Pengaturan Batas Stok Minimum (Low Stock Alert) | ✅ Complete | [KK-INV-04.md](KK-INV-04.md) |
| **`KK-INV-05`** | Penyesuaian Stok Manual (Rusak, Expired, Owner Use, Opname) | ⏳ Ready to Refine | [KK-INV-05.md](KK-INV-05.md) |
| **`KK-INV-06`** | Buku Besar Riwayat Mutasi Stok (Stock Audit Ledger) | ⏳ In Backlog | [KK-INV-06.md](KK-INV-06.md) |
| **`KK-INV-07`** | Valuasi Total Nilai Aset Stok Warung | ⏳ In Backlog | [KK-INV-07.md](KK-INV-07.md) |

---

## 🗄️ Database Reference
Untuk spesifikasi seluruh field, tipe data, dan aturan validasi entitas (`products`, `product_units`, `stock_movements`), buka [data-dictionary.md](data-dictionary.md).
