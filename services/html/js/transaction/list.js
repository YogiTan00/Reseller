console.log('list.js file loaded successfully!');

let currentPage = 1;
const limit = 10;
let orderBy = 'created_at';
let sort = 'desc';

export function getTransactionList(page = 1, q = '', startDate = '', endDate = '') {
    const input = document.getElementById('date-range');
    if (input && input.value) {
        const dates = input.value.split(' - ');
        startDate = dates[0];
        endDate = dates[1];
    }
    const offset = (page - 1) * limit;
    const queryParams = new URLSearchParams({
        q: q,
        orderBy: orderBy,
        limit: limit,
        offset: offset,
        sort: sort,
        startDate: startDate,
        endDate: endDate
    });
    fetch(`http://localhost:8080/api/v1/transaction/list?${queryParams}`, {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json'
        }
    })
        .then(response => response.json())
        .then(data => {
            const transactions = data.data;
            const transactionList = document.getElementById('transaction-list');
            const currentPageDisplay = document.getElementById('current-page');

            transactionList.innerHTML = '';
            currentPageDisplay.textContent = `Halaman ${currentPage}`;

            const totalCount = data.meta ? data.meta.count : transactions.length;
            const totalPages = Math.ceil(totalCount / limit);

            const nextPageButton = document.getElementById('next-page');
            const prevPageButton = document.getElementById('prev-page');

            nextPageButton.disabled = currentPage >= totalPages;
            prevPageButton.disabled = currentPage <= 1;

            if (transactions.length === 0) {
                transactionList.innerHTML = `
                    <div class="empty-state">
                        <i class="bi bi-inbox"></i>
                        <h4>Belum ada transaksi</h4>
                        <p>Silakan tambahkan transaksi baru dengan klik tombol di atas</p>
                    </div>
                `;
                return;
            }

            transactions.forEach(transaction => {
                const date = new Date(transaction.salesDate);
                const formattedDate = date.toLocaleDateString('id-ID', {
                    day: 'numeric',
                    month: 'long',
                    year: 'numeric'
                });
                const isReturn = transaction.returnItem === true || transaction.returnItem === 'true';
                const badgeClass = isReturn ? 'badge-return' : 'badge-sale';
                const badgeText = isReturn ? 'Retur' : 'Penjualan';
                
                const transactionItem = `
                    <li class="col-12 list-group-item d-flex align-items-center">
                        <div class="col-sm-4">
                            <span class="transaction-badge ${badgeClass}">${badgeText}</span>
                            <h5 class="mt-2">${transaction.name}</h5> 
                        </div>
                        <div class="col-sm-3">
                            <p class="mb-1"><i class="bi bi-calendar"></i> ${formattedDate}</p>
                            <p class="mb-0"><i class="bi bi-box-seam"></i> ${transaction.unit} unit</p>
                        </div>
                        <div class="col-sm-2">
                            <p class="mb-0 text-muted">${transaction.description || '-'}</p>
                        </div>
                        <div class="col-sm-3 col-update">
                            <div class="transaction-actions">
                                <button class="btn-action btn-edit" onclick="editTransaction('${transaction.id}')">
                                    <i class="bi bi-pencil"></i> Edit
                                </button>
                                <button class="btn-action btn-delete" onclick="deleteTransaction('${transaction.id}')">
                                    <i class="bi bi-trash"></i> Hapus
                                </button>
                            </div>
                        </div>
                    </li>
                `;
                transactionList.insertAdjacentHTML('beforeend', transactionItem);
            });
        })
        .catch(error => {
            console.error('Error fetching transactions:', error);
            const transactionList = document.getElementById('transaction-list');
            transactionList.innerHTML = `
                <div class="empty-state">
                    <i class="bi bi-exclamation-triangle"></i>
                    <h4>Terjadi Kesalahan</h4>
                    <p>Gagal memuat data transaksi. Silakan coba lagi.</p>
                </div>
            `;
        });
}

const dateRangeBtn = document.getElementById('date-range');
if (dateRangeBtn) {
    $(dateRangeBtn).daterangepicker({
        locale: {
            format: 'YYYY-MM-DD',
            applyLabel: 'Terapkan',
            cancelLabel: 'Batal',
            fromLabel: 'Dari',
            toLabel: 'Sampai',
            customRangeLabel: 'Kustom'
        }
    });
    
    $(dateRangeBtn).on('apply.daterangepicker', function(ev, picker) {
        const startDate = picker.startDate.format('YYYY-MM-DD');
        const endDate = picker.endDate.format('YYYY-MM-DD');
        const search = document.getElementById('q')?.value || "";
        currentPage = 1;
        getTransactionList(currentPage, search, startDate, endDate);
    });
}

document.querySelector('.date-picker-btn').addEventListener('click', function() {
    $(dateRangeBtn).trigger('click');
});

const applyFilterBtn = document.getElementById('apply-filter');
if (applyFilterBtn) {
    applyFilterBtn.addEventListener('click', () => {
        const search = document.getElementById('q').value;
        currentPage = 1;
        getTransactionList(currentPage, search);
    });
}

const nextPagerBtn = document.getElementById('next-page');
if (nextPagerBtn) {
    nextPagerBtn.addEventListener('click', () => {
        currentPage += 1;
        const search = document.getElementById('q').value;
        getTransactionList(currentPage, search);
    });
}

const prevPagerBtn = document.getElementById('prev-page');
if (prevPagerBtn) {
    prevPagerBtn.addEventListener('click', () => {
        if (currentPage > 1) {
            currentPage -= 1;
            const search = document.getElementById('q').value;
            getTransactionList(currentPage, search);
        }
    });
}

const sortBtn = document.getElementById('sort');
if (sortBtn) {
    sortBtn.addEventListener('click', () => {
        sort = (sort === 'asc') ? 'desc' : 'asc';
        const search = document.getElementById('q').value;
        getTransactionList(currentPage, search);
    });
}

window.editTransaction = function(transactionId) {
    alert('Fitur edit akan segera tersedia! ID Transaksi: ' + transactionId);
};

window.deleteTransaction = function(transactionId) {
    if (confirm('Apakah Anda yakin ingin menghapus transaksi ini?')) {
        fetch(`http://localhost:8080/api/v1/transaction/delete/${transactionId}`, {
            method: 'DELETE',
            headers: {
                'Content-Type': 'application/json'
            }
        })
        .then(response => {
            if (response.ok) {
                alert('Transaksi berhasil dihapus!');
                getTransactionList(currentPage, document.getElementById('q').value);
            } else {
                alert('Gagal menghapus transaksi. Silakan coba lagi.');
            }
        })
        .catch(error => {
            console.error('Error deleting transaction:', error);
            alert('Terjadi kesalahan saat menghapus transaksi.');
        });
    }
};

document.addEventListener('DOMContentLoaded', () => {
    getTransactionList();
    // Don't load dropdown here - it will be called after dynamic content loads
});

function loadProductDropdown() {
    console.log('Loading product dropdown...');
    
    fetch('http://localhost:8080/api/v1/product/list', {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json'
        }
    })
    .then(response => {
        console.log('Product list response status:', response.status);
        return response.json();
    })
    .then(data => {
        console.log('Product list data:', data);
        console.log('Number of products:', data.data ? data.data.length : 0);
        window.allProducts = data.data || [];
        
        populateSingleDropdown();
        populateMultipleDropdowns();
    })
    .catch(error => {
        console.error('Error fetching product dropdown:', error);
        alert('Gagal memuat daftar produk. Silakan refresh halaman.');
    });
}

function populateSingleDropdown() {
    console.log('populateSingleDropdown called');
    const idNameSelect = document.getElementById('idName');
    if (!idNameSelect) {
        console.log('idName element not found');
        return;
    }
    console.log('idName element found, options before:', idNameSelect.options.length);

    idNameSelect.innerHTML = '';
    const defaultOption = document.createElement("option");
    defaultOption.value = '';
    defaultOption.text = '-- Pilih Produk --';
    defaultOption.disabled = true;
    defaultOption.selected = true;
    idNameSelect.appendChild(defaultOption);

    window.allProducts.forEach(product => {
        console.log('Processing product:', product);
        const option = document.createElement("option");
        option.value = product.id;
        const name = product.name || 'Unknown Product';
        const typeSize = product.typeSize || '-';
        console.log('Product data - Name:', name, 'Type Size:', product.typeSize, 'Final typeSize:', typeSize);
        option.text = `${name} (${typeSize})`;
        idNameSelect.appendChild(option);
    });

    console.log('idName element options after:', idNameSelect.options.length);

    idNameSelect.addEventListener("change", function() {
        const selectedId = idNameSelect.value;
        const selectedProduct = window.allProducts.find(product => product.id === selectedId);
        if (selectedProduct) {
            const productNameInput = document.getElementById('productName');
            if (productNameInput) {
                productNameInput.value = selectedProduct.name || 'Unknown Product';
            }
        }
    });
}

function populateMultipleDropdowns() {
    console.log('populateMultipleDropdowns called');
    const dropdowns = document.querySelectorAll('.id-name-multiple');
    console.log('Found dropdowns:', dropdowns.length);
    dropdowns.forEach(dropdown => {
        dropdown.innerHTML = '';
        const defaultOption = document.createElement("option");
        defaultOption.value = '';
        defaultOption.text = '-- Pilih Produk --';
        defaultOption.disabled = true;
        defaultOption.selected = true;
        dropdown.appendChild(defaultOption);

        window.allProducts.forEach(product => {
            const option = document.createElement("option");
            option.value = product.id;
            const name = product.name || 'Unknown Product';
            const typeSize = product.typeSize || '-';
            option.text = `${name} (${typeSize})`;
            dropdown.appendChild(option);
        });
    });
}

function populateDropdowns() {
    populateMultipleDropdowns();
}

window.loadProductDropdown = loadProductDropdown;
window.populateDropdowns = populateDropdowns;
