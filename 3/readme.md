## Penjelasan Kompleksitas Algoritma

### Analisis Kode:

1. Iterasi Melalui Karakter: Algoritma ini akan melalui setiap karakter dalam string input, yang membutuhkan waktu O(N).

2. Operasi Stack:

    - Push (Menambahkan elemen ke stack): Operasi ini bekerja dengan kompleksitas waktu O(1), karena elemen ditambahkan di bagian akhir slice.
    - Pop (Menghapus elemen dari stack): Operasi ini bekerja dengan kompleksitas waktu O(1), karena elemen dihapus dari bagian akhir slice.

3. Pengecekan Cocokan Tanda Kurung: Untuk setiap karakter, kita melakukan paling banyak satu operasi push atau pop pada stack. Jumlah operasi push dan pop maksimum adalah N/2 karena setiap karakter tanda kurung dapat digunakan untuk memasukkan atau mengeluarkan elemen dari stack. Oleh karena itu, kompleksitas waktu operasi ini adalah O(N).

### Kesimpulan

Karena seluruh algoritma bergantung pada iterasi karakter dan operasi stack, kompleksitas waktu total algoritma adalah O(N).




