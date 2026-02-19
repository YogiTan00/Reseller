let currentPage = 1;
const limit = 10;
let orderBy = 'created_at';
let sort = 'desc';

export function getProductList(page = 1, q = '', orderByParam = '', sortParam = '') {
    if (orderByParam) orderBy = orderByParam;
    if (sortParam) sort = sortParam;
    
    const offset = (page - 1) * limit;
    const queryParams = new URLSearchParams({
        q: q,
        orderBy: orderBy,
        limit: limit,
        offset: offset,
        sort: sort
    });
    fetch(`http://localhost:8080/api/v1/product/list?${queryParams}`, {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json'
        }
    })
        .then(response => response.json())
        .then(data => {
            const products = data.data;
            const productList = document.getElementById('product-list');
            const currentPageDisplay = document.getElementById('current-page');

            productList.innerHTML = '';
            currentPageDisplay.textContent = `Halaman ${currentPage}`;

            const totalCount = data.meta ? data.meta.count : products.length;
            const totalPages = Math.ceil(totalCount / limit);

            const nextPageButton = document.getElementById('next-page');
            const prevPageButton = document.getElementById('prev-page');

            nextPageButton.disabled = currentPage >= totalPages;
            prevPageButton.disabled = currentPage <= 1;

            if (products.length === 0) {
                productList.innerHTML = `
                    <div class="empty-state">
                        <i class="bi bi-inbox"></i>
                        <h4>Belum ada produk</h4>
                        <p>Silakan tambahkan produk baru dengan klik tombol "Tambah Produk"</p>
                    </div>
                `;
                return;
            }

            products.forEach(product => {
                const productImage = product.image && product.image !== '' ? product.image : '/html/images/no-image.png';
                const price = parseInt(product.price).toLocaleString('id-ID');
                const defaultPrice = parseInt(product.defaultPrice).toLocaleString('id-ID');

                const productItem = `
                    <li class="col-12 list-group-item d-flex align-items-center">
                        <div class="col-sm-2">
                            <img src="${productImage}" alt="${product.name}">
                        </div>
                        <div class="col-sm-4">
                            <h5>${product.name}</h5>
                            <p class="mb-0"><small class="text-muted">Ukuran: ${product.typeSize || '-'}</small></p>
                        </div>
                        <div class="col-sm-3">
                            <p class="price">Rp ${price}</p>
                            <p class="default-price mb-0">Rp ${defaultPrice}</p>
                        </div>
                        <div class="col-sm-3 col-update">
                            <div class="product-actions">
                                <button class="btn-action btn-edit" onclick="editProduct('${product.id}')">
                                    <i class="bi bi-pencil"></i> Edit
                                </button>
                                <button class="btn-action btn-delete" onclick="deleteProduct('${product.id}')">
                                    <i class="bi bi-trash"></i> Hapus
                                </button>
                            </div>
                        </div>
                    </li>
                `;
                productList.insertAdjacentHTML('beforeend', productItem);
            });
        })
        .catch(error => {
            console.error('Error fetching products:', error);
            const productList = document.getElementById('product-list');
            productList.innerHTML = `
                <div class="empty-state">
                    <i class="bi bi-exclamation-triangle"></i>
                    <h4>Terjadi Kesalahan</h4>
                    <p>Gagal memuat data produk. Silakan coba lagi.</p>
                </div>
            `;
        });
}

const applyFilterBtn = document.getElementById('apply-filter');
if (applyFilterBtn) {
    applyFilterBtn.addEventListener('click', () => {
        const search = document.getElementById('q').value;
        currentPage = 1;
        getProductList(currentPage, search, orderBy, sort);
    });
}

const nextPagerBtn = document.getElementById('next-page');
if (nextPagerBtn) {
    nextPagerBtn.addEventListener('click', () => {
        currentPage += 1;
        const search = document.getElementById('q').value;
        getProductList(currentPage, search, orderBy, sort);
    });
}

const prevPagerBtn = document.getElementById('prev-page');
if (prevPagerBtn) {
    prevPagerBtn.addEventListener('click', () => {
        if (currentPage > 1) {
            currentPage -= 1;
            const search = document.getElementById('q').value;
            getProductList(currentPage, search, orderBy, sort);
        }
    });
}

const sortBtn = document.getElementById('sort');
if (sortBtn) {
    sortBtn.addEventListener('click', function() {
        sort = (sort === 'asc') ? 'desc' : 'asc';
        const search = document.getElementById('q').value;
        getProductList(currentPage, search, orderBy, sort);
    });
}

window.editProduct = function(productId) {
    alert('Fitur edit akan segera tersedia! ID Produk: ' + productId);
};

window.deleteProduct = function(productId) {
    if (confirm('Apakah Anda yakin ingin menghapus produk ini?')) {
        fetch(`http://localhost:8080/api/v1/product/delete/${productId}`, {
            method: 'DELETE',
            headers: {
                'Content-Type': 'application/json'
            }
        })
        .then(response => {
            if (response.ok) {
                alert('Produk berhasil dihapus!');
                getProductList(currentPage, document.getElementById('q').value, orderBy, sort);
            } else {
                alert('Gagal menghapus produk. Silakan coba lagi.');
            }
        })
        .catch(error => {
            console.error('Error deleting product:', error);
            alert('Terjadi kesalahan saat menghapus produk.');
        });
    }
};

document.addEventListener('DOMContentLoaded', () => {
    getProductList();
});
