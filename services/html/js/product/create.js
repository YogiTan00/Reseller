document.getElementById('create-product').addEventListener('click', createProduct);

export function createProduct() {
    const name = document.getElementById('name').value.trim();
    const typeSize = document.getElementById('typeSize').value.trim();
    const image = document.getElementById('image').value.trim();
    const defaultPrice = document.getElementById('defaultPrice').value;
    const price = document.getElementById('price').value;

    if (!name || !typeSize || !defaultPrice || !price) {
        alert('Mohon lengkapi semua field yang wajib diisi!');
        return;
    }

    if (parseFloat(defaultPrice) <= 0 || parseFloat(price) <= 0) {
        alert('Harga harus lebih dari 0!');
        return;
    }

    const btn = document.getElementById('create-product');
    const originalText = btn.innerHTML;
    btn.innerHTML = '<i class="bi bi-spinner"></i> Menyimpan...';
    btn.disabled = true;

    fetch('http://localhost:8080/api/v1/product/create', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            name: name,
            typeSize: typeSize,
            image: image,
            defaultPrice: parseFloat(defaultPrice),
            price: parseFloat(price)
        })
    })
        .then(response => response.json())
        .then(data => {
            alert('Produk berhasil ditambahkan!');
            document.getElementById('product-form').reset();
            btn.innerHTML = originalText;
            btn.disabled = false;
            getProductList();
        })
        .catch(error => {
            console.error('Error:', error);
            alert('Gagal menambahkan produk. Silakan coba lagi.');
            btn.innerHTML = originalText;
            btn.disabled = false;
        });
}
